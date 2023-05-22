package os

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestExec01(t *testing.T) {
	cmd := exec.Command("bash", "-c", "cat /data/svr/config/svrid")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("err:%v \n", err)
	}
	context := string(output)
	fmt.Println(context)
}

func TestExec02(t *testing.T) {
	cmd := exec.Command("cat", "/data/svr/config/svrid")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("err:%v \n", err)
	}
	context := string(output)
	fmt.Println(context)
}
