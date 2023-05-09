package target

import (
	"github.com/nooize/ltt"
	"os"
)

func ToStderr() ltt.Target {
	return &writerTarget{prefix: "STDERR: ", out: os.Stderr}
}

func ToStdout() ltt.Target {
	return &writerTarget{prefix: "STDOUT: ", out: os.Stdout}
}
