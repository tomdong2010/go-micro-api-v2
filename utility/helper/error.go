package helper

import (
	"log"
)

// 检查错误
func CheckErr(msg string, err error, exit bool) {
	if err == nil {
		log.Printf("%s Finish!!!\n", msg)
		return
	}

	if exit {
		log.Fatalf("%s Error: %v\n", msg, err)
	} else {
		log.Printf("%s Error: %v\n", msg, err)
	}
}
