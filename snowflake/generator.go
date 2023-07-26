package snowflake

import (
	"fmt"
	"strconv"

	snowflake "github.com/bwmarrin/snowflake"
)

func init() {
	snowflake.Epoch = 1577836800000 // 2020-01-01 00:00:00 UTC
	snowflake.NodeBits = 8
	snowflake.StepBits = 14
}

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
	// todo: get node from config
	node, err := snowflake.NewNode(0)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return node
}
