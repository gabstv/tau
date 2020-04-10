package tau

// Archetype is a recipe to create entities with a preset of components.
type Archetype struct {
	World      *World
	Components []*Component
}

// NewArchetype returns a new archetype. This func unsures that no
// duplicated components are added.
func NewArchetype(world *World, comps ...*Component) *Archetype {
	cmap := make(map[*Component]bool)
	components := make([]*Component, 0, len(comps))
	for _, c := range comps {
		if cmap[c] {
			// duplicated component
			continue
		}
		components = append(components, c)
		cmap[c] = true
	}
	arch := &Archetype{
		World:      world,
		Components: components,
	}
	return arch
}

// NewEntity adds a new entity with the component data to the world
//
// The best way to create an archetype entity is to ensure that the
// component data follows the same order that the components were
// created in NewArchetype
func (a *Archetype) NewEntity(compdata ...interface{}) Entity {
	entity := a.World.NewEntity()
	compds := clonecompslc(a.Components)
	for _, cdata := range compdata {
		for i, c := range compds {
			if c.Validate(cdata) {
				if err := a.World.AddComponentToEntity(entity, c, cdata); err != nil {
					// this should never happen
					panic(err)
				}
				compds = append(compds[:i], compds[i+1:]...)
				break
			}
		}
	}
	return entity
}

func clonecompslc(v []*Component) []*Component {
	x := make([]*Component, len(v), len(v))
	copy(x, v)
	return x
}
