package core

type Caller[T any] struct {
	Calls []func(T)
}

func (c *Caller[T]) Add(callback func(T)) {
	c.Calls = append(c.Calls, callback)
}

func (c *Caller[T]) Invoke(data T) {
	for i := range c.Calls {
		c.Calls[i](data)
	}
}
