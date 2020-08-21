package amf

import (
	"bytes"
	"fmt"
	"log"
)

const (
	ADD = 0x0
	DEL = 0x3
)

const (
	SetDataFrame string = "@setDataFrame"
	OnMetaData   string = "onMetaData"
)

var setFrameFrame []byte

/**
@TODO update implement detail
*/
func init() {
	b := bytes.NewBuffer(nil)

	if _, err := WriteValue(b, SetDataFrame); err != nil {
		log.Fatal(err)
	}
	setFrameFrame = b.Bytes()
}

func MetaDataReform(p []byte, flag byte) ([]byte, error) {
	r := bytes.NewReader(p)
	decoder := ReadValue
	switch flag {
	case ADD:
		v, err := decoder(r)
		if err != nil {
			return nil, err
		}
		switch v.(type) {
		case string:
			vv := v.(string)
			if vv != SetDataFrame {
				tmplen := len(setFrameFrame)
				b := make([]byte, tmplen+len(p))
				copy(b, setFrameFrame)
				copy(b[tmplen:], p)
				p = b
			}
		default:
			return nil, fmt.Errorf("setFrameFrame error")
		}
	case DEL:
		v, err := decoder(r)
		if err != nil {
			return nil, err
		}
		switch v.(type) {
		case string:
			vv := v.(string)
			if vv == SetDataFrame {
				p = p[len(setFrameFrame):]
			}
		default:
			return nil, fmt.Errorf("metadata error")
		}
	default:
		return nil, fmt.Errorf("invalid flag:%d", flag)
	}
	return p, nil
}
