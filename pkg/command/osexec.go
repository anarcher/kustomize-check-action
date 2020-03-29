package command

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	au "github.com/logrusorgru/aurora"
)

type OSExec struct {
	out bytes.Buffer
	err bytes.Buffer
}

func NewOSExec() *OSExec {
	e := &OSExec{}
	return e
}

func (e *OSExec) Run(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = &e.out
	cmd.Stderr = &e.err

	err := cmd.Run()
	out := e.out.String()

	fmt.Println(au.Cyan("CMD:"), au.Gray(8-1, fmt.Sprintf("%s %s", name, strings.Join(args, " "))))
	fmt.Println(au.Gray(8-1, e.out.String()).BgGray(4 - 1))
	fmt.Println(au.Red(e.err.String()))

	e.out.Reset()
	e.err.Reset()

	return out, err
}
