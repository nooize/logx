package rule

import (
	"github.com/nooize/logx"
	"golang.org/x/exp/slog"
	"testing"
)

func TestAnd(t *testing.T) {
	testFun := LevelGreater(slog.LevelInfo)
	if testFun(slog.Record{Level: slog.LevelDebug}) {
		t.Error("LevelGreater(Info) with Debug check return true")
	}
	if !testFun(slog.Record{Level: slog.LevelError}) {
		t.Error("LevelGreater(Info) with Error check return false")
	}
}

func TestOr(t *testing.T) {
	testFun := LevelLower(slog.LevelInfo)
	if !testFun(slog.Record{Level: slog.LevelDebug}) {
		t.Error("LevelLower(Info) with Debug check return false")
	}
	if testFun(slog.Record{Level: slog.LevelError}) {
		t.Error("LevelLower(Info) with Error check return  true")
	}
}

func TestNot(t *testing.T) {
	tests := []logx.Rule{
		func(_ slog.Record) bool { return true },
		func(_ slog.Record) bool { return false },
	}
	for _, origFn := range tests {
		origRes := origFn(slog.Record{})
		if res := Not(origFn)(slog.Record{Level: slog.LevelDebug}); res == origRes {
			t.Errorf("Not(...) expect %v but got %v", !origRes, res)
		}
	}
}

func TestBool(t *testing.T) {
	if !Bool(true)(slog.Record{}) {
		t.Error("Bool(true) return false")
	}
	if Bool(false)(slog.Record{}) {
		t.Error("Bool(true) return true")
	}
}
