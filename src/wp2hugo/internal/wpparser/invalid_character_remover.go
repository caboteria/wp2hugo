package wpparser

import (
	"bytes"
	"io"
)

type InvalidatorCharacterRemover struct {
	reader io.Reader
}

func (i InvalidatorCharacterRemover) Read(p []byte) (int, error) {
	tmp := make([]byte, len(p))
	n, err := i.reader.Read(tmp)
	if err != nil {
		return n, err
	}
	// Characters from 1 to 31 seem to be disallowed in XML
	// One gets errors like "XML syntax error on line <>: illegal character code U+0001"
	// Ref: https://github.com/ashishb/wp2hugo/issues/27
	for i := 0; i <= 31; i++ {
		tmp = bytes.ReplaceAll(tmp, []byte{byte(i)}, []byte(""))
	}
	copy(p, tmp)
	return len(tmp), nil
}
