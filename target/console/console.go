package console

import (
	"bytes"
	"context"
	"github.com/fatih/color"
	"golang.org/x/exp/slog"
	"os"
	"sync"
	"time"
)

const (
	DefaultFormat     = "%T %L %M %A"
	DefaultTimeFormat = time.DateTime
)

var (
	consoleBufPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 0, 100))
		},
	}
	timeCol = color.New(color.FgWhite)
	errCol  = color.New(color.FgRed)
	wrnCol  = color.New(color.FgYellow)
	infCol  = color.New(color.FgCyan)
	dbgCol  = color.New(color.FgHiMagenta)
	tagCol  = color.New(color.FgGreen)
	darkCol = color.New(color.FgHiBlack)
)

type TargetOption func(*consoleHandler)

func NewConsoleHandler(options ...TargetOption) slog.Handler {
	t := &consoleHandler{
		separator:  []byte(" "),
		timeFormat: DefaultTimeFormat,
	}
	t.Handler = slog.NewTextHandler(os.Stdout, nil)
	time.Now().Format(t.timeFormat)
	for _, opt := range options {
		opt(t)
	}
	return t
}

type consoleHandler struct {
	slog.Handler

	separator  []byte
	colored    bool
	timeFormat string
}

// implement slog.Handler interface

func (h *consoleHandler) Handle(ctx context.Context, rec slog.Record) error {
	// Fix color on Windows

	buf := new(bytes.Buffer)
	h.writeTime(buf, rec.Time)
	buf.Write(h.separator)
	h.writeLevel(buf, rec.Level)
	buf.Write(h.separator)
	h.writeMsg(buf, rec.Message)
	if rec.NumAttrs() > 0 {
		buf.Write(h.separator)
	}
	rec.Attrs(func(attr slog.Attr) bool {
		h.writeTag(buf, attr)
		return true
	})
	buf.Write([]byte("\n"))

	_, err := buf.WriteTo(h.out)
	return err
}

func (m *consoleHandler) writeTime(buf *bytes.Buffer, t time.Time) {
	buf.WriteString(timeCol.Sprint(t.Format(m.timeFormat)))
}

func (m *consoleHandler) writeLevel(buf *bytes.Buffer, l slog.Level) {
	switch l {
	//case lux.Trace:
	//	buf.WriteString("TRC")
	case slog.LevelDebug:
		dbgCol.Fprint(buf, "DBG")
	case slog.LevelInfo:
		infCol.Fprint(buf, "INF")
	case slog.LevelWarn:
		wrnCol.Fprint(buf, "WRN")
	case slog.LevelError:
		errCol.Fprint(buf, "ERR")
	default:
		timeCol.Fprint(buf, "-?-")
	}

}

func (m *consoleHandler) writeTag(buf *bytes.Buffer, attr slog.Attr) {
	buf.WriteString(" ")
	timeCol.Fprint(buf, attr.Key)
	darkCol.Fprint(buf, "=")
	tagCol.Fprint(buf, attr.Value.String())
}

func (t consoleHandler) writeMsg(buf *bytes.Buffer, msg string) {
	buf.WriteString(msg)
}

func parseFormat(v string) {
	//parts := strings.Split(v, "")
}
