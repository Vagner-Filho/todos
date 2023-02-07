package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	menu()

}

var menu_opts = [5]string{
	"Cadastrar To-do",
	"Editar To-do",
	"Deletar To-do",
	"Buscar To-do",
	"Ver Todos",
}

func menu() {
	fmt.Println("Selecione uma opção abaixo.")

	for i := 0; i < len(menu_opts); i++ {
		fmt.Printf("%v %v\n", i+1, menu_opts[i])
	}
}
