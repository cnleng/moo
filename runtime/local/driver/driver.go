package driver

import (
	"io"

	"github.com/moobu/moo/builder"
)

type Driver interface {
	// Fork creates a new process
	Fork(*Runnable) (*Process, error)
	// Kill kills a process
	Kill(*Process) error
	// Wait waits a process to exit
	Wait(*Process) error
}

type Runnable struct {
	Bundle *builder.Bundle
	Env    []string
	Args   []string
}

type Process struct {
	ID     int
	Stdin  io.Writer
	Stdout io.Reader
	Stderr io.Reader
}
