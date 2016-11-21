package tachoio

import (
	"io"
	"sync"
	"time"
)

type Reader struct {
	ts  time.Time
	cnt int
	rd  io.Reader

	sync.Mutex
}

// NewReader retrurs a new Reader
func NewReader(rd io.Reader) *Reader {
	r := new(Reader)
	r.ts = time.Now()
	return r
}

func (r *Reader) Read(p []byte) (n int, err error) {
	r.Lock()
	defer r.Unlock()

	n, err = r.rd.Read(p)
	r.cnt += n
	return
}

func (r *Reader) ReadMeter() (n int, d time.Duration) {
	r.Lock()
	defer r.Unlock()

	d = time.Since(r.ts)
	r.ts = time.Now()
	n = r.cnt
	r.cnt = 0
	return
}

type Writer struct {
	ts  time.Time
	cnt int
	wr  io.Writer

	sync.Mutex
}

// NewWriter returns a new Writer
func NewWriter(wr io.Writer) *Writer {
	w := new(Writer)
	return w
}

func (w *Writer) Write(p []byte) (n int, err error) {
	w.Lock()
	defer w.Unlock()

	n, err = w.wr.Write(p)
	w.cnt += n
	return
}

func (w *Writer) WriteMeter() (n int, d time.Duration) {
	w.Lock()
	defer w.Unlock()

	d = time.Since(w.ts)
	w.ts = time.Now()
	n = w.cnt
	w.cnt = 0
	return
}
