package lesson_05a

type tree struct {
	v int
	l *tree
	r *tree
}

// function to sum up the values in the binary tree
// is this like an extension function on the struct above?
// yes: it is called a "method receiver"
// it makes sense that we would use a pointer here so that we don't need
// to recursively allocate
func (t *tree) Sum() int {

	sum := t.v

	if t.l != nil {
		sum += t.l.Sum()
	}

	if t.r != nil {
		sum += t.r.Sum()
	}

	return sum
}

// the default value of a pointer is nil, so we can improve this
func (t *tree) SumImproved() int {

	if t == nil {
		return 0
	}

	return t.v + t.l.SumImproved() + t.r.SumImproved()
}
