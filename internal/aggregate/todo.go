package aggregate

type Todos []*Todo

type Todo struct {
	ID         uint64
	ActivityID uint64
	Activity   MapActivities
	Title      string
	IsActive   int
	Priority   string
}

func NewTodo(activity MapActivities, title string, isActive int, priority string) (*Todo, error) {
	return &Todo{
		Activity: activity,
		Title:    title,
		IsActive: isActive,
		Priority: priority,
	}, nil
}

func RebuildTodos(id uint64, activity uint64, title string, isActive int, priority string) (*Todo, error) {
	return &Todo{
		ID:         id,
		ActivityID: activity,
		Title:      title,
		IsActive:   isActive,
		Priority:   priority,
	}, nil
}

func RebuildTodo(id uint64, activity MapActivities, title string, isActive int, priority string) (*Todo, error) {
	return &Todo{
		ID:       id,
		Activity: activity,
		Title:    title,
		IsActive: isActive,
		Priority: priority,
	}, nil
}
