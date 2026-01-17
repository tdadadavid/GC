package marksweep

import "github.com/tdadadavid/gc/memory"

//NOTE: I found out that the name "heap" given to where objects live is not related to the heap data structure lol 😂

type GC struct {
	heap memory.Memory
}

func NewGC(memory *memory.Memory) *GC {
	return &GC{}
}

func (gc *GC) Clean() {}
