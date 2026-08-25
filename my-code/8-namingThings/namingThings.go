package namingthings

// Why is naming important?
// Critial for Readability = Maintainability

// Tip: Use the first few letters of the words
// Tip: Use fewer letters in smaller scope
// Tip: Use the complete words in larger scope (like package names)
// Tip: Use mixedCaps like this (type PlayerScore struct; here if 'PlayerScore' is being exported then this mixedCaps pattern is fine)
// Tip: Use all capitals for acronyms (like var localAPI string)
// Tip: Do not stutter (dont use player.PlayerScore, here 'player' word is repeated. But use player.Score instead.)
// Tip: Do not use under_scores in names. (MAX_TIME - bad, MaxTime - good, N - good)

func ReadNonIdiomatic(buffer *Buffer, inBuffer []byte) (size int, err error) {
	if buffer.empty() {
		buffer.Reset()
	}

	size = copy(
		inBuffer,
		buffer.buffer[buffer.offset:])

	buffer.offset += size
	return size, nil
}

func ReadIdiomatic(b *Buffer, p []byte) (n int, err error) {
	if b.empty() {
		b.Reset()
	}

	n = copy(p, b.b[b.off:])

	b.off += n

	return n, nil
}

// Common abbreviations in Go
var s string // string
var i int // index
var num int // number
var msg string // message
var v string // value
var val string // value
var fv string // flag value
var err error // error value
var args []string // arguments
var seen bool // has seen?
var parsed bool // parsed ok?
var buf []byte // buffer
var off int // offset
var op int // operation
var opRead int // read operation
var l int // length
var n int // number or number of
var m int // another number
var c int // capacity
var a int // array
var r rune // rune (aka int32)
var sep string // seperator
var src int // source
var dst int // destination
var b byte // byte
var b []byte // buffer
var buf [] byte // buffer
var w io.Writer // writer
var r io.Reader // reader
var pos int // position