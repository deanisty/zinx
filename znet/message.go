package znet

import (
	"github.com/deanisty/zinx/ziface"
)

type Message struct {
	Id uint32    // 消息id
	Data []byte  // 消息内容
	DataLen uint32 // 消息长度
}

func (m *Message) GetMsgId() uint32 {
	return m.Id
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}

func (m *Message) SetDataLen(len uint32) {
	m.DataLen = len
}

func NewMessagePacket(msgId uint32, data []byte) ziface.IMessage {
	msg := &Message{
		Id:      msgId,
		Data:    data,
		DataLen: uint32(len(data)),
	}

	return msg
}