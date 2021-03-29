// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"
)

// Complex128s returns stringer/JSON marshaler interface implementation for the complex128 slice type.
func Complex128s(a ...complex128) complex128sV { return complex128sV{A: a} }

type complex128sV struct{ A []complex128 }

func (a complex128sV) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a complex128sV) MarshalText() ([]byte, error) {
	if a.A == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	for i, v := range a.A {
		if i != 0 {
			buf.WriteString(" ")
		}
		b, err := complex128V{V: v}.MarshalText()
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func (a complex128sV) MarshalJSON() ([]byte, error) {
	if a.A == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range a.A {
		if i != 0 {
			buf.WriteString(",")
		}
		b, err := complex128V{V: v}.MarshalJSON()
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
	}
	buf.WriteString("]")
	return buf.Bytes(), nil
}