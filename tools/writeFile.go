package tools

import (
	"bufio"
	"log"
	"os"
)

func WriteFile(path, content string) {

	fn, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal("Open file failed: ", err)

	}
	defer fn.Close()
	writer := bufio.NewWriter(fn)
	_, _ = writer.WriteString(content)

	_ = writer.Flush()

}
