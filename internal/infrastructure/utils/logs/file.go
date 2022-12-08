package logs

import (
	"fmt"
	"os"
	"time"
)

func AddLogFile(message string, typeLog string) {
	f, _ := os.OpenFile("./logs-"+typeLog+".brvlog", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	_, err := f.Write([]byte(fmt.Sprintf("%s ..->.. %s\n", time.Now().Format("2006-01-02 15:04:05"), message)))
	if err != nil {
		return
	}
}
