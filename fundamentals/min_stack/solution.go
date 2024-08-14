package min_stack

// Start 14:22
// End 14:46
// https://leetcode.com/problems/min-stack

type Element struct {
	value    int
	prevLast *Element
	prevMin  *Element
}

type MinStack struct {
	last *Element
	min  *Element
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	newEl := new(Element)
	newEl.value = val
	if this.last == nil {
		this.last = newEl
		this.min = newEl
	} else {
		newEl.prevLast = this.last
		this.last = newEl
		if newEl.value <= this.min.value {
			newEl.prevMin = this.min
			this.min = newEl
		}
	}
}

func (this *MinStack) Pop() {
	if this.last != nil {
		if this.last == this.min {
			this.min = this.min.prevMin
		}
		this.last = this.last.prevLast
	}
}

func (this *MinStack) Top() int {
	return this.last.value
}

func (this *MinStack) GetMin() int {
	return this.min.value
}
