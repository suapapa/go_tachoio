package tachoio_test

import (
	"crypto/rand"
	"io"
	"io/ioutil"
	"testing"

	tachoio "github.com/suapapa/go_tachoio"
)

func BenchmarkRandReader(b *testing.B) {
	buf := make([]byte, 256*1024)
	for i := 0; i < b.N; i++ {
		rand.Reader.Read(buf)
	}
}

func BenchmarkDiscardWriter(b *testing.B) {
	buf := make([]byte, 256*1024)
	for i := 0; i < b.N; i++ {
		ioutil.Discard.Write(buf)
	}
}

func BenchmarkRandReaderDiscardWriter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		io.CopyN(ioutil.Discard, rand.Reader, 256*1024)
	}
}

func BenchmarkRandTachoReader(b *testing.B) {
	for i := 0; i < b.N; i++ {
		io.CopyN(ioutil.Discard, tachoio.NewReader(rand.Reader), 256*1024)
	}
}

func BenchmarkDiscardTachoWriter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		io.CopyN(tachoio.NewWriter(ioutil.Discard), rand.Reader, 256*1024)
	}
}

func BenchmarkTachoRandReaderTachoDiscardWriter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		io.CopyN(tachoio.NewWriter(ioutil.Discard), tachoio.NewReader(rand.Reader), 256*1024)
	}
}
