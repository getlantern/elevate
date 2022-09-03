package elevate

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/getlantern/byteexec"
)

//go:embed binaries/osx/cocoasudo
var cocoasudo []byte

func buildCommand(prompt string, icon string, name string, args ...string) (*exec.Cmd, error) {
	argsLen := len(args)
	if icon != "" {
		argsLen += 1
	}
	if prompt != "" {
		argsLen += 1
	}
	allArgs := make([]string, 0, argsLen)
	if icon != "" {
		allArgs = append(allArgs, "--icon="+icon)
	}
	if prompt != "" {
		allArgs = append(allArgs, "--prompt="+prompt)
	}
	allArgs = append(allArgs, name)
	allArgs = append(allArgs, args...)
	_, program := filepath.Split(os.Args[0])
	be, err := byteexec.New(cocoasudo, program)
	if err != nil {
		return nil, fmt.Errorf("unable to build byteexec for cocoasudo: %v", err)
	}

	return be.Command(allArgs...), nil
}
