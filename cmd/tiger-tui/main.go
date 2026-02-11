package main

import (
	"context"
	"fmt"
	"os"
	"syscall"

	"github.com/fd1az/tiger-tui/internal/logger"
	"github.com/fd1az/tiger-tui/pkg/ui"
)

func main() {
	// Logger to file (TUI owns stdout)
	logFile, err := os.OpenFile("tiger-tui.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening log file: %v\n", err)
		os.Exit(1)
	}
	defer logFile.Close()

	// Redirect stderr (fd 2) to the log file.
	// The TigerBeetle native client writes warnings directly to fd 2,
	// which would corrupt the TUI. This sends them to the log file instead.
	syscall.Dup2(int(logFile.Fd()), 2)

	log := logger.New(logFile, logger.LevelInfo, "tiger-tui", nil)
	log.Info(context.Background(), "starting tiger-tui")

	if err := ui.Run(); err != nil {
		log.Error(context.Background(), "tui error", "error", err)
		os.Exit(1)
	}

	log.Info(context.Background(), "tiger-tui exited")
}
