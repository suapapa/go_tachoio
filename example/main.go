// Copyright 2016, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"crypto/rand"
	"io"
	"io/ioutil"
	"log"
	"time"

	tachoio "github.com/suapapa/go_tachoio"
)

func main() {
	in := tachoio.NewReader(rand.Reader)
	out := tachoio.NewWriter(ioutil.Discard)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go func(ctx context.Context) {
		log.Println("Start copying...")
		n, err := io.Copy(out, in)
		if err != nil {
			panic(err)
		}
		log.Printf("Copy %v bytes\n", n)
	}(ctx)

	chTick := time.Tick(time.Second)
copyLoop:
	for {
		select {
		case <-chTick:
			go func(ctx context.Context) {
				rn, rd := in.ReadMeter()
				log.Printf("Read %d in %v. %.02fBPS", rn, rd, float64(rn)/rd.Seconds())
			}(ctx)
			go func(ctx context.Context) {
				wn, wd := out.WriteMeter()
				log.Printf("Write %d in %v. %.02fBPS", wn, wd, float64(wn)/wd.Seconds())
			}(ctx)
		case <-ctx.Done():
			break copyLoop
		}
	}
	log.Println(in)
	log.Println(out)
	log.Println("All done")
}
