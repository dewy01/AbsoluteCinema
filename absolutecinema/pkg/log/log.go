package log

import (
	"fmt"
	"log/slog"
	"os"
)

func New(level slog.Level, filePath string, isProd bool) (*slog.Logger, error) {
	if !isProd {
		return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:       level,
			AddSource:   false,
			ReplaceAttr: nil,
		})), nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("log file open: %w", err)
	}

	return slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level:       level,
		AddSource:   false,
		ReplaceAttr: nil,
	})), nil
}
