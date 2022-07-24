package student

import (
	"context"
	api "github.com/danilashushkanov/student-service/pkg/studentServiceApi"
	"github.com/danilashushkanov/studentClient/internal/closer"
	"github.com/danilashushkanov/studentClient/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Services interface {
	api.StudentServiceClient
	api.TeacherServiceClient
}

type Client struct {
	StudentClient api.StudentServiceClient
	TeacherClient api.TeacherServiceClient
}

func NewClient(cfg *config.Config) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.StudentTimeout)*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		cfg.GRPCClientAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(),
	)
	if err != nil {
		return nil, err
	}
	closer.Add(conn.Close)

	res := &Client{
		StudentClient: api.NewStudentServiceClient(conn),
		TeacherClient: api.NewTeacherServiceClient(conn),
	}

	return res, nil
}
