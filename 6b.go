package main

import (
	"reflect"
	"strings"
)

func main() {
	fileContent := readFileReturnArray("./inputs/6.input")

	packet := fileContent[0]

	packetArr := strings.Split(packet, "")

	uniqueSequenceLen := 14

	for idx, _ := range packetArr {

		set := make(map[string]bool)

		for i := idx; i < idx+uniqueSequenceLen; i++ {
			set[packetArr[i]] = true
		}

		if len(reflect.ValueOf(set).MapKeys()) == uniqueSequenceLen {
			print(idx + uniqueSequenceLen)
			return
		}
	}
}
