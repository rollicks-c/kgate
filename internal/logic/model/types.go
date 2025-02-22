package model

import (
	"golang.org/x/net/context"
	"sync"
)

type Status int

const (
	Running Status = iota + 1
	Stopped        // Not used
	Failure
	Restart
)

type Update struct {
	ID          string
	SortIndex   int
	Group       string
	PortForward string
	Status      Status
	Message     string
}

type Process interface {
	ID() string
	Group() string
	Describe() string
	Run(c Controller)
}

type Controller interface {
	StopChannel() chan struct{}
	StopWaitGroup() *sync.WaitGroup
	UpdateProcess(proc Process, state Status, msg string)
	TogglePause()
	Quit()
}
type Frontend interface {
	Run(controller Controller)
	Stop()
	Update(update Update)
	ShowMessage(msg string)
}

type Context struct {
	context.Context
	WG *sync.WaitGroup
}
