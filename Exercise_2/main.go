package Exercise_2

// My first documentation.
type Test_A []interface{}

func (p *Test_A) Push(v interface{}) {
	*p = append(*p, v.(int))
}

// My two documentation.
func (p *Test_A) Pop() interface{} {
	head := (*p)[0]
	*p = (*p)[1:]
	return head
}
func (p *Test_A) IsEmpty() bool {
	return len(*p) == 0
}
