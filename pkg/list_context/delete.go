package list_context

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
