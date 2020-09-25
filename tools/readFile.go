package tools

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

//read domain file
func ReadDomainFile(domainFile string) ([]string, error) {
	var domainList []string

	domain, err := os.Open(domainFile)
	if err != nil {
		log.Printf("Cannot open text file: %s, err: [%v]", domainFile, err)
		return nil, err
	}
	defer domain.Close()

	b := bufio.NewReader(domain)
	for {
		line, err := b.ReadString('\n')
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			domainList = append(domainList, line)
		}
		if err != nil {
			if err == io.EOF {
				log.Println("Read File Done.")
				return domainList, nil
			}
			log.Println("Read File Error:")
			return nil, err
		}
	}
}
