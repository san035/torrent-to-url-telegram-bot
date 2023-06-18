package list_context

func (listContext *ListContext) GetById(id uint64) *DataContext {
	for _, dataContext := range listContext.List {
		if dataContext.Id == id {
			return dataContext
		}
	}
	return nil
}
