package logs

import (
	"fmt"
	"os"
	"time"
)

func addLogFile(typeLog string, message string) {

	// Open the file
	f, _ := os.OpenFile("./"+time.Now().Format("2006-01-02")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	// Close the file when this function is finished
	defer f.Close()

	// Write in the file the log message
	f.Write([]byte(fmt.Sprintf("%s -> %s\n", time.Now().Format("2006-01-02 15:04:05"), message)))
}
