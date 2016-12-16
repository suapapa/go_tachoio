// Copyright 2016, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tachoio // import "github.com/suapapa/go_tachoio"

// NoopReader does nothing when read from it
var NoopReader = &noopRead{}

type noopRead struct{}

func (n *noopRead) Read(p []byte) (int, error) {
	return len(p), nil
}
