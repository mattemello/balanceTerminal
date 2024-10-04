package errorhand

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

func SetLogFile() *os.File {
	var fileLog, errF = os.OpenFile("./tmp/logfile", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if errF != nil {
		log.Fatal("Error: ", errF, " -> can't open the log file")
		os.Exit(1)
	}

	return fileLog

}

func HandlerError(err error, txt string) {
	if err != nil {
		log.Println(" Error: ", err, " -> ", txt)

	}
}

func TakeFileLine() string {
	_, file, line, _ := runtime.Caller(1)

	return (file + " " + strconv.Itoa(line))
}

func BadSaving(err error) {
	if err != nil {
		log.Println("Error in the sqlite: ", err)
	}
}

func Controll(s time.Time) {
	log.Fatalln("item id (dropdown): ", s)
}
