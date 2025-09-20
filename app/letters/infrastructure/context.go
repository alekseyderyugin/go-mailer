package infrastructure

type Context struct {
	errorHandler func(err error)
}

func NewContext(errorHandler func(err error)) *Context {
	return &Context{
		errorHandler: errorHandler,
	}
}

func (ctx *Context) HandleError(err error) {
	if err != nil {
		ctx.errorHandler(err)
	}
}
