package memory

type Memory struct{}

func NewMemory(size int) *Memory {
	return &Memory{}
}

func (m *Memory) Add(obj any) {}
