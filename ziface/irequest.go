package ziface

/*
	IRequest接口
	封装了客户端连接信息和请求数据
 */
type IRequest interface {
	GetConnection() IConnection // 客户端连接
	GetData() []byte // 请求数据
}
