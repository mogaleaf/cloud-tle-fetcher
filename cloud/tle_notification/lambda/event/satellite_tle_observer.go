package event

import (
	"sync"

	"github.com/google/uuid"
)

type SatelliteTleObserver interface {
	Register(satelliteNames []string) (uuid string, com <-chan NewTleEvent)
	Unregister(name string)

	NewTle(satellite string, tle string)
}

type satelliteTleObserver struct {
	mutex       sync.Mutex
	listeners   map[string]*listener
	satNameUids map[string][]*listener
}

type listener struct {
	name       string
	satellites []string
	connection chan NewTleEvent
}

type NewTleEvent struct {
	SatelliteName string
	Tle           string
}

func NewSatelliteTleObserver() SatelliteTleObserver {
	return &satelliteTleObserver{
		listeners:   make(map[string]*listener),
		satNameUids: make(map[string][]*listener),
	}
}

func (s *satelliteTleObserver) Register(satelliteNames []string) (string, <-chan NewTleEvent) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	newString := uuid.NewString()
	com := make(chan NewTleEvent, 10)
	l := &listener{
		name:       newString,
		satellites: satelliteNames,
		connection: com,
	}
	s.listeners[newString] = l
	for _, sat := range satelliteNames {
		s.satNameUids[sat] = append(s.satNameUids[sat], l)
	}
	return newString, com
}

func (s *satelliteTleObserver) Unregister(name string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	listner, ok := s.listeners[name]
	if !ok {
		return
	}
	close(listner.connection)
	delete(s.listeners, name)
}

func (s *satelliteTleObserver) NewTle(satellite string, tle string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for satName, listners := range s.satNameUids {
		if satName == satellite {
			for _, listner := range listners {
				listner.connection <- NewTleEvent{
					SatelliteName: satellite,
					Tle:           tle,
				}
			}
		}
	}
}
