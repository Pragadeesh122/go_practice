package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)



type Note struct{
	title string
	content string
	createdAt time.Time
}


func (note *Note) Display(){
	fmt.Printf("Your note title %v has the following content:\n\n%v\n", note.title,note.content)
}

func (note Note) MarshalJSON()([]byte,error){
	return json.Marshal(
		struct{
		Title string  `json:"title"`        
		Content string	`json:"content"`
		CreatedAt time.Time `json:"created_at"`
	}{
		Title: note.title,
		Content: note.content,
		CreatedAt: note.createdAt,
	})
}

func (note Note) Save()(error){
	fileName := strings.ReplaceAll(note.title," ","_")
	fileName = strings.ToLower(fileName)+".json"


	json,err := json.Marshal(note)

	if err != nil{
		return errors.New("there was an error converting the file to json")
	}

	return os.WriteFile(fileName,json,0644)
}


func New(title string,content string)(*Note){
	return &Note{
		title: title,
		content: content,
		createdAt: time.Now(),
	}
}

