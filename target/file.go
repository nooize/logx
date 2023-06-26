package target

import (
	"errors"
	"github.com/nooize/lux"
	"io"
	"os"
)

func File(file *os.File) (lux.Target, error) {
	if file == nil {
		return nil, errors.New("file is nil")
	}
	return &writerTarget{out: file}, nil
}

func FileWithPath(path string) (lux.Target, error) {
	if err := checkFileTarget(path); err != nil {
		return nil, err
	}
	return &writerTarget{}, nil
}

func checkFileTarget(path string) error {
	return nil
}

type writerTarget struct {
	prefix string
	out    io.Writer
}

func (ft *writerTarget) Handle(e lux.Event) error {
	_, err := ft.out.Write([]byte(ft.prefix + e.Message() + "\n"))
	return err
}
