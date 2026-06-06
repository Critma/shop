package main

import "github.com/critma/auth/internal/injection"

func main() {
	injection.BuildGraph().Run()
}
