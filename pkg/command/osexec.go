package command

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
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

	{
		if _, err := io.Copy(os.Stdout, &e.out); err != nil {
			log.Print(err)
		}
		if _, err := io.Copy(os.Stderr, &e.err); err != nil {
			log.Print(err)
		}
	}

	return out, err
}
