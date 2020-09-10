package binary

import (
	"github.com/davyxu/goobjfmt"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
)

type binaryCodec struct {
}

func (self *binaryCodec) Name() string {
	return "binary"
}

func (self *binaryCodec) MimeType() string {
	return "application/binary"
}

func (self *binaryCodec) Encode(msgObj interface{}, ctx cellnet.ContextSet) (data interface{}, err error) {

	return goobjfmt.BinaryWrite(msgObj)

}

func (self *binaryCodec) Decode(data interface{}, msgObj interface{}) error {

	return goobjfmt.BinaryRead(data.([]byte), msgObj)
}

func init() {

	codec.RegisterCodec(new(binaryCodec))
}
