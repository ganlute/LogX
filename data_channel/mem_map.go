package data_channel

import (
	"errors"
	"github.com/ganlute/LogX/utils"
	"sync"
)

const (
	Record_Not_Found = "Record Not Found"
)

var memDataMap sync.Map

type MemMap struct {
}

func (m *MemMap) InitContext() (err error) {
	// do nothing
	return nil
}

func (m *MemMap) GetContext() (ctx interface{}, err error) {
	key := utils.GoroutineId()
	if value, ok := memDataMap.Load(key); ok {
		return value, nil
	}
	return nil, errors.New(Record_Not_Found)
}
func (m *MemMap) SetContext(ctx interface{}) (err error) {
	key := utils.GoroutineId()
	memDataMap.Store(key, ctx)
	return nil
}
func (m *MemMap) DeleteContext() (err error) {
	key := utils.GoroutineId()
	memDataMap.Delete(key)
	return nil
}
func (m *MemMap) UpdateContext(ctx interface{}) (err error) {
	return m.SetContext(ctx)
}
