package master

type Config struct {
	MapperCount  int
	ReducerCount int
}

type Master struct{}

func (m *Master) Run() {
	// configure & start filesystems

	// start mappers

	// wait for mappers to finish

	// start reducers

	// wait for reducers to finish

	// return statistical result
}
