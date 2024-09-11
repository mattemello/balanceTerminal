package errorhand

import (
	"log"
	"os"
)

func HandlerError(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
		os.Exit(1)
	}
}

func Controll(s int) {
	log.Fatalln("item id (dropdown): ", s)
	os.Exit(1)
}
