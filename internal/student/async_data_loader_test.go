package student

import (
	"context"
	"errors"
	api "github.com/danilashushkanov/student-service/pkg/studentServiceApi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestDataLoaderImpl_LoadStudents(t *testing.T) {
	t.Run("load students err", func(t *testing.T) {
		te := newTestEnv(t)
		loader := te.client.NewDataLoader(te.ctx)
		studentIDs, _ := getStudentsTestData()

		listStudentsReq := &api.ListStudentRequest{StudentIds: studentIDs}
		expectedErr := errors.New("list students error")
		te.ssServicesMock.EXPECT().ListStudents(gomock.Any(), listStudentsReq).Return(nil, expectedErr)

		loader.ListStudents(listStudentsReq)
		actualErr := loader.Wait()
		assert.ErrorIs(t, expectedErr, actualErr)
	})

	t.Run("load students success", func(t *testing.T) {
		te := newTestEnv(t)
		loader := te.client.NewDataLoader(te.ctx)
		studentIDs, students := getStudentsTestData()

		listStudentsReq := &api.ListStudentRequest{StudentIds: studentIDs}
		listStudentsResp := &api.ListStudentResponse{Students: students}
		te.ssServicesMock.EXPECT().ListStudents(gomock.Any(), listStudentsReq).Return(listStudentsResp, nil)

		loader.ListStudents(listStudentsReq)
		actualErr := loader.Wait()
		actual := loader.Students()
		assert.Equal(t, students, actual)
		assert.NoError(t, actualErr)
	})
}

func getStudentsTestData() ([]int64, []*api.Student) {
	return []int64{
			1,
		},
		[]*api.Student{
			{
				Id:       1,
				FullName: "name",
				Age:      10,
				Salary:   10000,
				Teachers: []*api.Teacher{
					{
						Id:           1,
						PositionType: 1,
						FullName:     "name",
						StudentId:    1,
					},
				},
			},
		}
}

func TestDataLoaderImpl_LoadTeachers(t *testing.T) {
	t.Run("load teachers err", func(t *testing.T) {
		te := newTestEnv(t)
		loader := te.client.NewDataLoader(te.ctx)
		teacherIDs, _ := getTeachersTestData()

		listTeachersReq := &api.ListTeacherRequest{TeacherIds: teacherIDs}
		expectedErr := errors.New("list teachers error")
		te.ssServicesMock.EXPECT().ListTeachers(gomock.Any(), listTeachersReq).Return(nil, expectedErr)

		loader.ListTeachers(listTeachersReq)
		actualErr := loader.Wait()
		assert.ErrorIs(t, expectedErr, actualErr)
	})

	t.Run("load teachers success", func(t *testing.T) {
		te := newTestEnv(t)
		loader := te.client.NewDataLoader(te.ctx)
		teacherIDs, teachers := getTeachersTestData()

		listTeachersReq := &api.ListTeacherRequest{TeacherIds: teacherIDs}
		listTeachersResp := &api.ListTeacherResponse{Teachers: teachers}
		te.ssServicesMock.EXPECT().ListTeachers(gomock.Any(), listTeachersReq).Return(listTeachersResp, nil)

		loader.ListTeachers(listTeachersReq)
		actualErr := loader.Wait()
		actual := loader.Teachers()
		assert.Equal(t, teachers, actual)
		assert.NoError(t, actualErr)
	})
}

func getTeachersTestData() ([]int64, []*api.Teacher) {
	return []int64{
			1,
		},
		[]*api.Teacher{
			{
				Id:           1,
				FullName:     "name",
				PositionType: 1,
				StudentId:    1,
			},
		}
}

func Test_DataLoader(t *testing.T) {
	t.Run("load Students error", func(t *testing.T) {
		te := newTestEnv(t)
		loader := te.client.NewDataLoader(te.ctx)

		studentIDs, _ := getStudentsTestData()
		teacherIDs, teachers := getTeachersTestData()

		listStudentsReq := &api.ListStudentRequest{StudentIds: studentIDs}
		listTeachersReq := &api.ListTeacherRequest{TeacherIds: teacherIDs}

		expectedErr := errors.New("list students error")
		listTeachersResp := &api.ListTeacherResponse{Teachers: teachers}

		te.ssServicesMock.EXPECT().ListStudents(gomock.Any(), listStudentsReq).Return(nil, expectedErr)
		te.ssServicesMock.EXPECT().ListTeachers(gomock.Any(), listTeachersReq).Return(listTeachersResp, nil)

		loader.ListStudents(listStudentsReq)
		loader.ListTeachers(listTeachersReq)

		actualErr := loader.Wait()
		assert.ErrorIs(t, actualErr, expectedErr)
	})

	t.Run("load teachers error", func(t *testing.T) {
		te := newTestEnv(t)
		loader := te.client.NewDataLoader(te.ctx)

		studentIDs, students := getStudentsTestData()
		teacherIDs, _ := getTeachersTestData()

		listStudentsReq := &api.ListStudentRequest{StudentIds: studentIDs}
		listTeachersReq := &api.ListTeacherRequest{TeacherIds: teacherIDs}

		expectedErr := errors.New("list teachers error")
		listStudentsResp := &api.ListStudentResponse{Students: students}

		te.ssServicesMock.EXPECT().ListStudents(gomock.Any(), listStudentsReq).Return(listStudentsResp, nil)
		te.ssServicesMock.EXPECT().ListTeachers(gomock.Any(), listTeachersReq).Return(nil, expectedErr)

		loader.ListStudents(listStudentsReq)
		loader.ListTeachers(listTeachersReq)

		actualErr := loader.Wait()
		assert.ErrorIs(t, actualErr, expectedErr)
	})

	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)
		loader := te.client.NewDataLoader(te.ctx)

		studentIds, students := getStudentsTestData()
		teacherIds, teachers := getTeachersTestData()

		listStudentsReq := &api.ListStudentRequest{StudentIds: studentIds}
		listTeachersReq := &api.ListTeacherRequest{TeacherIds: teacherIds}

		listStudentsResp := &api.ListStudentResponse{Students: students}
		listTeachersResp := &api.ListTeacherResponse{Teachers: teachers}

		te.ssServicesMock.EXPECT().ListStudents(gomock.Any(), listStudentsReq).Return(listStudentsResp, nil)
		te.ssServicesMock.EXPECT().ListTeachers(gomock.Any(), listTeachersReq).Return(listTeachersResp, nil)

		loader.ListStudents(listStudentsReq)
		loader.ListTeachers(listTeachersReq)

		actualErr := loader.Wait()

		actualStudents := loader.Students()
		actualTeachers := loader.Teachers()

		assert.Equal(t, students, actualStudents)
		assert.Equal(t, teachers, actualTeachers)

		assert.NoError(t, actualErr)
	})

	t.Run("error and cancel", func(t *testing.T) {
		te := newTestEnv(t)
		loader := te.client.NewDataLoader(te.ctx)

		studentIDs, students := getStudentsTestData()
		teacherIDs, _ := getTeachersTestData()

		listStudentsReq := &api.ListStudentRequest{StudentIds: studentIDs}
		listTeachersReq := &api.ListTeacherRequest{TeacherIds: teacherIDs}

		listStudentsResp := &api.ListStudentResponse{Students: students}
		firstError := errors.New("list teachers error")

		first := make(chan struct{})

		te.ssServicesMock.EXPECT().ListStudents(gomock.Any(), listStudentsReq).
			Do(func(context.Context, *api.ListStudentRequest, ...grpc.CallOption) {
				first <- struct{}{}
			}).
			Return(listStudentsResp, nil)

		te.ssServicesMock.EXPECT().ListTeachers(gomock.Any(), listTeachersReq).
			Do(func(context.Context, *api.ListTeacherRequest, ...grpc.CallOption) {
				<-first
			}).
			Return(nil, firstError)

		loader.ListStudents(listStudentsReq)
		loader.ListTeachers(listTeachersReq)

		actualErr := loader.Wait()
		assert.ErrorIs(t, actualErr, firstError)
	})

	t.Run("parent cancel", func(t *testing.T) {
		te := newTestEnv(t)
		ctxWithCancel, cancel := context.WithCancel(te.ctx)
		loader := te.client.NewDataLoader(ctxWithCancel)

		studentIDs, students := getStudentsTestData()
		teacherIDs, _ := getTeachersTestData()

		listStudentsReq := &api.ListStudentRequest{StudentIds: studentIDs}
		listTeachersReq := &api.ListTeacherRequest{TeacherIds: teacherIDs}

		listStudentsResp := &api.ListStudentResponse{Students: students}
		operationClosedError := errors.New("context done")

		te.ssServicesMock.EXPECT().ListStudents(gomock.Any(), listStudentsReq).
			Return(listStudentsResp, nil)

		te.ssServicesMock.EXPECT().ListTeachers(gomock.Any(), listTeachersReq).
			Do(func(ctx context.Context, req *api.ListTeacherRequest, opts ...grpc.CallOption) {
				select {
				case <-ctx.Done():
				case <-time.After(time.Second * 5):
					assert.FailNow(t, "context done not receive")
				}
			}).
			Return(nil, operationClosedError)

		loader.ListStudents(listStudentsReq)
		loader.ListTeachers(listTeachersReq)

		cancel()

		actualErr := loader.Wait()

		assert.ErrorIs(t, actualErr, ErrorContextCanceled)
	})
}
