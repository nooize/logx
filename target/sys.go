package target

import (
	"github.com/nooize/lux"
	"os"
)

func ToStderr() lux.Target {
	return &writerTarget{prefix: "STDERR: ", out: os.Stderr}
}

func ToStdout() lux.Target {
	return &writerTarget{prefix: "STDOUT: ", out: os.Stdout}
}
