package main

import "fmt"



func main(){

	urls := map[string]string{
		"portfolio" : "https://pragadeeshvs.dev",
		"financetracker" : "https://financetracker.one",
	}

	
	for key,value := range urls{
		fmt.Println(key,value)
	}

}