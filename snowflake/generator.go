package snowflake

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/snowflake"
)

func GenerateID() snowflake.ID {
	node := getNode()
	id := node.Generate()

	return id
}

func GenerateInt64() int64 {
	id := GenerateID()
	return id.Int64()
}

func GenerateBase36() string {
	s36 := strconv.FormatInt(GenerateInt64(), 36)
	return s36
}

func getNode() *snowflake.Node {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return node
}
