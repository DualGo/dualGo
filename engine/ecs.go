package engine

import (
	"sync/atomic"
)

var(
	idInc uint64
)

//identity struct

type Entity struct{
	id uint64
	components []Components
}

func (entity *Entity) initEntity(){

}

func (entity *Entity) updateEntity(){
	for _, component := range entity.components{
		component.action()
	}
}

func NewEntity() Entity{
	return Entity{id: atomic.AddUint64(&idInc, 1)}
}

func (entity *Entity) AddComponent(component Components){
	entity.components = append(entity.components, component)
}

func (entity Entity) GetComponents() *[]Components{
	return &entity.components
}

//component struct

type ComponentAction func()

type Components struct{
	action ComponentAction
}

//system struct
type System struct{
	entities []*Entity
	Init initCallback
	Update updateCallback
}

func (system *System) AddEntity(entity *Entity){
	entity.initEntity()
	system.entities = append(system.entities, entity)
}

func (system *System) RemoveEntity(entity *Entity){
	for i, e := range system.entities{
		if entity.id ==  e.id{
			system.entities = append(system.entities[:i], system.entities[i+1:]...)
			break
		}
	}
}

func (system *System) initSystem(){
	system.Init()
}

func (system System) updateSystem(delta float32){
	system.Update()
	for _, entity := range system.entities{
		entity.updateEntity()

	}
}
