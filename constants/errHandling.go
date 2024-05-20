package constants

import (
	"fmt"
	"log"
	"os"
)

func CheckErr(code int, msg string, err error) {
	switch code {
	case 200:
		{
			log.Println("INFO: " + msg + "\n")
		}
	case 300:
		{
			if err != nil {
				writeLog(fmt.Sprintf("FAILED: %v\n", msg+": "+err.Error()))
			}
		}
	case 500:
		{
			if err != nil {
				writeLog(fmt.Sprintf("ERROR: %v", msg+": "+err.Error()))
			}
		}
	}
}

func writeLog(line string) {

	f, _ := os.OpenFile("./error.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	_, _ = f.Write([]byte(line))
	_ = f.Close()
}
