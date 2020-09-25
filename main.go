package main

import (
	"fmt"
	"log"

	"tdns0.2/tools"
)

var (
	domainFile = "./domain.txt"
	//wg         sync.WaitGroup
	//lock       sync.Mutex
)

func main() {

	domain, err := tools.ReadDomainFile(domainFile)
	if err != nil {
		log.Fatal("read domain failed:", err)
	}
	//wg.Add(2)
	get10(domain)
	get100(domain)
	//wg.Wait()

	fmt.Println("Work Done.")
}

func get10(domain []string) {
	for i := 0; i < len(domain); i++ {
		var (
			name   = domain[i]
			failed float64
			count  = 0
		)
		for j := 0; j < 10; j++ {
			time := tools.ExecuteCommand(name)
			if time == 0 {
				failed += 1
			}
			count += time
			tools.WriteFile("get10", fmt.Sprintf("%d ", time))
		}
		tools.WriteFile("get10", fmt.Sprintf("\n failed==> %.2f %%  average==> %.2f msec  domain==> %s \n", failed*10, float64(count)/(10.0-failed), name))
	}
	//wg.Done()
}

func get100(domain []string) {
	for i := 0; i < len(domain); i++ {
		var (
			name   = domain[i]
			failed float64
			count  = 0
		)
		for k := 0; k < 100; k++ {
			time := tools.ExecuteCommand(name)
			if time == 0 {
				failed += 1
			}
			count += time
			tools.WriteFile("get100", fmt.Sprintf("%d ", time))
		}
		tools.WriteFile("get100", fmt.Sprintf("\n failed==> %.2f %%  average==> %.2f msec  domain==> %s \n", failed, float64(count)/(100.0-failed), name))
	}
	//wg.Done()
}
