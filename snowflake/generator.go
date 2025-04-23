package snowflake

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	snowflake "github.com/bwmarrin/snowflake"
)

var currentNode *snowflake.Node = nil

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
	// fmt.Println("generated id", id)
	return id.Int64()
}

func GenerateBase36() string {
	s36 := strconv.FormatInt(GenerateInt64(), 36)
	return s36
}

func getNode() *snowflake.Node {
	if currentNode != nil {
		return currentNode
	}
	nodeId := getNodeId()
	newNode, err := snowflake.NewNode(nodeId)
	currentNode = newNode
	fmt.Println("createing new node with Id", nodeId)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return currentNode
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
