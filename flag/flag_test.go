package flag

import (
	"flag"
	"fmt"
	"testing"
)

var (
	cliName2 string
	cliAge2  int
	cliCanu2 bool
)

//var cliName = flag.String("name", "nick", "Input your name")
//var cliAge = flag.Int("age", 10, "Input your age")
//var cliCanu = flag.Bool("canu", true, "Input canu")
//
//func TestFlag01(t *testing.T) {
//	flag.Parse()
//	flag.Args()
//	flag.NArg()
//	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
//
//	fmt.Println("name=", *cliName)
//	fmt.Println("age=", *cliAge)
//	fmt.Println("canu=", *cliCanu)
//}

func init() {
	flag.StringVar(&cliName2, "name", "nick", "Input your name")
	flag.IntVar(&cliAge2, "age", 10, "Input your age")
	flag.BoolVar(&cliCanu2, "canu", true, "Input canu")

}

func TestFlag02(t *testing.T) {
	flag.Parse()
	flag.Args()
	flag.NArg()
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())

	fmt.Println("name=", cliName2)
	fmt.Println("age=", cliAge2)
	fmt.Println("canu=", cliCanu2)
}
