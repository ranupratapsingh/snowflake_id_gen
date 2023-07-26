package main

import (
	"fmt"

	"github.com/ranupratapsingh/snowflake_id_k8s/snowflake"
)

func main() {
	id := snowflake.GenerateID()
	s36 := snowflake.GenerateBase36()

	// Print out the ID in a few different ways.
	fmt.Printf("Int64  ID: %d\n", id)
	fmt.Printf("Base64 ID: %s\n", id.Base64())
	fmt.Printf("Base32 ID: %s\n", id.Base32())
	fmt.Printf("Base36 ID: %s\n", s36)
}
