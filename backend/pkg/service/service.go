package service

// For fx
type Name string

func ServiceName(name string) Name {
	return Name(name)
}
