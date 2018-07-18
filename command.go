package wifiscan

import (
	"context"
	"io"
	"os/exec"
)

var commandContext func(context.Context, string, ...string) cmd = func(ctx context.Context, name string, args ...string) cmd {
	return exec.CommandContext(ctx, name, args...)
}

type cmd interface {
	StdoutPipe() (io.ReadCloser, error)
	Wait() error
}
