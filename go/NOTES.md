# Go notes

## String

a string is in effect a read-only slice of bytes.

## Slice

A slice is a data structure describing a contiguous section of an array stored separately from the slice variable itself. A slice is not an array. A slice describes a piece of an array.

## Converting a string into a byte array (and back)

```go
func main() {
	str := "foo"
	byteArray := []byte(str) // [102 111 111]

  backToString := string(byteArray[:]) // "foo"
}
```

## When to use panic