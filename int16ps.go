// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import "bytes"

// Int16ps returns stringer/JSON/text marshaler for the int16 pointer slice type.
func Int16ps(a ...*int16) int16PS { return int16PS{A: a} }

type int16PS struct{ A []*int16 }

func (a int16PS) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a int16PS) MarshalText() ([]byte, error) {
	if a.A == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	for i, p := range a.A {
		b, err := int16P{P: p}.MarshalText()
		if err != nil {
			return nil, err
		}
		if i != 0 {
			buf.WriteString(" ")
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func (a int16PS) MarshalJSON() ([]byte, error) {
	if a.A == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.A {
		b, err := int16P{P: p}.MarshalJSON()
		if err != nil {
			return nil, err
		}
		if i != 0 {
			buf.WriteString(",")
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
	}
	buf.WriteString("]")
	return buf.Bytes(), nil
}
