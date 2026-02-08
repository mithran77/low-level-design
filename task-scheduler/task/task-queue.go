package task

type TaskHeap []*Task

func (th TaskHeap) Len() int {
	return len(th)
}

func (th TaskHeap) Less(i, j int) bool {
	if th[i].ExecutionTime.Compare(th[j].ExecutionTime) == 0 {
		// Higher priority first
		return th[i].Priority > th[j].Priority
	}
	return th[i].ExecutionTime.Before(th[j].ExecutionTime)
}

func (th TaskHeap) Swap(i, j int) {
	th[i], th[j] = th[j], th[i]
}

func (th *TaskHeap) Push(t any) {
	*th = append(*th, t.(*Task))
}

func (th TaskHeap) Pop() any {
	n := th.Len()
	last := th[n-1]
	th[n-1] = nil
	th = th[0 : n-1]
	return last
}

func NewTaskHeap() *TaskHeap {
	return &TaskHeap{}
}
