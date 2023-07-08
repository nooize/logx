package target

import (
	"context"
	"errors"
	"github.com/nooize/logx"
	"golang.org/x/exp/slog"
	"io"
	"os"
)

func File(file *os.File) (slog.Handler, error) {
	if file == nil {
		return nil, errors.New("file is nil")
	}
	return &writerTarget{out: file}, nil
}

func FileWithPath(path string) (slog.Handler, error) {
	if err := checkFileTarget(path); err != nil {
		return nil, err
	}
	return &writerTarget{}, nil
}

func checkFileTarget(path string) error {
	return nil
}

type writerTarget struct {
	logx.BaseHandler
	prefix string
	out    io.Writer
}

func (ft *writerTarget) Handle(ctx context.Context, rec slog.Record) error {
	_, err := ft.out.Write([]byte(ft.prefix + rec.Message + "\n"))
	return err
}
