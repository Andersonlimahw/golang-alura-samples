package main

import "fmt"

const (
	EXIT             = 0
	START_MONITORING = 1
	SHOW_LOGS        = 2
)

func main() {
	var name string = "Lemon"
	version := 1.1

	fmt.Println("Olá, Mr", name)
	fmt.Println("Este programa esta na version", version)

	showMenu()

	var command int
	fmt.Scan(&command)

	if !validateCommand(command) {
		fmt.Println("Comando não reconhecido")
		return
	}

	switch command {
	case EXIT:
		fmt.Println("Saindo do programa...")
	case START_MONITORING:
		fmt.Println("Iniciando monitoramento...")
	case SHOW_LOGS:
		fmt.Println("Exibindo logs...")
	default:
		fmt.Println("Comando não reconhecido")
	}
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

// validateCommand if sempre deve retornar true ou false e nao usa parenteses
func validateCommand(command int) bool {
	return command >= 0 && command <= 2
}
