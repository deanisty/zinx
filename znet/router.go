package znet

import (
	"github.com/deanisty/zinx/ziface"
)

// 路由基类
type BaseRouter struct {}

// 这里之所以BaseRouter的方法都为空
// 是因为有的Router不希望有PreHandler和PostHandler
// 所有router全部继承BaseRouter的好处是，不需要实现PreHandler和PostHandler也可以是梨花
func (br *BaseRouter) PreHandler(req ziface.IRequest) {}
func (br *BaseRouter) Handler(req ziface.IRequest) {}
func (br *BaseRouter) PostHandler(req ziface.IRequest) {}