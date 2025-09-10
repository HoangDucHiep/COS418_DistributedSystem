package main

type Lock struct {
	ch chan bool
}

func NewLock() *Lock {
	lock := &Lock{make(chan bool, 1)}
	lock.ch <- true // initial unlocked state
	return lock
}

func (l *Lock) Lock() {
	<-l.ch
}

func (l *Lock) Unlock() {
	l.ch <- true
}
