package engine

type initCallback func()
type updateCallback func()

var(

	systems []*System
)
func initEngine(){

}

func loopEngine(delta float32){
	for _, s := range systems{
		s.updateSystem(delta)
	}
}

func AddSystem(system *System){
	system.initSystem()
	systems = append(systems, system)
}

func RemoveSystem(system *System){
	for i, s := range systems{
		if s == system{
			systems = append(systems[:i], systems[i+1:]...)
			break
		}
	}
}

