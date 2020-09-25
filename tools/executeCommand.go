package tools

import (
	"log"
	"os/exec"
	"strconv"
)

////Execute Command
//func ExecuteCommand(domain string) int {
//	cmd := exec.Command("dig", "@10.42.1.19", domain)
//
//	var out bytes.Buffer
//	var stderr bytes.Buffer
//	cmd.Stdout = &out
//	cmd.Stderr = &stderr
//
//	var times int
//	err := cmd.Run()
//	if err != nil {
//		log.Println("stderr:", err.Error())
//	} else {
//		time := GetBetweenStr(out.String(), "Query time: ", " ")
//		times, err = strconv.Atoi(time)
//		if err != nil {
//			log.Println("Atoi error: ", err)
//		}
//
//	}
//	return times
//}
//阻塞
func ExecuteCommand(domain string) int {
	cmd := exec.Command("dig", "@10.42.1.19", domain)

	var times int
	output, err := cmd.Output()
	if err != nil {
		log.Println("stderr: packet loss ==> domain =", domain, err.Error())
	} else {
		time := GetBetweenStr(string(output), "Query time: ", " ")
		times, err = strconv.Atoi(time)
		if err != nil {
			log.Println("Atoi error: ", err)
		}

	}
	return times
}
