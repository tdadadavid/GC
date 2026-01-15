package marksweep

import "github.com/tdadadavid/gc/memory"

type GC struct {
	heap memory.Memory
}

func NewGC(memory *memory.Memory) *GC {
	return &GC{}
}

func (gc *GC) Clean() {}
