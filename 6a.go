package main

import (
	"reflect"
	"strings"
)

func main() {
	fileContent := readFileReturnArray("./inputs/6.input")

	packet := fileContent[0]

	packetArr := strings.Split(packet, "")

	for idx, _ := range packetArr {

		set := make(map[string]bool)

		for i := idx; i < idx+4; i++ {
			set[packetArr[i]] = true
		}

		if len(reflect.ValueOf(set).MapKeys()) == 4 {
			print(idx + 4)
			return
		}
	}
}
