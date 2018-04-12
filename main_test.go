package main

import (
	"flag"
	"testing"
)

func TestMain(t *testing.T) {
	flag.Set("file", "data.txt")
	main()
}
