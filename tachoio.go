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
	d  time.Duration
	n  int
	rd io.Reader

	sync.RWMutex
}

// NewReader retrurs a new Reader
func NewReader(rd io.Reader) *Reader {
	return &Reader{rd: rd}
}

func (r *Reader) Read(p []byte) (n int, err error) {
	r.Lock()
	defer r.Unlock()

	start := time.Now()
	n, err = r.rd.Read(p)
	r.n += n
	r.d += time.Since(start)

	runtime.Gosched()
	return
}

// ReadMeter returns read bytes and duration since last check
func (r *Reader) ReadMeter() (n int, d time.Duration) {
	r.Lock()
	defer r.Unlock()

	d = r.d
	n = r.n
	r.n = 0
	r.d = 0
	return
}

func (r *Reader) String() string {
	r.RLock()
	defer r.RUnlock()

	return fmt.Sprintf("tachoio.Reader(n: %d, d: %v)", r.n, r.d)
}

// Writer implements tacho-meter for io.Writer object.
type Writer struct {
	d  time.Duration
	n  int
	wr io.Writer

	sync.RWMutex
}

// NewWriter returns a new Writer
func NewWriter(wr io.Writer) *Writer {
	return &Writer{wr: wr}
}

func (w *Writer) Write(p []byte) (n int, err error) {
	w.Lock()
	defer w.Unlock()

	start := time.Now()
	n, err = w.wr.Write(p)
	w.d += time.Since(start)
	w.n += n

	runtime.Gosched()
	return
}

// WriteMeter returns written bytes and duration since last check
func (w *Writer) WriteMeter() (n int, d time.Duration) {
	w.Lock()
	defer w.Unlock()

	n = w.n
	d = w.d
	w.n = 0
	w.d = 0
	return
}

func (w *Writer) String() string {
	w.RLock()
	defer w.RUnlock()

	return fmt.Sprintf("tachoio.Writer(n: %d, d: %v)", w.n, w.d)
}
