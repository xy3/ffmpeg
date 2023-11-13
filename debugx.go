//go:build !debug
// +build !debug

package ffmpeg

func DebugNodes(node []DagNode) {}

func DebugOutGoingMap(node []DagNode, m map[int]map[Label][]NodeInfo) {}
