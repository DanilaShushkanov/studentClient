package bffstudentservice

import (
	"context"
	"github.com/danilashushkanov/studentClient/internal/app/bffstudentservice/adapters"
	"github.com/danilashushkanov/studentClient/internal/dto"
)

func (s *Service) CreateStudent(ctx context.Context, req *dto.StudentItem) (*dto.StudentItem, error) {
	student, err := s.studentServiceClient.StudentClient.CreateStudent(ctx, adapters.CreateStudentItemToPb(req))
	if err != nil {
		return nil, err
	}

	return adapters.StudentItemFromPb(student), nil
}
