package main

import "fmt"
import "os"

const (
	EXIT             = 0
	START_MONITORING = 1
	SHOW_LOGS        = 2
)

func main() {
	showIntro()
	showMenu()
	var command int = getCommand()
	if !validateCommand(command) {
		fmt.Println("Comando não reconhecido")
		return
	}
	handleCommand(command)
}

func showIntro() {
	var name string = "Lemon"
	version := 1.1

	fmt.Println("Olá, Mr", name)
	fmt.Println("Este programa esta na version", version)

}

func showMenu() {
	fmt.Println("0- Sair do Programa")
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
}

func getCommand() int {
	var command int
	fmt.Scan(&command)
	return command
}

// validateCommand if sempre deve retornar true ou false e nao usa parenteses
func validateCommand(command int) bool {
	return command >= 0 && command <= 2
}

func handleCommand(command int) {
	switch command {
	case EXIT:
		fmt.Println("Saindo do programa...")
		os.Exit(0) // Finaliza o programa com sucesso, para erro é -1
	case START_MONITORING:
		fmt.Println("Iniciando monitoramento...")
	case SHOW_LOGS:
		fmt.Println("Exibindo logs...")
	default:
		fmt.Println("Comando não reconhecido")
		os.Exit(-1)
	}
}
