package main

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/snowflake"
)

func main() {

	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate a snowflake ID.
	id := node.Generate()

	// Print out the ID in a few different ways.
	fmt.Printf("Int64  ID: %d\n", id)
	fmt.Printf("Base64 ID: %s\n", id.Base64())
	fmt.Printf("Base32 ID: %s\n", id.Base32())
	s36 := strconv.FormatInt(id.Int64(), 36)
	fmt.Printf("Base36 ID: %s\n", s36)
}
