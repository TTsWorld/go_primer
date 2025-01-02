package errgroup

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

func TestErrGroupClosure(t *testing.T) {
	errCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var a = 1
	group, errCtx := errgroup.WithContext(errCtx)
	group.Go(func() error {
		a = 10
		time.Sleep(3 * time.Second)
		fmt.Println("running")
		return nil
	})
	group.Go(func() error {
		return errors.New("error")
	})
	_ = group.Wait()
	fmt.Println(a)
}
