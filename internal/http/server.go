package http

import (
	"context"
	"errors"
	"github.com/danilashushkanov/studentClient/internal/config"
	"github.com/danilashushkanov/studentClient/internal/student"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Server struct {
	config *config.Config
	ctx    context.Context
	e      *echo.Echo
	client *student.Client
}

func New(ctx context.Context, cfg *config.Config, client *student.Client) *Server {
	return &Server{ctx: ctx, config: cfg, client: client}
}

func (s *Server) Start() {
	s.e = echo.New()
	s.e.HidePort = true
	s.e.HideBanner = true
	s.SetupRoutes()

	go func() {
		if err := s.e.Start(s.config.HTTPAddr); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("error start http server, err: %+v", err)
			}
		}
	}()
}

func (s *Server) Stop() error {
	if err := s.e.Shutdown(s.ctx); err != nil {
		if errors.Is(err, context.Canceled) {
			return nil
		}
		log.Printf("error stop http server, err: %+v", err)
		return err
	}
	return nil
}
