package student

import (
	"context"
	"fmt"
	api "github.com/danilashushkanov/student-service/pkg/studentServiceApi"
	"sync"
)

var ErrorContextCanceled = fmt.Errorf("global context canceled: %w", context.Canceled)

type DataLoader struct {
	studentServiceClient *Client
	ctx                  context.Context
	cancel               context.CancelFunc
	wg                   *sync.WaitGroup
	errCh                chan error

	mu       *sync.RWMutex
	students []*api.Student
	teachers []*api.Teacher
}

func (c *Client) NewDataLoader(ctx context.Context) *DataLoader {
	ctx, cancel := context.WithCancel(ctx)
	return &DataLoader{
		studentServiceClient: c,
		wg:                   &sync.WaitGroup{},
		mu:                   &sync.RWMutex{},
		errCh:                make(chan error),

		ctx:    ctx,
		cancel: cancel,
	}
}

func (d *DataLoader) Students() []*api.Student {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.students
}

func (d *DataLoader) Teachers() []*api.Teacher {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.teachers
}

func (d *DataLoader) ListStudents(req *api.ListStudentRequest) {
	d.wg.Add(1)
	go func(ctx context.Context, req *api.ListStudentRequest) {
		defer d.wg.Done()

		students, err := d.studentServiceClient.StudentClient.ListStudents(ctx, req)
		if err != nil {
			d.handleLoadError(err)
			return
		}

		d.mu.Lock()
		defer fmt.Println("listStudents loaded")
		defer d.mu.Unlock()
		d.students = students.Students
	}(d.ctx, req)
}

func (d *DataLoader) ListTeachers(req *api.ListTeacherRequest) {
	d.wg.Add(1)
	go func(ctx context.Context, req *api.ListTeacherRequest) {
		defer d.wg.Done()

		teachers, err := d.studentServiceClient.TeacherClient.ListTeachers(ctx, req)
		if err != nil {
			d.handleLoadError(err)
			return
		}

		d.mu.Lock()
		defer fmt.Println("listTeachers loaded")
		defer d.mu.Unlock()
		d.teachers = teachers.Teachers
	}(d.ctx, req)
}

func (d *DataLoader) handleLoadError(err error) {
	select {
	case <-d.ctx.Done():
	default:
		select {
		case d.errCh <- err:
		case <-d.ctx.Done():
		}
	}
}

func (d *DataLoader) Wait() error {
	done := make(chan struct{}, 1)
	go func() {
		d.wg.Wait()
		done <- struct{}{}
		close(done)
	}()

	select {
	case <-done:
		defer d.cancel()
		select {
		case <-d.ctx.Done():
			return ErrorContextCanceled
		default:
			return nil
		}

	case err := <-d.errCh:
		d.cancel()
		<-done
		return err
	}
}
