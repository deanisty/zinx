package znet

import (
	"fmt"
	"testing"
)

func TestMessage(t *testing.T) {
	data := []byte("abc")
	msg := NewMessagePacket(1, data)

	fmt.Printf("%T", msg)
}
