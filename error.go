package main

import "log"

func IfError(err error, message string) {
	if err != nil {
		log.Println("/~~>\n message:",message,"\n error:",err.Error(),"\n ~~/")
	}
} 