// Code generated by ecs https://github.com/gabstv/ecs; DO NOT EDIT.

package hello

import (
    "sort"
    

    "github.com/gabstv/ecs/v2"
)








const uuidMoveComponent = "94091420-5E23-457E-8F8D-A422A03E36AF"
const capMoveComponent = 256

type drawerMoveComponent struct {
    Entity ecs.Entity
    Data   Move
}

// WatchMove is a helper struct to access a valid pointer of Move
type WatchMove interface {
    Entity() ecs.Entity
    Data() *Move
}

type slcdrawerMoveComponent []drawerMoveComponent
func (a slcdrawerMoveComponent) Len() int           { return len(a) }
func (a slcdrawerMoveComponent) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a slcdrawerMoveComponent) Less(i, j int) bool { return a[i].Entity < a[j].Entity }


type mWatchMove struct {
    c *MoveComponent
    entity ecs.Entity
}

func (w *mWatchMove) Entity() ecs.Entity {
    return w.entity
}

func (w *mWatchMove) Data() *Move {
    
    
    id := w.c.indexof(w.entity)
    if id == -1 {
        return nil
    }
    return &w.c.data[id].Data
}

// MoveComponent implements ecs.BaseComponent
type MoveComponent struct {
    initialized bool
    flag        ecs.Flag
    world       ecs.BaseWorld
    wkey        [4]byte
    data        []drawerMoveComponent
    
}

// GetMoveComponent returns the instance of the component in a World
func GetMoveComponent(w ecs.BaseWorld) *MoveComponent {
    return w.C(uuidMoveComponent).(*MoveComponent)
}

// SetMoveComponentData updates/adds a Move to Entity e
func SetMoveComponentData(w ecs.BaseWorld, e ecs.Entity, data Move) {
    GetMoveComponent(w).Upsert(e, data)
}

// GetMoveComponentData gets the *Move of Entity e
func GetMoveComponentData(w ecs.BaseWorld, e ecs.Entity) *Move {
    return GetMoveComponent(w).Data(e)
}

// WatchMoveComponentData gets a pointer getter of an entity's Move.
//
// The pointer must not be stored because it may become invalid overtime.
func WatchMoveComponentData(w ecs.BaseWorld, e ecs.Entity) WatchMove {
    return &mWatchMove{
        c: GetMoveComponent(w),
        entity: e,
    }
}

// UUID implements ecs.BaseComponent
func (MoveComponent) UUID() string {
    return "94091420-5E23-457E-8F8D-A422A03E36AF"
}

// Name implements ecs.BaseComponent
func (MoveComponent) Name() string {
    return "MoveComponent"
}

func (c *MoveComponent) indexof(e ecs.Entity) int {
    i := sort.Search(len(c.data), func(i int) bool { return c.data[i].Entity >= e })
    if i < len(c.data) && c.data[i].Entity == e {
        return i
    }
    return -1
}

// Upsert creates or updates a component data of an entity.
// Not recommended to be used directly. Use SetMoveComponentData to change component
// data outside of a system loop.
func (c *MoveComponent) Upsert(e ecs.Entity, data interface{}) {
    v, ok := data.(Move)
    if !ok {
        panic("data must be Move")
    }
    
    id := c.indexof(e)
    
    if id > -1 {
        
        dwr := &c.data[id]
        dwr.Data = v
        
        return
    }
    
    rsz := false
    if cap(c.data) == len(c.data) {
        rsz = true
        c.world.CWillResize(c, c.wkey)
        
    }
    newindex := len(c.data)
    c.data = append(c.data, drawerMoveComponent{
        Entity: e,
        Data:   v,
    })
    if len(c.data) > 1 {
        if c.data[newindex].Entity < c.data[newindex-1].Entity {
            c.world.CWillResize(c, c.wkey)
            
            sort.Sort(slcdrawerMoveComponent(c.data))
            rsz = true
        }
    }
    
    if rsz {
        
        c.world.CResized(c, c.wkey)
        c.world.Dispatch(ecs.Event{
            Type: ecs.EvtComponentsResized,
            ComponentName: "MoveComponent",
            ComponentID: "94091420-5E23-457E-8F8D-A422A03E36AF",
        })
    }
    
    c.world.CAdded(e, c, c.wkey)
    c.world.Dispatch(ecs.Event{
        Type: ecs.EvtComponentAdded,
        ComponentName: "MoveComponent",
        ComponentID: "94091420-5E23-457E-8F8D-A422A03E36AF",
        Entity: e,
    })
}

// Remove a Move data from entity e
//
// Warning: DO NOT call remove inside the system entities loop
func (c *MoveComponent) Remove(e ecs.Entity) {
    
    
    i := c.indexof(e)
    if i == -1 {
        return
    }
    
    //c.data = append(c.data[:i], c.data[i+1:]...)
    c.data = c.data[:i+copy(c.data[i:], c.data[i+1:])]
    c.world.CRemoved(e, c, c.wkey)
    
    c.world.Dispatch(ecs.Event{
        Type: ecs.EvtComponentRemoved,
        ComponentName: "MoveComponent",
        ComponentID: "94091420-5E23-457E-8F8D-A422A03E36AF",
        Entity: e,
    })
}

func (c *MoveComponent) Data(e ecs.Entity) *Move {
    
    
    index := c.indexof(e)
    if index > -1 {
        return &c.data[index].Data
    }
    return nil
}

// Flag returns the 
func (c *MoveComponent) Flag() ecs.Flag {
    return c.flag
}

// Setup is called by ecs.BaseWorld
//
// Do not call this directly
func (c *MoveComponent) Setup(w ecs.BaseWorld, f ecs.Flag, key [4]byte) {
    if c.initialized {
        panic("MoveComponent called Setup() more than once")
    }
    c.flag = f
    c.world = w
    c.wkey = key
    c.data = make([]drawerMoveComponent, 0, 256)
    c.initialized = true
    
}


func init() {
    ecs.RegisterComponent(func() ecs.BaseComponent {
        return &MoveComponent{}
    })
}
