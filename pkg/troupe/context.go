package troupe

import (
	"context"
	"time"

	"github.com/gabstv/ecs"
)

type Context interface {
	ecs.Context
	Engine() *Engine
	CurrentFPS() float64
	IsDrawingSkipped() bool
}

type ctxt struct {
	c          context.Context
	dt         float64
	system     *System
	world      Worlder
	engine     *Engine
	fps        float64
	drwskipped bool
}

func (c ctxt) Deadline() (deadline time.Time, ok bool) {
	return c.c.Deadline()
}

func (c ctxt) Done() <-chan struct{} {
	return c.c.Done()
}

func (c ctxt) Err() error {
	return c.c.Err()
}

func (c ctxt) Value(key interface{}) interface{} {
	return c.c.Value(key)
}

func (c ctxt) DT() float64 {
	return c.dt
}

func (c ctxt) System() *System {
	return c.system
}

func (c ctxt) World() Worlder {
	return c.world
}

func (c ctxt) Engine() *Engine {
	return c.engine
}

func (c ctxt) CurrentFPS() float64 {
	return c.fps
}

func (c ctxt) IsDrawingSkipped() bool {
	return c.drwskipped
}
