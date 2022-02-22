package observer

import "fmt"

type ISubject interface {
	Add(observer IObserver)
	Del(observer IObserver)
	Notify(msg string)
}

type IObserver interface {
	Update(msg string)
}

type Subject struct {
	observers []IObserver
}

func (s *Subject) Add(observer IObserver) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Del(observer IObserver) {
	for i, ob := range s.observers {
		if ob == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
		}
	}
}

func (s *Subject) Notify(msg string) {
	for _, ob := range s.observers {
		ob.Update(msg)
	}
}


type Observer1 struct {}
func (o Observer1) Update(msg string) {
	fmt.Printf("Observer1: %s\n", msg)
}

type Observer2 struct {}
func (o Observer2) Update(msg string) {
	fmt.Printf("Observer2: %s\n", msg)
}