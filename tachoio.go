// Copyright 2016, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tachoio // import "github.com/suapapa/go_tachoio"

import (
	"fmt"
	"io"
	"runtime"
	"sync"
	"time"
)

// Reader implements tacho-meter for io.Reader object.
type Reader struct {
	ts  time.Time
	cnt int
	rd  io.Reader

	sync.RWMutex
}

// NewReader retrurs a new Reader
func NewReader(rd io.Reader) *Reader {
	r := Reader{
		rd: rd,
		ts: time.Now(),
	}
	return &r
}

func (r *Reader) Read(p []byte) (n int, err error) {
	r.Lock()
	defer r.Unlock()

	n, err = r.rd.Read(p)
	r.cnt += n

	runtime.Gosched()
	return
}

// ReadMeter returns read bytes and duration since last check
func (r *Reader) ReadMeter() (n int, d time.Duration) {
	r.Lock()
	defer r.Unlock()

	d = time.Since(r.ts)
	r.ts = time.Now()
	n = r.cnt
	r.cnt = 0
	return
}

func (r *Reader) String() string {
	r.RLock()
	defer r.RUnlock()

	return fmt.Sprintf("tachoio.Reader(transpered %d bytes in %s)",
		r.cnt, time.Since(r.ts))
}

// Writer implements tacho-meter for io.Writer object.
type Writer struct {
	ts  time.Time
	cnt int
	wr  io.Writer

	sync.RWMutex
}

// NewWriter returns a new Writer
func NewWriter(wr io.Writer) *Writer {
	w := Writer{
		wr: wr,
		ts: time.Now(),
	}
	return &w
}

func (w *Writer) Write(p []byte) (n int, err error) {
	w.Lock()
	defer w.Unlock()

	n, err = w.wr.Write(p)
	w.cnt += n

	runtime.Gosched()
	return
}

// WriteMeter returns written bytes and duration since last check
func (w *Writer) WriteMeter() (n int, d time.Duration) {
	w.Lock()
	defer w.Unlock()

	d = time.Since(w.ts)
	w.ts = time.Now()
	n = w.cnt
	w.cnt = 0
	return
}

func (w *Writer) String() string {
	w.RLock()
	defer w.RUnlock()

	return fmt.Sprintf("tachoio.Writer(transpered %d bytes in %s)",
		w.cnt, time.Since(w.ts))
}
