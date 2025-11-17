package utils

import "time"

func GetDate() (string){
	date := time.Now().Format("2006-01-02")

	return  date
}