package console

import (
	"bytes"
	"io"
	"os"
	"sync"
	"time"
)

const (
	consoleDefaultTimeFormat = time.Kitchen
)

var (
	consoleBufPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 0, 100))
		},
	}
)

func NewConsoleTarget() ltt.Target {
	t := ConsoleTarget{
		Out:        os.Stdout,
		TimeFormat: consoleDefaultTimeFormat,
		Tags:       []string{},
	}

	//for _, opt := range options {
	//	opt(&w)
	//}

	// Fix color on Windows
	if t.Out == os.Stdout || t.Out == os.Stderr {
		t.Out = colorable.NewColorable(t.Out.(*os.File))
	}

	return t
}

type ConsoleTarget struct {
	// Out is the output destination.
	Out io.Writer

	// TimeFormat specifies the format for timestamp in output.
	TimeFormat string

	// Tags defines the order of parts in output.
	// if is empty then all not listed tags wil be skipped.
	Tags []string
}

// Write transforms the JSON input with formatters and appends to w.Out.
func (t ConsoleTarget) Handle(time.Time, ltt.Level, string, *ltt.Tags) error {
	// Fix color on Windows
	if t.Out == os.Stdout || t.Out == os.Stderr {
		t.Out = colorable.NewColorable(t.Out.(*os.File))
	}

	//if w.PartsOrder == nil {
	//	w.PartsOrder = consoleDefaultPartsOrder()
	//}

	var buf = consoleBufPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		consoleBufPool.Put(buf)
	}()

	var evt map[string]interface{}

	for _, p := range w.PartsOrder {
		t.writePart(buf, evt, p)
	}

	w.writeFields(evt, buf)

	if w.FormatExtra != nil {
		err = w.FormatExtra(evt, buf)
		if err != nil {
			return n, err
		}
	}

	err := buf.WriteByte('\n')
	if err != nil {
		return err
	}
	_, err = buf.WriteTo(t.Out)
	return err
}

// writePart appends a formatted part to buf.
func (t ConsoleTarget) writePart(buf *bytes.Buffer, evt map[string]interface{}, p string) {
	var f Formatter

	switch p {
	case LevelFieldName:
		if w.FormatLevel == nil {
			f = consoleDefaultFormatLevel(w.NoColor)
		} else {
			f = w.FormatLevel
		}
	case TimestampFieldName:
		if w.FormatTimestamp == nil {
			f = consoleDefaultFormatTimestamp(w.TimeFormat, w.NoColor)
		} else {
			f = w.FormatTimestamp
		}
	case MessageFieldName:
		if w.FormatMessage == nil {
			f = consoleDefaultFormatMessage
		} else {
			f = w.FormatMessage
		}
	case CallerFieldName:
		if w.FormatCaller == nil {
			f = consoleDefaultFormatCaller(w.NoColor)
		} else {
			f = w.FormatCaller
		}
	default:
		if w.FormatFieldValue == nil {
			f = consoleDefaultFormatFieldValue
		} else {
			f = w.FormatFieldValue
		}
	}

	var s = f(evt[p])

	if len(s) > 0 {
		if buf.Len() > 0 {
			buf.WriteByte(' ') // Write space only if not the first part
		}
		buf.WriteString(s)
	}
}
