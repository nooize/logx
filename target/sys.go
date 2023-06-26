package target

import (
	"github.com/nooize/lux"
	"os"
)

func ToStderr() lwr.Target {
	return &writerTarget{prefix: "STDERR: ", out: os.Stderr}
}

func ToStdout() lwr.Target {
	return &writerTarget{prefix: "STDOUT: ", out: os.Stdout}
}
