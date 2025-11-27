package iface

type World interface {
	AddObject(object Object) bool
	RemoveObject(id int) bool
	GetRemovals() []int
	GetObjects() map[int]Object
	Size() int
	NewId() int
	Cleanup()
}
