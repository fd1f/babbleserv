package main

import (
	"os"
	"crypto/ed25519"
)

func main() {
	f, err := os.Create("key.key")
	if err != nil {
		panic(err)
	}
	_, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
	f.Write(priv)
	f.Close()
	// this was to see if it wrote
	// data, err := os.ReadFile("key.key")
	// if err != nil {
	// 	panic(err)
	// }
	os.Exit(0)
}
