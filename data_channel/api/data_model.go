package api

type DataChannel interface {
	InitContext() (err error)
	GetContext() (ctx interface{}, err error)
	SetContext(ctx interface{}) (err error)
	DeleteContext() (err error)
	UpdateContext(ctx interface{}) (err error)
}
