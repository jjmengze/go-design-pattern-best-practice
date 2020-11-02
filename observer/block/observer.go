package main

import (
	"errors"
	"sync"
)

//implement subject can notify those who observer
type subject interface {
	Attach(name string, observer observer) error
	Detach(name string, observer observer)
	Notify(body *info) error
}

//implement observer object  can receive those subject notice
type observer interface {
	Update(body *info) error
}

//default mail implement subject interface, if you don't want to implement subject you cloud reuse this object
//observer field could add the user's name in the map and store their observer object.
type defaultMail struct {
	observer map[string]observer
	sync.Mutex
}

func NewDefaultMail() *defaultMail {
	return &defaultMail{observer: map[string]observer{}}
}

func (d *defaultMail) Attach(name string, observer observer) error {
	d.Lock()
	defer d.Unlock()
	if _, ok := d.observer[name]; !ok {
		d.observer[name] = observer
		return nil
	}
	return errors.New("already add user")
}

func (d *defaultMail) Detach(name string, observer observer) {
	d.Lock()
	defer d.Unlock()
	delete(d.observer, name)
}

//Notify all observer should change their status
func (d defaultMail) Notify(body *info) error {
	for _, observer := range d.observer {
		observer.Update(body)
	}
	return nil
}
