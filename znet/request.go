package znet

import (
	"github.com/deanisty/zinx/ziface"
)

type Request struct {
	conn ziface.IConnection    // 客户端建立的连接
	data ziface.IMessage         // 客户端请求的数据
}

// 获取连接信息
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// 获取请求的数据
func (r *Request) GetData() []byte {
	return r.data.GetData()
}

func (r *Request) GetMsgId() uint32 {
	return r.data.GetMsgId()
}
