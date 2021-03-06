package p2p

import (
	"math/rand"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

func newMessage(payLoad string) *Message {
	return &Message{
		Code:    ctlMsgPingCode,
		Payload: []byte(payLoad),
	}
}

func getRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func Test_message(t *testing.T) {
	randStr1 := getRandomString(zipBytesLimit - 50)
	msg1 := newMessage(randStr1)

	err := msg1.Zip()
	assert.Equal(t, err, nil)
	assert.Equal(t, string(msg1.Payload), randStr1)

	err = msg1.UnZip()
	assert.Equal(t, err, nil)
	assert.Equal(t, string(msg1.Payload), randStr1)

	randStr2 := getRandomString(zipBytesLimit + 50)
	msg2 := newMessage(randStr2)

	err = msg2.Zip()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(msg2.Payload) < len([]byte(randStr2)), true)

	err = msg2.UnZip()
	assert.Equal(t, err, nil)
	assert.Equal(t, string(msg2.Payload), randStr2)
}
