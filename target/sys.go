package target

import (
	"golang.org/x/exp/slog"
	"os"
)

func ToStderr() slog.Handler {
	return &writerTarget{prefix: "STDERR: ", out: os.Stderr}
}

func ToStdout() slog.Handler {
	return &writerTarget{prefix: "STDOUT: ", out: os.Stdout}
}
