package vm

type StatusType string

const (
	INITIALIZED = "INITIALIZED"
	FORKED      = "FORKED"
	KILLED      = "KILLED"
	RUNNING     = "RUNNING"
	SUCCEEDED   = "SUCCEEDED"
	FAILED      = "FAILED"
)

type Index struct {
	Start int
	End   int
}

type IndexMap map[string]Index

type Context struct {
	Status StatusType

	StartCodeCursor   int
	CurrentCodeCursor int
	EndCodeCursor     int

	StartInputCursor   int
	CurrentInputCursor int
	EndInputCursor     int

	IndexMap IndexMap
}

func NewContext() *Context {
	m := make(IndexMap)
	return (&Context{}).SetStatus(INITIALIZED).SetIndexMap(m)
}

func (context *Context) Clone() *Context {
	m := context.CloneIndexMap()
	return (&Context{}).
		SetStatus(context.Status).
		SetStartInputCursor(context.StartInputCursor).
		SetCurrentInputCursor(context.CurrentInputCursor).
		SetEndInputCursor(context.EndInputCursor).
		SetStartCodeCursor(context.StartCodeCursor).
		SetCurrentCodeCursor(context.CurrentCodeCursor).
		SetEndCodeCursor(context.EndCodeCursor).
		SetIndexMap(m)
}

func (context *Context) CloneIndexMap() IndexMap {
	m := make(IndexMap)
	for k, v := range context.IndexMap {
		m[k] = v
	}
	return m
}

func (context *Context) SetStatus(status StatusType) *Context {
	context.Status = status
	return context
}

func (context *Context) SetStartCodeCursor(cursor int) *Context {
	context.StartCodeCursor = cursor
	return context
}

func (context *Context) SetCurrentCodeCursor(cursor int) *Context {
	context.CurrentCodeCursor = cursor
	return context
}

func (context *Context) SetEndCodeCursor(cursor int) *Context {
	context.EndCodeCursor = cursor
	return context
}

func (context *Context) SetStartInputCursor(cursor int) *Context {
	context.StartInputCursor = cursor
	return context
}

func (context *Context) SetCurrentInputCursor(cursor int) *Context {
	context.CurrentInputCursor = cursor
	return context
}

func (context *Context) SetEndInputCursor(cursor int) *Context {
	context.EndInputCursor = cursor
	return context
}

func (context *Context) SetIndexMap(m IndexMap) *Context {
	context.IndexMap = m
	return context
}
