package ctx

import (
	"context"
	"fmt"
	"testing"
	"time"
)

//WithTimeout
func TestContext01(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go func(ctx context.Context, duration time.Duration) {
		select {
		case <-ctx.Done():
			fmt.Println("handle", ctx.Err())
		case <-time.After(duration):
			fmt.Println("process request with", duration)
		}
	}(ctx, 1500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func TestContext02(t *testing.T) {
	ctx := context.WithValue(context.Background(), "k", "v")
	ctx = context.WithValue(ctx, "k1", "v")
	ctx = context.WithValue(ctx, "k2", "v")
	ctx = context.WithValue(ctx, "k3", "v")
	ctx = context.WithValue(ctx, "k4", "v")
	ctx = context.WithValue(ctx, "k5", "v5")
	fmt.Printf("%+v\n", ctx)

	a := ctx.Value("k5").(string)
	println(a)

}
