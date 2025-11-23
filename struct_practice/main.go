package main

import (
	"bufio"
	"errors"
	"fmt"
	"go_practice/note"
	"os"
	"strings"
)

func getNoteData()(string,string,error){
	title,titleErr := getUserInput("Note tile : ")
	if titleErr != nil{
		fmt.Println(titleErr)
		return "","",titleErr
	}
	content,contentErr:=getUserInput("Note content : ")

	if contentErr != nil{
		fmt.Println(contentErr)
		return "","",contentErr
	}

	return title,content,nil
}

func main(){

	title, content, err := getNoteData()

	if err != nil{
		fmt.Println(err)
		return
	}
	
	note := note.New(title,content)

	note.Display()
	err = note.Save()


	if err != nil{
		fmt.Println("Saving the note failed", err)
		return
	}

}

func getUserInput(prompt string)(string,error){
	fmt.Print(prompt)
	
	reader :=bufio.NewReader(os.Stdin)
	value, err :=reader.ReadString('\n')
	if err != nil{
		return "", errors.New("the value should not be empty")
	}
	value = strings.TrimSuffix(value,"\n")
	value = strings.TrimSuffix(value,"\r")
	return value, nil
}