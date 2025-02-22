package forwarding

import (
	"fmt"
	"github.com/rollicks-c/kgate/internal/config"
	"github.com/rollicks-c/kgate/internal/logic/model"
	"time"
)

func CreateForwarder(group config.PortGroup, def config.PortForward) model.Process {

	return &managedForwarder{
		group:       group,
		serviceName: def.Service,
		namespace:   def.Namespace,
		localPort:   def.LocalPort,
		remotePort:  def.RemotePort,
		readyCh:     make(chan struct{}),
		timeout:     time.Second * 5,
	}

}

func (m managedForwarder) ID() string {
	return m.hash()
}

func (m managedForwarder) Group() string {
	return m.group.Name
}

func (m managedForwarder) Describe() string {
	return fmt.Sprintf("%s:%s/%s:%s", m.localPort, m.namespace, m.serviceName, m.remotePort)
}

func (m managedForwarder) Run(c model.Controller) {

	// setup port-forwarder
	pf, err := m.createPortForwarder(c.StopChannel())
	if err != nil {
		c.UpdateProcess(m, model.Failure, err.Error())
		return
	}

	m.run(c, pf)

}
