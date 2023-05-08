package main

import "picket/src/cmd"

func main() {
	root := cmd.Root()

	if err := root.Execute(); err != nil {
		panic(err)
	}
}
