package other

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"testing"
)

//2 个代码输出判定题，看有什么问题？notion2024 =
//

func TestM(t *testing.T) {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A= ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B= ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
func parseFile(file string) string {
	f, err := os.Open(file)

	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
}
func TestM2(t *testing.T) {

	parseFile("")

}
