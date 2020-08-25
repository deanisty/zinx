package ziface

type IMessage interface {
	GetDataLen() uint32 // 获取数据长度
	GetMsgId() uint32   // 获取消息id
	GetData() []byte    // 获取消息

	SetDataLen(uint32)  // 设置消息长度
	SetMsgId(uint32)    // 设置消息id
	SetData([]byte)     // 设置消息
}
