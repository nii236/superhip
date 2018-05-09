package main

import (
	"encoding/json"
	"io"
)

func mustDecode(r io.Reader, target interface{}) {
	err := json.NewDecoder(r).Decode(target)
	if err != nil {
		panic(err)
	}
}

func mustMarshal(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}
