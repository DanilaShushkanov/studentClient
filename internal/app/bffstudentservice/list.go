package bffstudentservice

import (
	"context"
	"github.com/danilashushkanov/studentClient/internal/app/bffstudentservice/adapters"
	"github.com/danilashushkanov/studentClient/internal/dto"
)

func (s *Service) List(ctx context.Context, req *dto.ListStudentAndTeachersRequest) (*dto.ListStudentAndTeachersResponse, error) {
	loader := s.studentServiceClient.NewDataLoader(ctx)

	loader.ListStudents(adapters.ListStudentToPb(req))
	loader.ListTeachers(adapters.ListTeacherToPb(req))

	err := loader.Wait()
	if err != nil {
		return nil, err
	}

	return &dto.ListStudentAndTeachersResponse{
		StudentList: adapters.StudentsFromPb(loader.Students()),
		TeacherList: adapters.TeachersFromPb(loader.Teachers()),
	}, nil
}
