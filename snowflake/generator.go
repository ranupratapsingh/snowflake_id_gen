package snowflake

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	snowflake "github.com/bwmarrin/snowflake"
)

var nodeId int64 = -1

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
	if nodeId < 0 {
		nodeId = getNodeId()
		fmt.Println("setting nodeId", nodeId)
	}
	node, err := snowflake.NewNode(nodeId)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return node
}

// node id is a uniq number from 0 to 255
// it is set by the env variable HOSTNAME
// the HOSTNAME is automatically set by k8s statefulset
func getNodeId() int64 {
	hostName := os.Getenv("HOSTNAME")
	if hostName == "" {
		return 0
	}

	podIndex, _ := strings.CutPrefix(hostName, "snowflake-id-gen-")
	intIndex, err := strconv.Atoi(podIndex)

	if err != nil {
		return 0
	}
	return int64(intIndex)
}
