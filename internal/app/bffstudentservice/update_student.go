package bffstudentservice

import (
	"context"
	"github.com/danilashushkanov/studentClient/internal/app/bffstudentservice/adapters"
	"github.com/danilashushkanov/studentClient/internal/dto"
)

func (s *Service) UpdateStudent(ctx context.Context, req *dto.StudentItem) (*dto.StudentItem, error) {
	student, err := s.studentServiceClient.StudentClient.PatchStudent(ctx, adapters.UpdateStudentItemToPb(req))
	if err != nil {
		return nil, err
	}

	return adapters.StudentItemFromPb(student), nil
}
