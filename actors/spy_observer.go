package actors

type SpyObserver struct {
	ch chan struct{}
}

func (o *SpyObserver) Update() {
	o.ch <- struct{}{}
}
