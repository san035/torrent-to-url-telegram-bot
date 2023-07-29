package list_context

import (
	"context"
	"sync"
)

type DataContext struct {
	Context   *context.Context
	MagnetUrl *string
	Id        uint64
}

const MaxCountContext = 100

type ListContext struct {
	Mutex    sync.Mutex
	List     *[MaxCountContext]*DataContext
	LastItem int
}

var lastId uint64

func NewListContext() (listContext *ListContext) {
	listContext = new(ListContext)
	listContext.List = new([MaxCountContext]*DataContext)
	listContext.LastItem = -1
	return
}

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
