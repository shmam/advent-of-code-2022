package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	parent   *Node
	fileSize int
	children map[string]*Node
}

type FileTree struct {
	root *Node
}

const threshold = 100000

func treeTraversal(n *Node) int {
	if n == nil {
		return 0
	}

	sum := 0
	if n.fileSize <= threshold {
		sum = n.fileSize
	}
	for _, child := range n.children {
		sum += treeTraversal(child)
	}
	return sum
}

func main() {

	fileContent := readFileReturnArray("./inputs/7.input")

	commands := [][]string{}
	currentCommand := []string{}

	for _, eachLine := range fileContent {
		if eachLine[0:1] == "$" {
			commands = append(commands, currentCommand)
			currentCommand = []string{eachLine}
		} else {
			currentCommand = append(currentCommand, eachLine)
		}
	}

	tree := FileTree{}

	tree.root = &Node{
		fileSize: 0,
		children: map[string]*Node{},
	}

	currentNode := tree.root

	// build tree from commands
	for _, eachCommand := range commands {

		if len(eachCommand) > 0 {
			command := eachCommand[0]
			results := eachCommand[1:]

			commandParts := strings.Split(command, " ")

			if commandParts[1] == "cd" {
				// move positon to another directory
				targetDirectory := commandParts[2]

				if targetDirectory == ".." {
					currentNode = currentNode.parent
				} else if targetDirectory == "/" {
					currentNode = tree.root
				} else {
					currentNode = currentNode.children[targetDirectory]
				}

			} else if commandParts[1] == "ls" {

				// list files and dirs
				for _, eachResult := range results {
					lsResult := strings.Split(eachResult, " ")

					// creating new directories
					if lsResult[0] == "dir" {
						dirName := lsResult[1]

						currentNode.children[dirName] = &Node{
							parent:   currentNode,
							fileSize: 0,
							children: map[string]*Node{},
						}

					} else {
						// listing files
						fileSize, _ := strconv.Atoi(lsResult[0])

						currentNode.fileSize += fileSize

						tempNode := currentNode.parent
						for tempNode != nil {
							tempNode.fileSize += fileSize
							tempNode = tempNode.parent
						}
					}
				}
			}

		}
	}

	// recursively traverse tree
	currentNode = tree.root
	sum := treeTraversal(currentNode)

	fmt.Println(sum)
}
