package main

/*
Non-idiomatic means not preferred usage
Idiomatic means preferred usage
*/

func NamingThings() {
	/*
	Why naming is important?
	1. Critical for Readability = Maintainability
	2. Use the first few letters of the words 
		like var fv string // flag value
	3. Use fewer letters in smaller scopes 
		like var bytesRead int // number of bytes | BAD!! (this name is too verbose )
		var n int // number of bytes | GOOD!!
	4. Use the complete words in larger scopes
	5. Use mixedCaps like this type PlayerScore struct
	6. Use all capitals for acronyms like var localAPI string
	7. Do not stutter like 
		player.PlayerScore | BAD!!
		player.Score | GOOD!!
	8. Do not use under_scores or LIKE_THIS like
		const MAX_TIME int | BAD!!
		const MaxTime int | GOOD!!
		const N int | GOOD!!
	*/
}

// Non-Idiomatic example
// func readNonIdiomatic(buffer *Buffer, inBuffer []byte) (size int, err error) {
// 	if buffer.empty() {
// 		buffer.Reset()
// 	}

// 	size = copy(
// 		inBuffer,
// 		buffer.buffer[buffer.offset:])

// 	buffer.offset += size
// 	return size, nil
// }

// Idiomatic example
// func readIdiomatic(b *Buffer, p []byte) (n int, err error) {
// 	if b.empty() {
// 		b.Reset()
// 	}

// 	n = copy(p, b.buf[b.off:])
// 	b.off += n

// 	return n, nil
// }

/*
var s string // string
var i int // index
var num int // number
var msg string // message
var v string // value
var val string // value
var fv string //flag value
var err // error value
var args []string // arguments
var seen bool // has seen?
var parsed bool // parsing ok?
var buf []byte // buffer
var off int // offset
var op int // operation
var opRead int // read operation
var l int // length
var n int // number or number of
var m  int // another number
var c int // capacity
var c int // character
var a int // array 
var r rune // rune
var sep string // separator
var src int // source
var dst int // destination
var b byte // byte
var b []byte // buffer
var buf []byte // buffer
var w io.Writer // writer
var r io.Reader // reader
var pos int // position
...list goes on and on...
*/