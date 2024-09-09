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

func Controll(s []string) {
	log.Fatalln("dimension: ", len(s))
	for i, c := range s {
		log.Fatal(i, " --- Value: ", c)
	}
	os.Exit(1)
}
