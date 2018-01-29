package model

//Galinha é portuguesa
type Galinha interface {
	Cacareja() string //método cacarejar
}

//Pato é tosco!
type Pato interface {
	Grasna() string //método fazer barulho
}

//Ave é uma ave que jamais será um pombo!
type Ave struct {
	Nome string
}

//Cacareja retorna o som que o animal português faz
func (a Ave) Cacareja() string {
	return "có!"
}

//Grasna retorna o barulho do animal das trevas
func (a Ave) Grasna() string {
	return "quack!!"
}
