package internal

type Arg struct {
	Name string
	Type string
}

type Return struct {
	Type string
}

type MethodDefinition struct {
	Name    string
	Args    []*Arg
	Returns []*Return
}

type Definition struct {
	Interface string
	Methods   []*MethodDefinition
	Implement string
}
