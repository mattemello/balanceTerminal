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

func Controll(s rune) {
	log.Fatalln("item id (dropdown): ", s)
	os.Exit(1)
}
