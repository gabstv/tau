// Code generated by ecs https://github.com/gabstv/ecs; DO NOT EDIT.

package core

import (
    
    "sort"

    "github.com/gabstv/ecs/v2"
    
)









const uuidTrTweeningSystem = "820C75AB-CAD6-47AE-A84C-1EC7BAECE328"

type viewTrTweeningSystem struct {
    entities []VITrTweeningSystem
    world ecs.BaseWorld
    
}

type VITrTweeningSystem struct {
    Entity ecs.Entity
    
    Transform *Transform 
    
    TrTweening *TrTweening 
    
}

type sortedVITrTweeningSystems []VITrTweeningSystem
func (a sortedVITrTweeningSystems) Len() int           { return len(a) }
func (a sortedVITrTweeningSystems) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortedVITrTweeningSystems) Less(i, j int) bool { return a[i].Entity < a[j].Entity }

func newviewTrTweeningSystem(w ecs.BaseWorld) *viewTrTweeningSystem {
    return &viewTrTweeningSystem{
        entities: make([]VITrTweeningSystem, 0),
        world: w,
    }
}

func (v *viewTrTweeningSystem) Matches() []VITrTweeningSystem {
    
    return v.entities
    
}

func (v *viewTrTweeningSystem) indexof(e ecs.Entity) int {
    i := sort.Search(len(v.entities), func(i int) bool { return v.entities[i].Entity >= e })
    if i < len(v.entities) && v.entities[i].Entity == e {
        return i
    }
    return -1
}

// Fetch a specific entity
func (v *viewTrTweeningSystem) Fetch(e ecs.Entity) (data VITrTweeningSystem, ok bool) {
    
    i := v.indexof(e)
    if i == -1 {
        return VITrTweeningSystem{}, false
    }
    return v.entities[i], true
}

func (v *viewTrTweeningSystem) Add(e ecs.Entity) bool {
    
    
    // MUST NOT add an Entity twice:
    if i := v.indexof(e); i > -1 {
        return false
    }
    v.entities = append(v.entities, VITrTweeningSystem{
        Entity: e,
        Transform: GetTransformComponent(v.world).Data(e),
TrTweening: GetTrTweeningComponent(v.world).Data(e),

    })
    if len(v.entities) > 1 {
        if v.entities[len(v.entities)-1].Entity < v.entities[len(v.entities)-2].Entity {
            sort.Sort(sortedVITrTweeningSystems(v.entities))
        }
    }
    return true
}

func (v *viewTrTweeningSystem) Remove(e ecs.Entity) bool {
    
    
    if i := v.indexof(e); i != -1 {

        v.entities = append(v.entities[:i], v.entities[i+1:]...)
        return true
    }
    return false
}

func (v *viewTrTweeningSystem) clearpointers() {
    
    
    for i := range v.entities {
        e := v.entities[i].Entity
        
        v.entities[i].Transform = nil
        
        v.entities[i].TrTweening = nil
        
        _ = e
    }
}

func (v *viewTrTweeningSystem) rescan() {
    
    
    for i := range v.entities {
        e := v.entities[i].Entity
        
        v.entities[i].Transform = GetTransformComponent(v.world).Data(e)
        
        v.entities[i].TrTweening = GetTrTweeningComponent(v.world).Data(e)
        
        _ = e
        
    }
}

// TrTweeningSystem implements ecs.BaseSystem
type TrTweeningSystem struct {
    initialized bool
    world       ecs.BaseWorld
    view        *viewTrTweeningSystem
    enabled     bool
    
}

// GetTrTweeningSystem returns the instance of the system in a World
func GetTrTweeningSystem(w ecs.BaseWorld) *TrTweeningSystem {
    return w.S(uuidTrTweeningSystem).(*TrTweeningSystem)
}

// Enable system
func (s *TrTweeningSystem) Enable() {
    s.enabled = true
}

// Disable system
func (s *TrTweeningSystem) Disable() {
    s.enabled = false
}

// Enabled checks if enabled
func (s *TrTweeningSystem) Enabled() bool {
    return s.enabled
}

// UUID implements ecs.BaseSystem
func (TrTweeningSystem) UUID() string {
    return "820C75AB-CAD6-47AE-A84C-1EC7BAECE328"
}

func (TrTweeningSystem) Name() string {
    return "TrTweeningSystem"
}

// ensure matchfn
var _ ecs.MatchFn = matchTrTweeningSystem

// ensure resizematchfn
var _ ecs.MatchFn = resizematchTrTweeningSystem

func (s *TrTweeningSystem) match(eflag ecs.Flag) bool {
    return matchTrTweeningSystem(eflag, s.world)
}

func (s *TrTweeningSystem) resizematch(eflag ecs.Flag) bool {
    return resizematchTrTweeningSystem(eflag, s.world)
}

func (s *TrTweeningSystem) ComponentAdded(e ecs.Entity, eflag ecs.Flag) {
    if s.match(eflag) {
        if s.view.Add(e) {
            // TODO: dispatch event that this entity was added to this system
            
        }
    } else {
        if s.view.Remove(e) {
            // TODO: dispatch event that this entity was removed from this system
            
        }
    }
}

func (s *TrTweeningSystem) ComponentRemoved(e ecs.Entity, eflag ecs.Flag) {
    if s.match(eflag) {
        if s.view.Add(e) {
            // TODO: dispatch event that this entity was added to this system
            
        }
    } else {
        if s.view.Remove(e) {
            // TODO: dispatch event that this entity was removed from this system
            
        }
    }
}

func (s *TrTweeningSystem) ComponentResized(cflag ecs.Flag) {
    if s.resizematch(cflag) {
        s.view.rescan()
        
    }
}

func (s *TrTweeningSystem) ComponentWillResize(cflag ecs.Flag) {
    if s.resizematch(cflag) {
        
        s.view.clearpointers()
    }
}

func (s *TrTweeningSystem) V() *viewTrTweeningSystem {
    return s.view
}

func (*TrTweeningSystem) Priority() int64 {
    return 90
}

func (s *TrTweeningSystem) Setup(w ecs.BaseWorld) {
    if s.initialized {
        panic("TrTweeningSystem called Setup() more than once")
    }
    s.view = newviewTrTweeningSystem(w)
    s.world = w
    s.enabled = true
    s.initialized = true
    
}


func init() {
    ecs.RegisterSystem(func() ecs.BaseSystem {
        return &TrTweeningSystem{}
    })
}
