package tau

import (
	"testing"

	"github.com/gabstv/ecs"
	"github.com/stretchr/testify/assert"
)

type compAData struct {
	X int
	Y int
}
type compBData struct {
	Name string
}

func TestNewArchetype(t *testing.T) {
	w := NewWorld(&Engine{})

	c1, err := w.NewComponent(ecs.NewComponentInput{
		Name: "COMP_A",
		ValidateDataFn: func(data interface{}) bool {
			_, ok := data.(*compAData)
			return ok
		},
	})
	assert.NoError(t, err)
	c2, err := w.NewComponent(ecs.NewComponentInput{
		Name: "COMP_B",
		ValidateDataFn: func(data interface{}) bool {
			_, ok := data.(*compBData)
			return ok
		},
	})
	assert.NoError(t, err)
	arche1 := NewArchetype(w, c1, c2)

	w.NewSystem("", 1, func(ctx ecs.Context) {
		m := ctx.System().View().Matches()
		for _, v := range m {
			da := v.Components[c1].(*compAData)
			db := v.Components[c2].(*compBData)
			if db.Name == "tau" {
				da.X++
			} else {
				da.Y++
			}
		}
	}, c1, c2)

	// most optimal way to instantiate an archetype is to follow the order
	// of the components specified by NewArchetype
	e1 := arche1.NewEntity(&compAData{
		X: 10,
		Y: 20,
	}, &compBData{
		Name: "tau",
	})
	e2 := arche1.NewEntity(&compAData{
		X: 10,
		Y: 20,
	}, &compBData{
		Name: "Trends",
	})
	w.Run(1)

	ed1 := c1.Data(e1).(*compAData)
	ed2 := c1.Data(e2).(*compAData)

	assert.Equal(t, 11, ed1.X)
	assert.Equal(t, 20, ed1.Y)
	assert.Equal(t, 10, ed2.X)
	assert.Equal(t, 21, ed2.Y)
}
