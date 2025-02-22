package gate

import (
	"context"
	"github.com/rollicks-c/kgate/internal/config"
	"github.com/rollicks-c/kgate/internal/logic/forwarding"
	"github.com/rollicks-c/kgate/internal/logic/model"
	"github.com/rollicks-c/kgate/internal/logic/ui"
	"github.com/rollicks-c/term"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	defaultMessage = "press [yellow::b]q[-::-] to quit | [yellow::b]s[-::-] to stop/resume"
)

type controller struct {

	// data
	groups  []config.PortGroup
	session *session

	// rendering
	view       model.Frontend
	statusChan chan model.Update

	// runtime
	ctx     model.Context
	ctxStop context.CancelFunc
}

func (c controller) StopChannel() chan struct{} {
	return c.session.stopChan
}

func (c controller) StopWaitGroup() *sync.WaitGroup {
	return c.session.wg
}

func newController(groups ...config.PortGroup) *controller {

	// create context
	runContext, ctxStop := context.WithCancel(context.Background())
	ctx := model.Context{
		Context: runContext,
		WG:      &sync.WaitGroup{},
	}

	// create controller
	c := &controller{
		groups:     groups,
		ctx:        ctx,
		ctxStop:    ctxStop,
		statusChan: make(chan model.Update),
	}
	c.session = newSession()
	c.view = ui.NewFancy()

	return c
}

func (c controller) Run() {

	// start frontend
	go c.view.Run(c)
	go c.runUpdateLoop()

	// start forwards
	c.view.ShowMessage("[orange]starting port forwards...")
	c.startSession()

	// listen to terminate signals
	c.view.ShowMessage(defaultMessage)
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigChan:
	case <-c.ctx.Done():
	}

	// controlled shutdown
	c.shutdown()
}

func (c controller) TogglePause() {
	go func() {
		if c.session.isStopped {
			c.session.reset()
			c.startSession()
		} else {
			c.stopSession()
		}
	}()
}

func (c controller) Quit() {
	go func() {
		c.view.ShowMessage("[orange]stopping port forwards...")
		c.stopSession()
		<-time.After(500 * time.Millisecond)
		c.ctxStop()
		_ = syscall.Kill(os.Getpid(), syscall.SIGTERM)
	}()
}

func (c controller) UpdateProcess(proc model.Process, state model.Status, msg string) {
	update := model.Update{
		ID:          proc.ID(),
		SortIndex:   c.session.oridnalView[proc.ID()],
		Group:       proc.Group(),
		PortForward: proc.Describe(),
		Status:      state,
		Message:     msg,
	}
	c.statusChan <- update
}

func (c controller) runUpdateLoop() {
	for {
		select {
		case event := <-c.statusChan:
			c.view.Update(event)
		case <-c.ctx.Done():
			return
		}
	}
}

func (c controller) shutdown() {

	term.Warnf("shutting down app...\n")

	// stop UI
	close(c.statusChan)
	c.view.Stop()

	// stop all routines
	c.ctxStop()
	c.ctx.WG.Wait()
}

func (c controller) startSession() {

	// iterate all port forwards
	for _, g := range c.groups {
		for _, pf := range g.PortForwards {

			// create and run processes
			proc := forwarding.CreateForwarder(g, pf)
			c.session.addProcess(proc)
			go proc.Run(c)
		}
	}
}

func (c controller) stopSession() {

	// session already aborted
	if c.session.isStopped {
		return
	}

	// stop all forwards
	close(c.session.stopChan)
	c.session.wg.Wait()

	// mark session as aborted
	c.session.isStopped = true
}
