package go_test_test

import (
	"fmt"
	"go_primer/go_test"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {

	u1 := go_test.UserInfo{ID: 1, UID: 0, IsOpenGame: 0, Coin: 0, Assets: 0, GarageLen: 0, ParkLen: 0}
	u2 := go_test.UserInfo{ID: 1, UID: 0, IsOpenGame: 0, Coin: 0, Assets: 0, GarageLen: 0, ParkLen: 0, CreateTime: time.Now(), UpdateTime: time.Now()}

	fmt.Println(cmp.Equal(u1, u2, cmp.Transformer("omitTime", go_test.OmitTime)))
}
