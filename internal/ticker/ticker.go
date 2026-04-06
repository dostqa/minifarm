package ticker

var DefaultTicker *Ticker

func init() {
	DefaultTicker = &Ticker{ticksPerSecond: 60}
}

type Ticker struct {
	ticksSinceLaunch int
	ticksPerSecond   int
}

func (t *Ticker) Update() {
	t.ticksSinceLaunch++
}

func (t *Ticker) Now() int {
	return t.ticksSinceLaunch
}

func (t *Ticker) NowTick() int {
	return t.ticksSinceLaunch % t.ticksPerSecond
}

func (t *Ticker) TicksPerSecond() int {
	return t.ticksPerSecond
}
