package forwarding

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/rollicks-c/kgate/internal/config"
	"github.com/rollicks-c/kgate/internal/logic/model"
	"io"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
	"net/http"
	"time"
)

type managedForwarder struct {
	group                  config.PortGroup
	serviceName, namespace string
	localPort, remotePort  string
	readyCh                chan struct{}
	timeout                time.Duration
}

type portForwarder interface {
	ForwardPorts() error
}

func (m managedForwarder) run(c model.Controller, pf portForwarder) {

	hasError := false

	// start session
	go func() {

		// setup cleanup
		defer c.StopWaitGroup().Done()
		c.StopWaitGroup().Add(1)

		// start port-forwarding
		err := pf.ForwardPorts()

		// handle errors
		if err != nil {
			c.UpdateProcess(
				m,
				model.Failure,
				err.Error(),
			)
			hasError = true
		} else {
			c.UpdateProcess(
				m,
				model.Stopped,
				"",
			)
		}
	}()

	// await readiness or timeout
	select {
	case <-m.readyCh:
		c.UpdateProcess(
			m,
			model.Running,
			"",
		)
	case <-time.After(m.timeout):
		if !hasError {
			c.UpdateProcess(
				m,
				model.Failure,
				fmt.Sprintf("timeout occured (%s)", m.timeout.String()),
			)
		}
	}

}

func (m managedForwarder) hash() string {
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%v:%s:%s:%s:%s", m.group, m.serviceName, m.namespace, m.localPort, m.remotePort)))
	return hex.EncodeToString(hash.Sum(nil))
}

func (m managedForwarder) createPortForwarder(stopChan chan struct{}) (*portforward.PortForwarder, error) {

	// gather pod
	client, conf, err := createClient(m.group.Target)
	if err != nil {
		return nil, err
	}
	podName, err := getPodForService(client, m.serviceName, m.namespace)
	if err != nil {
		return nil, fmt.Errorf("failed to get pod for service %s: %v", m.serviceName, err)
	}

	// prep port-forwarding request
	url := client.CoreV1().RESTClient().Post().
		Resource("pods").
		Namespace(m.namespace).
		Name(podName).
		SubResource("portforward").
		URL()
	transport, upgrader, err := spdy.RoundTripperFor(conf)
	if err != nil {
		return nil, fmt.Errorf("failed to create round-tripper: %v", err)
	}

	// create port forward
	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: transport}, http.MethodPost, url)
	ports := []string{fmt.Sprintf("%s:%s", m.localPort, m.remotePort)}
	pf, err := portforward.New(dialer, ports, stopChan, m.readyCh, io.Discard, io.Discard)
	if err != nil {
		return nil, fmt.Errorf("failed to create port-forward: %v", err)
	}

	// port-forwarding ready
	return pf, nil

}
