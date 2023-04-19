# jsonex: A simple fork of `encoding/json`

This package is a fork of `encoding/json` that allows decoding individual JSON slice and map items. It is intended to be used in conjunction with `jsonex.Decoder` to allow decoding of large JSON files without loading the entire file into memory.

It adds 2 functions `DecodeKey` and `DecodeValue` that allow decoding of individual JSON slice and map items.

Synced with [Go](https://github.com/golang/go/tree/9cad0cc6e6b2a84134c46ce7069e62de28459f26) on 2023-04-19.

## Example

```go
	package main

	import (
		"fmt"
		"strings"

		"go.oneofone.dev/jsonex"
	)

	func main() {
		dec := jsonex.NewDecoder(strings.NewReader(`{"foo": "bar", "baz": "qux"}`))

		// Decode the first key
		key, err := dec.DecodeKey()
		if err != nil {
			panic(err)
		}
		fmt.Println(key)

		// Decode the first value
		var value string
		err = dec.DecodeValue(&value)
		if err != nil {
			panic(err)
		}
		fmt.Println(value)

		// Decode the second key
		key, err = dec.DecodeKey()
		if err != nil {
			panic(err)
		}
		fmt.Println(key)

		// Decode the second value
		err = dec.DecodeValue(&value)
		if err != nil {
			panic(err)
		}
		fmt.Println(value)
	}
```

More examples in [decode_ex_test.go](decode_ex_test.go).

## License

Same as Go's `encoding/json` package. See [LICENSE](LICENSE) for details.
