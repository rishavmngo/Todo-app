package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"time"
	// "reflect"
)

type Todo struct {
	Name      string
	Complete  bool
	Timestamp time.Time
}

func main() {

	welcome()
	var command = 1
	var todos []Todo

	//if error reading data
	if !getData(&todos) {
		return
	}

	printMenu()
	for command != 0 {

		color.Set(color.FgMagenta)
		fmt.Print("Enter the command(0-9): ")
		color.Unset()
		fmt.Scan(&command)

		switch command {
		case 1:
			printListOfTodos(todos)
		case 2:
			addTodo(&todos)
			writeData(&todos)
		case 3:
			printListOfTodos(todos)
			toggleStatus(todos)
			writeData(&todos)
		case 4:
			printListOfTodos(todos)
			updateTodo(todos)
			writeData(&todos)
		case 5:
			listCompletedTodos(&todos)
		case 6:
			listUnCompleteTodos(&todos)
		case 8:
			clear()
		case 9:
			printMenu()
		case 0:
			break
		default:
			fmt.Println("Not a valid command!")
			printMenu()
		}
		fmt.Print("\n\n\n")
	}
}

func updateTodo(todos []Todo) {
	var number uint
	fmt.Println("Update a todo")
	fmt.Print("Enter a todo number: ")
	fmt.Scan(&number)

	if number < 1 || number > uint(len(todos)) {
		fmt.Println("Invalid todo list number! Try Again")
		return
	}
	fmt.Print("New Name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	todos[number-1].Name = scanner.Text()
	todos[number-1].Timestamp = time.Now()
}

func toggleStatus(todos []Todo) {
	var number uint
	fmt.Println("Toggle status of a todo")
	fmt.Print("Todo number: ")
	fmt.Scan(&number)
	if number < 1 || number > uint(len(todos)) {
		fmt.Println("Invalid todo list number! Try Again")
		return
	}

	todos[number-1].Complete = !todos[number-1].Complete

}

func printMenu() {
	color.Yellow("Commands")
	fmt.Println("\t(0 - exit, 1 - list all todos, 2 - add, 3 - toggle status, 4 - update,5 - list completed todos, 6 - list uncompleted todos, 8 - clear screen, 9 - list commands)")
}

func addTodo(todos *[]Todo) {

	fmt.Println("New Todo")
	fmt.Print("Name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	*todos = append(*todos, Todo{Name: scanner.Text(), Timestamp: time.Now(), Complete: false})
}

func printListOfTodos(todos []Todo) {
	for index, todo := range todos {
		statusSign(todo.Complete)
		fmt.Println(strconv.Itoa(index+1)+".", todo.Name)
		fmt.Printf("   %v\n", todo.Timestamp.Format(time.ANSIC))

		color.Unset()
	}
}

func statusSign(status bool) {
	if status {
		color.Set(color.FgGreen)
	} else {
		color.Set(color.FgRed)
	}
}
func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func welcome() {
	color.Magenta("Welcome to todo list")
}

func getData(todos *[]Todo) bool {
	data, err := ioutil.ReadFile("data/todo.json")
	if err != nil {

		fmt.Println(err)
		return false
	}

	if !json.Valid(data) {
		fmt.Println("invalid JSON string:", string(data))
		return false
	}
	json.Unmarshal(data, todos)
	return true
}

func writeData(todos *[]Todo) {
	data, err := json.Marshal(*todos)
	if err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile("data/todo.json", data, 0777)
}

func listCompletedTodos(todos *[]Todo) {

	for index, todo := range *todos {
		if todo.Complete {
			fmt.Println(strconv.Itoa(index+1)+".", todo.Name)
			fmt.Printf("   %v\n", todo.Timestamp.Format(time.ANSIC))
		}
	}
}

func listUnCompleteTodos(todos *[]Todo) {

	for index, todo := range *todos {
		if !todo.Complete {
			fmt.Println(strconv.Itoa(index+1)+".", todo.Name)
			fmt.Printf("   %v\n", todo.Timestamp.Format(time.ANSIC))
		}
	}
}
