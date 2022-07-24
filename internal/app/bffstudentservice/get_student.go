package bffstudentservice

import (
	"context"
	api "github.com/danilashushkanov/student-service/pkg/studentServiceApi"
	"github.com/danilashushkanov/studentClient/internal/app/bffstudentservice/adapters"
	"github.com/danilashushkanov/studentClient/internal/dto"
)

func (s *Service) GetStudent(ctx context.Context, req *api.GetStudentRequest) (*dto.StudentItem, error) {
	student, err := s.studentServiceClient.StudentClient.GetStudent(ctx, req)
	if err != nil {
		return nil, err
	}

	return adapters.StudentItemFromPb(student), nil
}
