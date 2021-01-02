package sjiswriter

import (
	"bytes"
	"testing"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func TestSJISWriter_Write(t *testing.T) {
	type fields struct {
		b *bytes.Buffer
	}
	type args struct {
		b []byte
	}
	type wants struct {
		n     int
		isErr bool
		b     []byte
	}

	runFunc := func(t *testing.T, fields fields, args args, wants wants) {
		sw := &SJISWriter{
			w: transform.NewWriter(fields.b, japanese.ShiftJIS.NewEncoder()),
		}
		gotN, err := sw.Write(args.b)
		if (err != nil) != wants.isErr {
			t.Errorf("SJISWriter.Write() error = %v, wantErr %v", err, wants.isErr)
			return
		}
		if gotN != wants.n {
			t.Errorf("SJISWriter.Write() = %v, want %v", gotN, wants.n)
		}
		if !bytes.Equal(fields.b.Bytes(), wants.b) {
			t.Errorf("SJISWriter.Write() []byte = %v, want %v", fields.b.Bytes(), wants.b)
		}
	}

	t.Run("aあア亜", func(t *testing.T) {
		var b bytes.Buffer
		runFunc(t, fields{b: &b}, args{b: []byte("aあア亜")}, wants{n: 10, isErr: false, b: []byte{97, 130, 160, 131, 65, 136, 159}})
	})

	t.Run("〜", func(t *testing.T) {
		var b bytes.Buffer
		runFunc(t, fields{b: &b}, args{b: []byte("〜")}, wants{n: 3, isErr: false, b: []byte{129, 96}})
	})

	t.Run("−", func(t *testing.T) {
		var b bytes.Buffer
		runFunc(t, fields{b: &b}, args{b: []byte("−")}, wants{n: 3, isErr: false, b: []byte{129, 124}})
	})

	t.Run("¢", func(t *testing.T) {
		var b bytes.Buffer
		runFunc(t, fields{b: &b}, args{b: []byte("¢")}, wants{n: 3, isErr: false, b: []byte{129, 145}})
	})

	t.Run("£", func(t *testing.T) {
		var b bytes.Buffer
		runFunc(t, fields{b: &b}, args{b: []byte("£")}, wants{n: 3, isErr: false, b: []byte{129, 146}})
	})

	t.Run("¬", func(t *testing.T) {
		var b bytes.Buffer
		runFunc(t, fields{b: &b}, args{b: []byte("¬")}, wants{n: 3, isErr: false, b: []byte{129, 202}})
	})

	t.Run("–", func(t *testing.T) {
		var b bytes.Buffer
		runFunc(t, fields{b: &b}, args{b: []byte("–")}, wants{n: 3, isErr: false, b: []byte{129, 124}})
	})

	t.Run("—", func(t *testing.T) {
		var b bytes.Buffer
		runFunc(t, fields{b: &b}, args{b: []byte("—")}, wants{n: 3, isErr: false, b: []byte{129, 92}})
	})

	t.Run("‖", func(t *testing.T) {
		var b bytes.Buffer
		runFunc(t, fields{b: &b}, args{b: []byte("‖")}, wants{n: 3, isErr: false, b: []byte{129, 97}})
	})

	t.Run("‾", func(t *testing.T) {
		var b bytes.Buffer
		runFunc(t, fields{b: &b}, args{b: []byte("‾")}, wants{n: 3, isErr: false, b: []byte{129, 80}})
	})

	t.Run("ø", func(t *testing.T) {
		var b bytes.Buffer
		runFunc(t, fields{b: &b}, args{b: []byte("ø")}, wants{n: 2, isErr: false, b: []byte{131, 179}})
	})

	t.Run("›", func(t *testing.T) {
		var b bytes.Buffer
		runFunc(t, fields{b: &b}, args{b: []byte("›")}, wants{n: 3, isErr: false, b: []byte{129, 114}})
	})

	t.Run("鑫", func(t *testing.T) {
		var b bytes.Buffer
		runFunc(t, fields{b: &b}, args{b: []byte("鑫")}, wants{n: 3, isErr: false, b: []byte{63}})
	})
}
