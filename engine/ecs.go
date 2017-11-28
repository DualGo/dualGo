package engine

type Entity struct{
	id int32
	components []Components
}

func (entity *Entity) AddComponent(component Components){
	entity.components = append(entity.components, component)
}

func (entity Entity) GetComponent() *[]Components{
	return &entity.components
}

type ComponentAction func()

type Components struct{
	action ComponentAction
}
