package student

import (
	"context"
	"github.com/golang/mock/gomock"
	"testing"
)

type testEnv struct {
	ctx  context.Context
	ctrl *gomock.Controller

	ssServicesMock *MockServices

	client *Client
}

func newTestEnv(t *testing.T) *testEnv {
	te := &testEnv{
		ctx:  context.Background(),
		ctrl: gomock.NewController(t),
	}

	te.ssServicesMock = NewMockServices(te.ctrl)
	te.client = &Client{
		StudentClient: te.ssServicesMock,
		TeacherClient: te.ssServicesMock,
	}

	return te
}
