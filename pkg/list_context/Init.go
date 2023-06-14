package list_context

import (
	"context"
	"errors"
	"sync"
)

type DataContext struct {
	Context   *context.Context
	MagnetUrl *string
}

const MaxCountContext = 100

type ListContext struct {
	List     *[MaxCountContext]*DataContext
	LastItem int
	Mutex    sync.Mutex
}

func NewListContext() (listContext *ListContext) {
	listContext = new(ListContext)
	listContext.List = new([MaxCountContext]*DataContext)
	listContext.LastItem = -1
	return
}

// AddContext return istContext.LastItem
func (listContext *ListContext) AddContext(newValueContext *DataContext) (int, error) {
	listContext.Mutex.Lock()
	defer listContext.Mutex.Unlock()

	for i, oldContext := range listContext.List {
		if oldContext == nil {
			continue
		}
		if oldContext.MagnetUrl == nil {
			continue
		}
		if *oldContext.MagnetUrl == *newValueContext.MagnetUrl {
			return i, errors.New("context is busy")
		}
	}

	listContext.LastItem = (listContext.LastItem + 1) % MaxCountContext
	listContext.List[listContext.LastItem] = newValueContext
	return listContext.LastItem, nil
}

func (listContext *ListContext) Delete(delValueContext *DataContext) {
	listContext.Mutex.Lock()
	defer listContext.Mutex.Unlock()

	for i, dataContext := range listContext.List {
		if dataContext == nil {
			continue
		}
		if dataContext.Context == delValueContext.Context {
			listContext.List[i] = nil
			return
		}
	}
	return
}

//func GetContextByKey(listContext []DataContext, key
//} string) *DataContext {
//	for i, context := range listContext {
//		if context.Key == key {
//			return &listContext[i]
//		}
//	}
//	return nil
//}
//
//func UpdateContextByKey(listContext []DataContext, key string, value interface{}) []DataContext {
//	for i, context := range listContext {
//		if context.Key == key {
//			context.Value = value
//			listContext[i] = context
//			return listContext
//		}
//	}
//	return listContext
//}
