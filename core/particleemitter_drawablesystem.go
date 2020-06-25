// Code generated by ecs https://github.com/gabstv/ecs; DO NOT EDIT.

package core

import (
    
    "sort"

    "github.com/gabstv/ecs/v2"
    
)









const uuidDrawableParticleEmitterSystem = "627C4B36-EE45-40C6-91AE-617D5CFDD8FC"

type viewDrawableParticleEmitterSystem struct {
    entities []VIDrawableParticleEmitterSystem
    world ecs.BaseWorld
    
}

type VIDrawableParticleEmitterSystem struct {
    Entity ecs.Entity
    
    Drawable *Drawable 
    
    ParticleEmitter *ParticleEmitter 
    
}

type sortedVIDrawableParticleEmitterSystems []VIDrawableParticleEmitterSystem
func (a sortedVIDrawableParticleEmitterSystems) Len() int           { return len(a) }
func (a sortedVIDrawableParticleEmitterSystems) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortedVIDrawableParticleEmitterSystems) Less(i, j int) bool { return a[i].Entity < a[j].Entity }

func newviewDrawableParticleEmitterSystem(w ecs.BaseWorld) *viewDrawableParticleEmitterSystem {
    return &viewDrawableParticleEmitterSystem{
        entities: make([]VIDrawableParticleEmitterSystem, 0),
        world: w,
    }
}

func (v *viewDrawableParticleEmitterSystem) Matches() []VIDrawableParticleEmitterSystem {
    
    return v.entities
    
}

func (v *viewDrawableParticleEmitterSystem) indexof(e ecs.Entity) int {
    i := sort.Search(len(v.entities), func(i int) bool { return v.entities[i].Entity >= e })
    if i < len(v.entities) && v.entities[i].Entity == e {
        return i
    }
    return -1
}

// Fetch a specific entity
func (v *viewDrawableParticleEmitterSystem) Fetch(e ecs.Entity) (data VIDrawableParticleEmitterSystem, ok bool) {
    
    i := v.indexof(e)
    if i == -1 {
        return VIDrawableParticleEmitterSystem{}, false
    }
    return v.entities[i], true
}

func (v *viewDrawableParticleEmitterSystem) Add(e ecs.Entity) bool {
    
    
    // MUST NOT add an Entity twice:
    if i := v.indexof(e); i > -1 {
        return false
    }
    v.entities = append(v.entities, VIDrawableParticleEmitterSystem{
        Entity: e,
        Drawable: GetDrawableComponent(v.world).Data(e),
ParticleEmitter: GetParticleEmitterComponent(v.world).Data(e),

    })
    if len(v.entities) > 1 {
        if v.entities[len(v.entities)-1].Entity < v.entities[len(v.entities)-2].Entity {
            sort.Sort(sortedVIDrawableParticleEmitterSystems(v.entities))
        }
    }
    return true
}

func (v *viewDrawableParticleEmitterSystem) Remove(e ecs.Entity) bool {
    
    
    if i := v.indexof(e); i != -1 {

        v.entities = append(v.entities[:i], v.entities[i+1:]...)
        return true
    }
    return false
}

func (v *viewDrawableParticleEmitterSystem) clearpointers() {
    
    
    for i := range v.entities {
        e := v.entities[i].Entity
        
        v.entities[i].Drawable = nil
        
        v.entities[i].ParticleEmitter = nil
        
        _ = e
    }
}

func (v *viewDrawableParticleEmitterSystem) rescan() {
    
    
    for i := range v.entities {
        e := v.entities[i].Entity
        
        v.entities[i].Drawable = GetDrawableComponent(v.world).Data(e)
        
        v.entities[i].ParticleEmitter = GetParticleEmitterComponent(v.world).Data(e)
        
        _ = e
        
    }
}

// DrawableParticleEmitterSystem implements ecs.BaseSystem
type DrawableParticleEmitterSystem struct {
    initialized bool
    world       ecs.BaseWorld
    view        *viewDrawableParticleEmitterSystem
    enabled     bool
    
}

// GetDrawableParticleEmitterSystem returns the instance of the system in a World
func GetDrawableParticleEmitterSystem(w ecs.BaseWorld) *DrawableParticleEmitterSystem {
    return w.S(uuidDrawableParticleEmitterSystem).(*DrawableParticleEmitterSystem)
}

// Enable system
func (s *DrawableParticleEmitterSystem) Enable() {
    s.enabled = true
}

// Disable system
func (s *DrawableParticleEmitterSystem) Disable() {
    s.enabled = false
}

// Enabled checks if enabled
func (s *DrawableParticleEmitterSystem) Enabled() bool {
    return s.enabled
}

// UUID implements ecs.BaseSystem
func (DrawableParticleEmitterSystem) UUID() string {
    return "627C4B36-EE45-40C6-91AE-617D5CFDD8FC"
}

func (DrawableParticleEmitterSystem) Name() string {
    return "DrawableParticleEmitterSystem"
}

// ensure matchfn
var _ ecs.MatchFn = matchDrawableParticleEmitterSystem

// ensure resizematchfn
var _ ecs.MatchFn = resizematchDrawableParticleEmitterSystem

func (s *DrawableParticleEmitterSystem) match(eflag ecs.Flag) bool {
    return matchDrawableParticleEmitterSystem(eflag, s.world)
}

func (s *DrawableParticleEmitterSystem) resizematch(eflag ecs.Flag) bool {
    return resizematchDrawableParticleEmitterSystem(eflag, s.world)
}

func (s *DrawableParticleEmitterSystem) ComponentAdded(e ecs.Entity, eflag ecs.Flag) {
    if s.match(eflag) {
        if s.view.Add(e) {
            // TODO: dispatch event that this entity was added to this system
            s.onEntityAdded(e)
        }
    } else {
        if s.view.Remove(e) {
            // TODO: dispatch event that this entity was removed from this system
            s.onEntityRemoved(e)
        }
    }
}

func (s *DrawableParticleEmitterSystem) ComponentRemoved(e ecs.Entity, eflag ecs.Flag) {
    if s.match(eflag) {
        if s.view.Add(e) {
            // TODO: dispatch event that this entity was added to this system
            s.onEntityAdded(e)
        }
    } else {
        if s.view.Remove(e) {
            // TODO: dispatch event that this entity was removed from this system
            s.onEntityRemoved(e)
        }
    }
}

func (s *DrawableParticleEmitterSystem) ComponentResized(cflag ecs.Flag) {
    if s.resizematch(cflag) {
        s.view.rescan()
        s.onResize()
    }
}

func (s *DrawableParticleEmitterSystem) ComponentWillResize(cflag ecs.Flag) {
    if s.resizematch(cflag) {
        s.onWillResize()
        s.view.clearpointers()
    }
}

func (s *DrawableParticleEmitterSystem) V() *viewDrawableParticleEmitterSystem {
    return s.view
}

func (*DrawableParticleEmitterSystem) Priority() int64 {
    return 10
}

func (s *DrawableParticleEmitterSystem) Setup(w ecs.BaseWorld) {
    if s.initialized {
        panic("DrawableParticleEmitterSystem called Setup() more than once")
    }
    s.view = newviewDrawableParticleEmitterSystem(w)
    s.world = w
    s.enabled = true
    s.initialized = true
    
}


func init() {
    ecs.RegisterSystem(func() ecs.BaseSystem {
        return &DrawableParticleEmitterSystem{}
    })
}
