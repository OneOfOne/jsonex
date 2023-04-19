package jsonex_test

import (
	"fmt"
	"strings"
	"testing"

	"go.oneofone.dev/jsonex"
)

func TestDecodeItem(t *testing.T) {
	data := `[{"a":1},{"a":2},{"c":{"a":5}},5,"test","x",null]`
	dec := jsonex.NewDecoder(strings.NewReader(data))
	var out []any
	for dec.More() {
		var v any
		if err := dec.DecodeValue(&v); err != nil {
			t.Fatal(err)
		}
		out = append(out, v)
	}
	j, _ := jsonex.Marshal(out)
	if string(j) != data {
		t.Fatalf("expected %s, got %s", data, j)
	}

	dec = jsonex.NewDecoder(strings.NewReader(`{"foo": "bar", "baz": "qux"}`))

	// Decode the first key
	key, err := dec.DecodeKey()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(key))

	// Decode the first value
	var value string
	err = dec.DecodeValue(&value)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(value))

	// Decode the second key
	key, err = dec.DecodeKey()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(key))

	// Decode the second value
	err = dec.DecodeValue(&value)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(value))
}

func TestDecodeKeyValue(t *testing.T) {
	data := `{"a":1,"b":2,"c":{"a":5}}`
	dec := jsonex.NewDecoder(strings.NewReader(data))
	out := map[string]any{}
	for dec.More() {
		key, err := dec.DecodeKey()
		if err != nil {
			t.Fatal(err)
		}
		var v any
		if err := dec.DecodeValue(&v); err != nil {
			t.Fatal(err)
		}
		out[key] = v
		t.Logf("%q: %v", key, v)
	}

	j, _ := jsonex.Marshal(out)
	if string(j) != data {
		t.Fatalf("expected %s, got %s", data, j)
	}
}
