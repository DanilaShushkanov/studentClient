package http

import (
	api "github.com/danilashushkanov/student-service/pkg/studentServiceApi"
	"github.com/danilashushkanov/studentClient/internal/app/bffstudentservice"
	"github.com/danilashushkanov/studentClient/internal/dto"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (s *Server) createStudentHandler(c echo.Context) error {
	req := &dto.StudentItem{}
	if err := c.Bind(req); err != nil {
		return err
	}

	studentService := bffstudentservice.New(s.client)
	student, err := studentService.CreateStudent(s.ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, student)
}

func (s *Server) getStudentHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	req := &api.GetStudentRequest{
		Id: int64(id),
	}

	studentService := bffstudentservice.New(s.client)
	student, err := studentService.GetStudent(s.ctx, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, student)
}

func (s *Server) patchStudentHandler(c echo.Context) error {
	studentDto := &dto.StudentItem{}
	if err := c.Bind(studentDto); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	studentDto.ID = int64(id)

	studentService := bffstudentservice.New(s.client)
	student, err := studentService.UpdateStudent(s.ctx, studentDto)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, student)
}

func (s *Server) deleteStudentHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	req := &api.GetStudentRequest{
		Id: int64(id),
	}

	studentService := bffstudentservice.New(s.client)
	err := studentService.DeleteStudent(s.ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "success")
}

func (s *Server) createTeacherHandler(c echo.Context) error {
	teacherDto := &dto.TeacherItem{}
	if err := c.Bind(teacherDto); err != nil {
		return err
	}

	studentService := bffstudentservice.New(s.client)
	teacher, err := studentService.CreateTeacher(s.ctx, teacherDto)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, teacher)
}

func (s *Server) patchTeacherHandler(c echo.Context) error {
	teacherDto := &dto.TeacherItem{}
	if err := c.Bind(teacherDto); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	teacherDto.ID = int64(id)

	studentService := bffstudentservice.New(s.client)
	teacher, err := studentService.UpdateTeacher(s.ctx, teacherDto)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, teacher)
}

func (s *Server) list(c echo.Context) error {
	listDto := &dto.ListStudentAndTeachersRequest{}
	if err := c.Bind(listDto); err != nil {
		return err
	}

	studentService := bffstudentservice.New(s.client)
	list, err := studentService.List(s.ctx, listDto)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, list)
}
