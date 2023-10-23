package handler

import (
	"encoding/json"
	"github.com/gookit/slog"
	"io"
)

func (h *Handler) parseBody(from io.ReadCloser, to interface{}) error {
	body, err := io.ReadAll(from)
	if err != nil || len(body) == 0 {
		slog.Errorf("Failed to read body: %v", err.Error())
		return err
	}

	err = json.Unmarshal(body, to)
	if err != nil {
		return err
	}

	return nil
}
