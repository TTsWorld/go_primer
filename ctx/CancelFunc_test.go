package ctx

import (
	"context"
	"testing"
)

func TestCancelFunc01(t *testing.T) {
	var cancelFunc context.CancelFunc
	cancelFunc = context.WithCancel(context.Background()).(context.CancelFunc)

}
