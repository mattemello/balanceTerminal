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

func BadSaving(err error) {
	if err != nil {
		log.Fatal("Error in the sqlite: ", err)
		os.Exit(1)
	}
}

func Controll(s string) {
	log.Fatalln("item id (dropdown): ", s)
	os.Exit(1)
}
