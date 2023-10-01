package animation

import "math/rand"

// TODO: Move into gggg.

type Frames interface {
	// Tick moves the animation forward (in milliseconds).
	Tick(ms int)

	// Frame returns the current frame.
	Frame() int

	// Randomize randomizes the frame.
	Randomize()
}

func NewFrames(maxFrame int, msPerFrame int) Frames {
	return &animation{
		maxFrame:   maxFrame,
		msPerFrame: msPerFrame,
	}
}

type animation struct {
	maxFrame   int
	msPerFrame int

	current int
}

func (anim *animation) Tick(ms int) {
	if anim.maxFrame == 0 {
		return
	}
	anim.current += ms
	if anim.Frame() > anim.maxFrame {
		anim.current -= anim.Frame() * anim.msPerFrame
	}
}

func (anim *animation) Frame() int {
	if anim.maxFrame == 0 {
		return 0
	}
	return anim.current / anim.msPerFrame
}

func (anim *animation) Randomize() {
	if anim.maxFrame == 0 {
		return
	}
	anim.current = rand.Intn(anim.maxFrame * anim.msPerFrame)
}
