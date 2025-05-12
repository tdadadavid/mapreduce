package mr

import (
	"github.com/tdadadavid/mapreduce/pkg/mappers"
	"github.com/tdadadavid/mapreduce/pkg/master"
	"github.com/tdadadavid/mapreduce/pkg/reducers"
	"github.com/tdadadavid/mapreduce/pkg/statistics"
)

type MR struct {
	config   master.Config
	mappers  *mappers.Mapper
	reducers *reducers.Reducer
	stats    statistics.Stats
}

func New() MR {
	return MR{}
}

func (m *MR) Run() {}

func (m *MR) RegisterMapper(mapp mappers.Mapper) *MR {
	m.mappers = &mapp
	return m
}

func (m *MR) RegisterReducer(r reducers.Reducer) *MR {
	m.reducers = &r
	return m
}

// SetMapCount sets the number of mappers to run
func (m *MR) SetMapCount(count int) *MR {
	m.config.MapperCount = count
	return m
}

// SetReduceCount sets the number of reducers to run
func (m *MR) SetReduceCount(count int) *MR {
	m.config.ReducerCount = count
	return m
}
