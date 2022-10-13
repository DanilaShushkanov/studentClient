package http

func (s *Server) SetupRoutes() {
	s.e.POST("student-service/studentsAndTeachers", s.list)

	s.e.POST("/student-service/student", s.createStudentHandler)
	s.e.GET("/student-service/student/:id", s.getStudentHandler)
	s.e.PUT("/student-service/student/:id", s.patchStudentHandler)
	s.e.DELETE("/student-service/student/:id", s.deleteStudentHandler)

	s.e.PUT("/student-service/teacher/:id", s.patchTeacherHandler)
	s.e.POST("/student-service/teacher", s.createTeacherHandler)
}
