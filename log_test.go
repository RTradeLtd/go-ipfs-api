package shell

import (
	"context"
	"testing"
)

func Test_Logger(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sh := NewLocalShell()
	logger, err := sh.GetLogs(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := logger.Close(); err != nil {
			t.Fatal(err)
		}
	}()
	_, err = logger.Next()
	if err != nil {
		t.Fatal(err)
	}
}
