package cmd

import (
	"os/exec"
	"testing"
)

func TestCmd(t *testing.T) {
	s := "ffmpeg -version"
	cmd := exec.Command("/bin/bash", "-c", s)
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(output)
	}
	println(string(output))
}
