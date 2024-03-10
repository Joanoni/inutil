package inutil

import (
	"fmt"
	"log"
)

func Print(values ...any) {
	fmt.Println(values...)
}

func PrintF(format string, values ...any) {
	fmt.Printf(format+"\n", values...)
}

func Log(values ...any) {
	log.Println(values...)
}

func LogF(format string, values ...any) {
	log.Printf(format+"\n", values...)
}
