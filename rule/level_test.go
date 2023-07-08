package rule

import (
	"golang.org/x/exp/slog"
	"testing"
)

func TestLevelGreater(t *testing.T) {
	testFun := LevelGreater(slog.LevelInfo)
	if testFun(slog.Record{Level: slog.LevelDebug}) {
		t.Error("LevelGreater(Info) with Debug check return true")
	}
	if !testFun(slog.Record{Level: slog.LevelError}) {
		t.Error("LevelGreater(Info) with Error check return false")
	}
}

func TestLevelLower(t *testing.T) {
	testFun := LevelLower(slog.LevelInfo)
	if !testFun(slog.Record{Level: slog.LevelDebug}) {
		t.Error("LevelLower(Info) with Debug check return false")
	}
	if testFun(slog.Record{Level: slog.LevelError}) {
		t.Error("LevelLower(Info) with Error check return  true")
	}
}

func TestLevelIn(t *testing.T) {
	testFun := LevelIn(slog.LevelWarn, slog.LevelError)
	if testFun(slog.Record{Level: slog.LevelDebug}) {
		t.Error("LevelIn(Wrn, Err) with Debug check return true")
	}
	if !testFun(slog.Record{Level: slog.LevelError}) {
		t.Error("LevelIn(Wrn, Err) with Error check return  false")
	}
}

func TestLevelNotIn(t *testing.T) {
	testFun := LevelNotIn(slog.LevelWarn, slog.LevelError)
	if !testFun(slog.Record{Level: slog.LevelDebug}) {
		t.Error("LevelNotIn(Wrn, Err) with Debug check return false")
	}
	if testFun(slog.Record{Level: slog.LevelError}) {
		t.Error("LevelNotIn(Wrn, Err) with Error check return true")
	}
}
