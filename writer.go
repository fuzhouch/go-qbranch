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

// Writer represents an interface to support writing fields to byte
// array to Bond buffer.
type Writer interface {
	WriteBool(index uint, value bool)
	WriteInt8(index uint, value int8)
	WriteUInt8(index uint, value uint8)
	WriteInt16(index uint, value int16)
	WriteUInt16(index uint, value uint16)
	WriteInt32(index uint, value int32)
	WriteUInt32(index uint, value uint32)
	WriteInt64(index uint, value int64)
	WriteUInt64(index uint, value uint64)
	WriteString(index uint, value string)
	WriteWString(index uint, value string)
	WriteNull(index uint)

	// Container type is implemented in a different way.
	// As we know all Bond containers are generic, while
	// generic is completely unsupported in Golang. Though
	// a clever way is to use interface{}, it breaks a
	// basic design goal of Bond, that we want generated
	// structures to be type safe.
	//
	// So, QBranch choose another way to make it work.
	// For each use of container type like map[string]string, we
	// create a type alias and provides an WriteAsByte() method.
}
