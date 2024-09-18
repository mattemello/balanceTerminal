package errorhand

import (
	"log"
	"os"
	"runtime"
	"strconv"
)

func HandlerError(err error, txt string) {
	if err != nil {
		log.Fatal("Error: ", err, " -> ", txt)
		os.Exit(1)
	}
}

func TakeFileLine() string {
	_, file, line, _ := runtime.Caller(1)

	return (file + " " + strconv.Itoa(line))
}

func BadSaving(err error) {
	if err != nil {
		log.Fatal("Error in the sqlite: ", err)
		os.Exit(1)
	}
}

func Controll(s int) {
	log.Fatalln("item id (dropdown): ", s)
	os.Exit(1)
}
