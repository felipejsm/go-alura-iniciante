package main

import "fmt"

func main() {
	nome := "Felipe"
	idade := 30
	versao := 1.1
	fmt.Println("Olá Sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

	var comando int
	//fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)
}
