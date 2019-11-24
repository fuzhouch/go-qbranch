// MIT License
//
// Copyright (c) 2018 Fuzhou Chen
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files
// (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge,
// publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
// CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
// SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package qbranch

// Reader represents an interface that used to read fields from
// serialized buffer.
//
// IMPORTANT: Reader is NOT THREAD SAFE. It keeps an internal state to
// support moving read pointers like Scan().
type Reader interface {
	Scan() (done bool, index uint)
	Error() error

	// All Read*() functions automatically advance internal pointer
	// to next position, but happens only when Scan() returns false
	// and no err. This is to allow we support reading nullable.

	ReadBool() (value bool, err error)
	ReadInt8() (value int8, err error)
	ReadUInt8() (value uint8, err error)
	ReadInt16() (value int16, err error)
	ReadUInt16() (value uint16, err error)
	ReadInt32() (value int32, err error)
	ReadUInt32() (value uint32, err error)
	ReadInt64() (value int64, err error)
	ReadUInt64() (value uint64, err error)
	ReadString() (value string, err error)
	ReadWString() (value string, err error)

	// Only "required nullable<T>" expect reading an explicit Null
	// from serialized buffer. To allow we do this, the generated
	// code always read "required nullable" types twice: first
	// read a null, if it returns error, then try to read via
	// standard Read*().
	ReadNull() (isNull bool, err error)
}
