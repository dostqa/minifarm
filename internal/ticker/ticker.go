package ticker

var DefaultTicker *Ticker

func init() {
	DefaultTicker = &Ticker{}
}

const (
	ticksPerSecond = 60
	ticksPerFrame  = 15
)

type Ticker struct {
	ticksSinceLaunch int
}

func (t *Ticker) Update() {
	t.ticksSinceLaunch++
}

func (t *Ticker) NowFrame() int {
	return (t.ticksSinceLaunch % ticksPerSecond) / ticksPerFrame
}
