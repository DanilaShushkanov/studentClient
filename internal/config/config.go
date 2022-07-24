package config

type Config struct {
	HTTPAddr       string `env:"HTTP_ADDR"`
	GRPCClientAddr string `env:"GRPC_CLIENT"`
	StudentTimeout int64  `env:"STUDENT_TIMEOUT"`
}
