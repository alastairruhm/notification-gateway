package mongodb

import "gopkg.in/mgo.v2"
import "time"

var service Service

// Service ...
type Service struct {
	baseSession *mgo.Session
	queue       chan int
	URL         string
	Open        int
}

// New ...
func (s *Service) New() error {
	var err error
	s.queue = make(chan int, MaxPool)
	for i := 0; i < MaxPool; i = i + 1 {
		s.queue <- 1
	}
	s.Open = 0
	s.baseSession, err = mgo.DialWithTimeout(s.URL, 5*time.Second)
	return err
}

// Session ...
func (s *Service) Session() *mgo.Session {
	<-s.queue
	s.Open++
	return s.baseSession.Copy()
}

// Close ...
func (s *Service) Close(c *Collection) {
	c.db.s.Close()
	s.queue <- 1
	s.Open--
}
