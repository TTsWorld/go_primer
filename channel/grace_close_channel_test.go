package channel

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// 1:1, one sender， one receiver
func TestClose01(t *testing.T) {
	ch := make(chan int)

	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()

	for i := 0; i < 10; i++ {
		ch <- i
		if i == 8 {
			close(ch)
			select {
			case <-time.After(time.Second * 3):
				return
			}
		}

	}

}

// 1:N, one sender, N receiver
func TestClose02(t *testing.T) {
	ch := make(chan int)
	ch1 := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			j, ok := <-ch
			if ok {
				fmt.Println(j)
				ch1 <- j
			}
		}()

	}

	for i := 0; i < 100; i++ {
		ch <- i
		<-ch1
		if i == 8 {
			close(ch)
			select {
			case <-time.After(time.Second * 3):
				return
			}
		}

	}

}

// N:1, N sender, 1 receiver
func TestClose03(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	const Max = 100000
	const NumSenders = 1000
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	// senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				select {
				case <-stopCh:
					return
				case dataCh <- rand.Intn(Max):
				}
			}
		}()
	}
	// the receiver
	go func() {
		for value := range dataCh {
			if value == Max-1 {
				fmt.Println("send stop signal to senders.")
				close(stopCh)
				return
			}
			fmt.Println(value)
		}
	}()
	select {
	case <-time.After(time.Hour):
	}

}

// N:M, N sender, M receiver
func TestClose04(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	const Max = 100000
	const NumReceivers = 10
	const NumSenders = 1000
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{}) // It must be a buffered channel.
	toStop := make(chan string, 1)
	var stoppedBy string
	stoppedBy = stoppedBy
	// moderator
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}
				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			for {
				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == Max-1 {
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}
					fmt.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}
	select {
	case <-time.After(time.Hour):
	}
}
