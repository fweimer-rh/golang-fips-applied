// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build cmd_go_bootstrap || !cgo
// +build cmd_go_bootstrap !cgo

package sha1

import (
	"hash"
)

var boringEnabled = false

func boringNewSHA1() hash.Hash { panic("!no_openssl: not available") }

func boringUnreachable() {}

func boringSHA1([]byte) [20]byte { panic("!no_openssl: not available") }
