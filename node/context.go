package node

// Context has a number which is basically used to create incremental stuff.
// Example incremental stuff: state number(q0, q1, q2)
type Context struct {
	N int
}

// NewContext returns a new Context.
// The default value of N is -1.
func NewContext() *Context {
	return &Context{
		N: -1,
	}
}

// Increment add 1 to N which held in Context struct,
// and returns the number.
func (ctx *Context) Increment() int {
	ctx.N++
	return ctx.N
}
