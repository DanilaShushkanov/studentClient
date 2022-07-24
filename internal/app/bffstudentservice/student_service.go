package bffstudentservice

import (
	"context"
	api "github.com/danilashushkanov/student-service/pkg/studentServiceApi"
	"github.com/danilashushkanov/studentClient/internal/dto"
	"github.com/danilashushkanov/studentClient/internal/student"
)

type StudentServiceInterface interface {
	CreateStudent(context.Context, *dto.StudentItem) (*dto.StudentItem, error)
	GetStudent(context.Context, *api.GetStudentRequest) (*dto.StudentItem, error)
	UpdateStudent(context.Context, *dto.StudentItem) (*dto.StudentItem, error)
	DeleteStudent(context.Context, *api.GetStudentRequest) error

	CreateTeacher(context.Context, *dto.TeacherItem) (*dto.TeacherItem, error)
	UpdateTeacher(context.Context, *dto.TeacherItem) (*dto.TeacherItem, error)

	List(context.Context, *dto.ListStudentAndTeachersRequest) (*dto.ListStudentAndTeachersResponse, error)
}

type Service struct {
	studentServiceClient *student.Client
}

func New(c *student.Client) *Service {
	return &Service{studentServiceClient: c}
}
