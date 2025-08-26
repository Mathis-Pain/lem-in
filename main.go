package main

import (
	"lem-in/utils"
	"lem-in/utils/builder"
	"os"
)

func main() {
	content := utils.FileMaker(os.Args)
	if content == nil {
		return
	}
	defer content.Close()

	AllRooms, Links := builder.ParseFile(content)

}
