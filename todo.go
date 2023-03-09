package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

var reader = bufio.NewReader(os.Stdin)

func get_cli_input() string {
	input, err := reader.ReadString('\n')

	if err == nil {
		return strings.Replace(input, "\n", "", -1)
	}
	fmt.Println(err.Error())
	return "none"
}

func main() {
	fmt.Println("Olá, por favor insira seu nome: ")
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')

	if len(username) <= 1 {
		fmt.Println("Por favor, insira um nome válido")
		return
	} else if err != nil {
		fmt.Println(err.Error())
	}

	option := 0
	for option != 6 {
		menu()
		option = get_menu_choosen_option()
		handleOption(option)
	}
}

var menu_opts = [6]string{
	"Cadastrar To-do",
	"Editar To-do",
	"Deletar To-do",
	"Buscar To-do",
	"Ver Todos",
	"Sair",
}

type Todo struct {
	done bool
	id   int
	text string
}

var todos = []Todo{}

func menu() {
	fmt.Println("Selecione uma opção abaixo.")

	for i := 0; i < len(menu_opts); i++ {
		fmt.Printf("%v %v\n", i+1, menu_opts[i])
	}
}

func get_menu_choosen_option() int {

	reader := bufio.NewReader(os.Stdin)
	opt, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	value, err := strconv.Atoi(opt[:len(opt)-1])
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	return value
}

func printTodos() {
	for v := range todos {
		fmt.Printf("%v. %v\n", todos[v].id, todos[v].text)
	}
}

func getTodoTextFromUser() (Todo, error) {
	fmt.Println("Todo: ")
	reader := bufio.NewReader(os.Stdin)
	todo, err := reader.ReadString('\n')

	newTodo := Todo{}

	if err == nil {
		newTodo := Todo{
			done: false,
			id:   len(todos),
			text: todo,
		}
		return newTodo, err
	} else {
		return newTodo, err
	}
}

func handleOption(opt int) {
	if opt == 1 {
		fmt.Println("Escreva o Todo: ")
		reader := bufio.NewReader(os.Stdin)
		todo, err := reader.ReadString('\n')

		if err == nil && len(todo) > 5 {
			var newTodo = Todo{
				done: false,
				id:   len(todos),
				text: todo,
			}
			todos = append(todos, newTodo)
		} else {
			fmt.Println(err.Error())
		}
	}
	if opt == 2 {
		fmt.Println("Digite o id de um Todo: ")
		printTodos()
		todo_id := get_cli_input()

		if todo_id == "none" {
			fmt.Println("Nenhum todo escolhido")
			return
		}

		id, conversion_err := strconv.Atoi(todo_id)

		if conversion_err != nil {
			fmt.Println(conversion_err.Error())
			return
		}

		edited_todo_text, err := getTodoTextFromUser()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		todos[id].text = edited_todo_text.text
	}
	if opt == 3 {
		printTodos()
		fmt.Println("Digite o Id do Todo que deseja deletar: ")
		id := get_cli_input()

		if id == "none" {
			return
		}

		intId, conversion_err := strconv.Atoi(id)
		if conversion_err != nil {
			fmt.Println(conversion_err.Error())
			return
		}
		todos = slices.Delete(todos, intId, intId)
	}
	if opt == 4 {
		fmt.Println("Insira a busca:")
		search := get_cli_input()

		if search == "none" {
			return
		}

		for _, td := range todos {
			if strings.Contains(td.text, search) {
				fmt.Printf("%v. %v\n", td.id, td.text)
			}
		}
	}
	if opt == 5 {
		printTodos()
	}
	// TODO: handle options and integrate with db
}
