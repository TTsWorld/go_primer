package ctx

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// WithTimeout
func TestContext01(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go func(ctx context.Context, duration time.Duration) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("handle", ctx.Err())
				return
			default:
				fmt.Println("xxxxxx")
			}
			time.Sleep(500 * time.Millisecond)
		}
		//下面函数的参数从 500 ms 调整为 3s，则携程会直接取消
	}(ctx, 500*time.Millisecond)

	time.Sleep(3 * time.Second)
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
