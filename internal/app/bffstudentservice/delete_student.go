package bffstudentservice

import (
	"context"
	api "github.com/danilashushkanov/student-service/pkg/studentServiceApi"
)

func (s *Service) DeleteStudent(ctx context.Context, req *api.GetStudentRequest) error {
	_, err := s.studentServiceClient.StudentClient.DeleteStudent(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
