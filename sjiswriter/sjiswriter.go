package sjiswriter

import (
	"bytes"
	"io"
	"unicode/utf8"
)

type SJISWriter struct {
	w io.Writer
}

func NewSJISWriter(w io.Writer) *SJISWriter {
	return &SJISWriter{w}
}

var conversionPair = [][][]byte{
	{[]byte("〜"), []byte("～")},
	{[]byte("−"), []byte("－")},
	{[]byte("¢"), []byte("￠")},
	{[]byte("£"), []byte("￡")},
	{[]byte("¬"), []byte("￢")},
	{[]byte("–"), []byte("－")},
	{[]byte("—"), []byte("―")},
	{[]byte("‖"), []byte("∥")},
	{[]byte("‾"), []byte("￣")},
	{[]byte("ø"), []byte("Φ")},
	{[]byte("›"), []byte("〉")},
}

var irreplacableCharByte = []byte("?")

func (sw *SJISWriter) Write(b []byte) (n int, err error) {
	for i := range conversionPair {
		b = bytes.ReplaceAll(b, conversionPair[i][0], conversionPair[i][1])
	}

	for len(b) > 0 {
		_, size := utf8.DecodeRune(b)
		if size == 0 {
			break
		}

		_, err = sw.w.Write(b[:size])
		if err != nil {
			_, err = sw.w.Write(irreplacableCharByte)
			if err != nil {
				break
			}
		}

		n += size
		b = b[size:]
	}
	return
}
