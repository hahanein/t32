// Package observer merely contains interfaces for the implementation of the
// observer pattern in which a subject maintains a list of so called observers
// and notifies them automatically of any state changes.

package observer

type Subject interface {
	Attach(Observer)
	Notify()
}

type Observer interface {
	Update()
}
