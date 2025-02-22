package gate

import (
	"github.com/rollicks-c/kgate/internal/logic/model"
	"sort"
	"sync"
)

type session struct {

	// data
	processes   map[string]model.Process
	oridnalView map[string]int
	isStopped   bool

	// runtime
	stopChan chan struct{}
	wg       *sync.WaitGroup
}

func newSession() *session {
	s := &session{}
	s.reset()
	return s
}

func (s *session) reset() {
	s.isStopped = false
	s.stopChan = make(chan struct{})
	s.processes = make(map[string]model.Process)
	s.oridnalView = make(map[string]int)
	s.wg = &sync.WaitGroup{}
}

func (s *session) addProcess(proc model.Process) {
	s.processes[proc.ID()] = proc
	s.buildOrdinalView()
}

func (s *session) buildOrdinalView() {
	list := make([]model.Process, 0, len(s.processes))
	for _, p := range s.processes {
		list = append(list, p)
	}
	sort.Slice(list, func(i, j int) bool {
		if list[i].Group() == list[j].Group() {
			return list[i].Describe() < list[j].Describe()
		}
		return list[i].Group() < list[j].Group()
	})
	for i, p := range list {
		s.oridnalView[p.ID()] = i
	}
}
