package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)


type Todo struct{
	content string
}

func (todo *Todo) Display(){
	fmt.Println(todo.content)
}

func (todo Todo) MarshalJSON()([]byte,error){
	return json.Marshal(
		struct{
		Content string	`json:"content"`
	}{
		Content: todo.content,
	})
}

func (todo Todo) Save()(error){
	fileName := "todo.json"

	json,err := json.Marshal(todo)

	if err != nil{
		return errors.New("there was an error converting the file to json")
	}

	return os.WriteFile(fileName,json,0644)
}


func New(content string)(*Todo){
	return &Todo{
		
		content: content,
		
	}
}

