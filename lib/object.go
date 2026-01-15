package lib

type Object struct {
	Name string
}

func New() *Object {
	return &Object{}
}
