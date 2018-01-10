package edd_log

import "testing"

var logger *Logger

func init() {
	logger=NewLogger("Woo")
}

func TestLogger_Info(t *testing.T) {
	logger.Info("Waht")
}
