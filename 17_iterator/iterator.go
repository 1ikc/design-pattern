package iterator

type Iterator interface {
	HasNext() bool
	Next()
	Current() interface{}
}

type ArrayInt []int

type ArrayIntIterator struct {
	arrayInt ArrayInt
	idx	int
}

func (i *ArrayIntIterator) HasNext() bool {
	return i.idx < len(i.arrayInt) - 1
}

func (i *ArrayIntIterator) Next() {
	i.idx++
}

func (i *ArrayIntIterator) Current() interface{} {
	return i.arrayInt[i.idx]
}

func (a ArrayInt) Iterator() Iterator {
	return &ArrayIntIterator{
		arrayInt: a,
	}
}