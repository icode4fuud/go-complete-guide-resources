package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// by convention if the interface has only one method, the name of the interface should be the method name with -er suffix
// e.g. Saver interface
type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

// add an interface to accommodate both save and display methods but the better implementation is to embed the interfaces
//
//	type ouputtable interface {
//		Save() error
//		Display()
//	}
type ouputtable interface {
	saver // embed the saver interface
	Display()
}

// start main
func main() {

	//exemplify using printSomething function
	printSomething("Hello, World!")
	printSomething(42)
	printSomething(3.14159)

	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	//call todo Display and Save
	// todo.Display()
	// err = saveData(todo) // saveData(todo) replaces todo.Save()
	//^simplify the code by using outputtable
	err = outputData(todo)

	if err != nil {
		fmt.Println("Saving the todo failed.")
		return
	}

	// userNote.Display()
	// err = saveData(userNote) // saveData(userNote) replaces userNote.Save()
	//^simplify the code by using outputtable
	outputData(userNote)

	//not needed
	// if err != nil {
	// 	fmt.Println("Saving the note failed.")
	// 	return
	// }

}

//end main

// using Switch statement to determine the type of the value
func printSomething(value interface{}) {
	switch value.(type) {
	case string:
		fmt.Println("String:", value)
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float64:", value)
	}
	fmt.Println(value)
}

// add a method to replace duplication of Display method and Save method
func outputData(data ouputtable) error {
	data.Display()
	return saveData(data)
}

// add SaveData function to replace repetitive code
func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the data failed.")
		return err
	}

	fmt.Println("Saving the data succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
