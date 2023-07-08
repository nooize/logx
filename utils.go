package logx

import (
	"context"
	"golang.org/x/exp/slices"
	"golang.org/x/exp/slog"
)

type BaseHandler struct {
	groupPrefix string   // for text: prefix of groups opened in preformatting
	groups      []string // all groups started from WithGroup
}

func (h *BaseHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}

func (h *BaseHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

func (h *BaseHandler) WithAttrs(as []slog.Attr) slog.Handler {
	//out := h.clone()
	// Pre-format the attributes as an optimization.
	/*
		prefix := bytes.NewBuffer([]byte{})
		defer prefix.Reset()
		prefix.WriteString(h.groupPrefix)
		state := out.newHandleState((*buffer.Buffer)(&out.preformattedAttrs), false, "", prefix)
		defer state.free()
		if len(h2.preformattedAttrs) > 0 {
			state.sep = h.attrSep()
		}
		state.openGroups()
		for _, a := range as {
			state.appendAttr(a)
		}
		// Remember the new prefix for later keys.
		out2.groupPrefix = state.prefix.String()
		// Remember how many opened groups are in preformattedAttrs,
		// so we don't open them again when we handle a Record.
		out.nOpenGroups = len(out.groups)
	*/
	return nil
}

func (h *BaseHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	out := h.clone()
	out.groups = append(out.groups, name)
	return out
}

func (h *BaseHandler) clone() *BaseHandler {
	// We can't use assignment because we can't copy the mutex.
	return &BaseHandler{
		//json:              h.json,
		//opts:              h.opts,
		//preformattedAttrs: slices.Clip(h.preformattedAttrs),
		groupPrefix: h.groupPrefix,
		groups:      slices.Clip(h.groups),
		//nOpenGroups:       h.nOpenGroups,
		//w:                 h.w,
	}
}

type nullTarget struct {
	BaseHandler
}

func (t *nullTarget) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}
