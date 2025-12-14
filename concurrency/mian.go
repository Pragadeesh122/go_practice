package main

import (
	"fmt"
	"time"
)



func greet(value string,index int, channel chan bool){
	fmt.Println(value,index)
	channel <- true
}


func slowGreet(value string,index int, channel chan bool){
	time.Sleep(3* time.Second)
	fmt.Println(value,index)
	channel <- true
}

func main(){

	greetChannel := make(chan bool, 6)

	for i :=0; i < 4; i++{
		go greet("Hello Go Module",i,greetChannel)
	}

	for i :=0; i < 2; i++{
		go slowGreet("Slow Greetings",i,greetChannel)
	}


	for range 6{
		<-greetChannel
	}
}

