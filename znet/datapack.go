package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/deanisty/zinx/utils"
	"github.com/deanisty/zinx/ziface"
)

type IDataPack struct {

}

func NewDataPack() *IDataPack {
	return &IDataPack{}
}

func (dp *IDataPack) GetHeadLen() uint32 {
	return 8
}

func (dp *IDataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	// 创建一个字节缓冲
	dataBuf := bytes.NewBuffer([]byte{})
	// 写入数据长度
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}
	// 写入messageID
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	// 写入数据data
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuf.Bytes(), nil
}

func (dp *IDataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	// data reader
	dataBuf := bytes.NewReader(binaryData)

	msg := &Message{}

	// dataLen
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	// MsgId
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	//
	if utils.GlobalObject.MaxPacketSize > 0 && msg.DataLen > utils.GlobalObject.MaxPacketSize {
		return nil, errors.New("too large message data received")
	}

	return msg, nil
}