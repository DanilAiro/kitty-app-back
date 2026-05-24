package utils

import (
	"os"
	"os/signal"
	"syscall"
)

// Обрабатывает сигнал завершения
func HandleTermination() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}