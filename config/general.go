package config

import (
	"fmt"
	"log"
)

func CatchError(e *error) {
	if err := recover(); err != nil {
		*e = fmt.Errorf("%v", err)
	}
}

func CatchErrorGeneral() {
	if err := recover(); err != nil {
		log.Println("Error =>", err)
	}
}
