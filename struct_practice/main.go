package main

import (
	"bufio"
	"errors"
	"fmt"
	"go_practice/note"
	"go_practice/todo"
	"os"
	"strings"
)


type Saver interface{
	Save() error
}

type OutPutable interface{
	Saver
	Display()
}

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

	todoText,err := getUserInput("Todo Text : ")

	if err != nil{
		fmt.Println(err)
		return
	}

	todo := todo.New(todoText)
	note := note.New(title,content)
	
	outputData(note)
	outputData(todo)


	res := add(7,4324)
	fmt.Println(res)

}

func outputData(data OutPutable) error{
	data.Display()
	
	return saveData(data)

}

func saveData(data Saver) error{
	err := data.Save()

	if err != nil{
		fmt.Println("Saving the todo failed", err)
		return err
	}
	fmt.Println("The todo file was successfully created")
	return nil
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



func add[T int|float64|string](a,b T) T {
	return a+b
}