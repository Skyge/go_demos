package MiniStack

/**
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
	push(x) -- Push element x onto stack.
	pop() -- Removes the element on top of the stack.
	top() -- Get the top element.
	getMin() -- Retrieve the minimum element in the stack.

Example:
	MinStack minStack = new MinStack();
	minStack.push(-2);
	minStack.push(0);
	minStack.push(-3);
	minStack.getMin();   --> Returns -3.
	minStack.pop();
	minStack.top();      --> Returns 0.
	minStack.getMin();   --> Returns -2.
*/
type MinStack struct {
	stack []int
	min   *node
}

type node struct {
	prevMin *node
	idx     int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(x int) {
	m.stack = append(m.stack, x)
	if m.min == nil || m.stack[m.min.idx] > x {
		n := &node{prevMin: m.min, idx: len(m.stack) - 1}
		m.min = n
	}
}

func (m *MinStack) Pop() {
	idx := len(m.stack) - 1
	if m.min != nil && m.min.idx == idx {
		m.min = m.min.prevMin
	}
	m.stack = m.stack[:idx]
}

func (m *MinStack) Top() int {
	return m.stack[len(m.stack)-1]
}

func (m *MinStack) GetMin() int {
	return m.stack[m.min.idx]
}
