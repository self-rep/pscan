package main

/*

PSCAN - {
	Fast Portscanner By Selfrep#6192
	ig: @h4_remiixx
	https://github.com/self-rep
}

*/
import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func checkport(ip string, port int, ms time.Duration) {

	dialer := net.Dialer{
		Timeout: time.Millisecond * ms, // configure if no ports are showing (1000 = 1 sec) (100 = 1/10 second)
	}

	prt := strconv.Itoa(port)
	conn, err := dialer.Dial("tcp", ip+":"+prt)
	if err != nil {
		//fmt.Println("No Open Port Found", prt) // debug
		return
	}
	fmt.Println("\x1b[38;5;111m[\x1b[38;5;93mPSCAN\x1b[38;5;111m]\x1b[38;5;111m Found:\x1b[38;5;93m", ip, "\x1b[38;5;111m:\x1b[38;5;93m", port)
	conn.Close()

}

func main() {
	var threads int = 0       // default
	var ip string = "0.0.0.0" // default
	var ms time.Duration = 0

	args := os.Args[1:]

	for i := range args {
		if len(args[i]) > 1 && args[i][0] == '-' {

			kv := strings.Split(args[i][1:], "=")
			switch kv[0] {
			case "threads":
				amount, err := strconv.Atoi(kv[1])
				if err != nil {
					fmt.Println("Invalid Port")
					return
				}
				threads = amount
				break
			case "ms":
				mms, err := strconv.Atoi(kv[1]) // convert int to string
				if err != nil {
					fmt.Println("Invalid Timeout Value Make sure your value is an int")
					return
				}
				ms = time.Duration(mms) * time.Millisecond // ms * milisecond (1000 = 1 second, 100 = 1/10 second)
				break
			}

			continue
		}

		ip = args[i]
	}

	if threads == 0 || ip == "0.0.0.0" || ms == 0 {
		fmt.Println("Invalid Options!")
		fmt.Println("./pscan <ip> -threads=<threads> -ms=<MS>")
		fmt.Println("The more threads you use the more inaccurate it could be i normally use 1000 - 10000 threads")
		fmt.Println("Default MS is 1000(1 second) the lower you put the faster it will scan, if you have it to low it might not scan properly")
		return
	}
	tm := time.Now()
	fmt.Println("Starting Time:", tm)
	var j int = 65535

	fmt.Println("")
	fmt.Println("\x1b[38;5;111m[\x1b[38;5;91mNOTICE\x1b[38;5;111m] Scanning Ip: "+ip+" With Threads(", threads, ") And Timeout(", ms, ")")
	var wg sync.WaitGroup

	for i := 0; i < j; i++ {
		wg.Add(1)
		go func(i int) {
			checkport(ip, i, ms)
			wg.Done() // finish goroutine
		}(i)
		if i%threads == 0 {
			wg.Wait()
		}
	}
	tn := time.Since(tm)
	fmt.Println("\x1b[38;5;93m[\x1b[38;5;111mPSCAN\x1b[38;5;93m] \x1b[38;5;111mFinished \x1b[38;5;111mScan")
	fmt.Println("Ending Time:", tn)

}
