package stack

/********************************* Int Stack implementation ***************************************/

// IntStack is an easy stack implementation that provides the usual stack operations
type IntStack struct {
	a []int
}

// NewIntStack creates a new stack instance
func NewIntStack(cap int) *IntStack {
	return &IntStack{a: make([]int, 0, cap)}
}

// Get returns the ith element from the stack
func (s *IntStack) Get(i int) int {
	return s.a[i]
}

// Push adds e to the top of the stack
func (s *IntStack) Push(e int) {
	s.a = append(s.a, e)
}

// Pop removes the top from the stack and returns the removed element
func (s *IntStack) Pop() int {
	e := s.a[len(s.a)-1]
	s.a = s.a[:len(s.a)-1]
	return e
}

// Peek returns the top element of the stack without removing it
func (s *IntStack) Peek() int {
	return s.a[len(s.a)-1]
}

/*********************************** String Stack implementation *****************************************/

// StringStack is an easy stack implementation that stores strings and provides the usual methods of a stack.
type StringStack struct {
	a []string
}

// NewStringStack creates a new stack instance.
func NewStringStack(cap int) *StringStack {
	return &StringStack{a: make([]string, 0, cap)}
}

// Get returns the ith element from the stack
func (s *StringStack) Get(i int) string {
	return s.a[i]
}

// Push adds e to the top of the stack
func (s *StringStack) Push(e string) {
	s.a = append(s.a, e)
}

// Pop removes the top from the stack and returns the removed element
func (s *StringStack) Pop() string {
	e := s.a[len(s.a)-1]
	s.a = s.a[:len(s.a)-1]
	return e
}

// Peek returns the top element of the stack without removing it
func (s *StringStack) Peek() string {
	return s.a[len(s.a)-1]
}
