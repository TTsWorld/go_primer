package signal

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestSignal01(t *testing.T) {
	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	<-sigchan
	fmt.Println("shutdown server")
}
