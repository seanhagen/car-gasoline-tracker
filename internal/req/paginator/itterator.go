package paginator

type (
	// Itterator allows you to itterate over a PagerCol
	Itterator struct {
		current int
		data    PagerColl
	}
)

// Next returns true if there are more elements in the itterator
func (it *Itterator) Next() bool {
	it.current++
	return it.current < it.data.Count()
}

// Current returns the current value from the itterator
func (it Itterator) Current() Paginatable {
	return it.data.Item(it.current)
}

// Reset resets the index to the start
func (it *Itterator) Reset() {
	it.current = -1
}

// NewItterator returns a new itterator for use with PagerColl interfaces
func NewItterator(data PagerColl) *Itterator {
	return &Itterator{data: data, current: -1}
}
