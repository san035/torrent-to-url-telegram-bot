package list_context

import (
	"errors"
	"sync/atomic"
)

// Add return уникальный id контекста
func (listContext *ListContext) Add(newValueContext *DataContext) (uint64, error) {
	listContext.Mutex.Lock()
	defer listContext.Mutex.Unlock()

	for _, oldContext := range listContext.List {
		if oldContext == nil {
			continue
		}
		if oldContext.MagnetUrl == nil {
			continue
		}
		if *oldContext.MagnetUrl == *newValueContext.MagnetUrl {
			return oldContext.Id, errors.New("context is busy")
		}
	}

	listContext.LastItem = (listContext.LastItem + 1) % MaxCountContext
	listContext.List[listContext.LastItem] = newValueContext
	atomic.AddUint64(&lastId, 1)
	listContext.List[listContext.LastItem].Id = lastId
	return lastId, nil
}
