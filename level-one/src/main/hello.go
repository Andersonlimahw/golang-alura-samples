package main

import (
	"fmt"
	"time"
)
import "os"
import "net/http"

const (
	Exit            = 0
	StartMonitoring = 1
	ShowLogs        = 2
	ShowSamples     = 3
	HandleFiles     = 4
)

const (
	DELAY_IN_SECONDS    = 3
	MAX_NUMBER_OF_SITES = 2
)

func main() {
	for { // loop infinito pois no go não temos while
		showIntro()
		showMenu()
		var command int = getCommand()
		if !validateCommand(command) {
			fmt.Println("Comando não reconhecido")
			return
		}
		handleCommand(command)
	}

}

func showIntro() {
	name, _ := getPerson() // o _ ignora o valor retornadothu
	version := 1.1

	fmt.Println("Olá, Mr", name)
	fmt.Println("Este programa esta na version", version)

}

func showMenu() {
	fmt.Println("0- Sair do Programa")
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- Executar exemplos")
	fmt.Println("4- Ler arquivos")
}

func getCommand() int {
	var command int
	fmt.Scan(&command)
	return command
}

// validateCommand if sempre deve retornar true ou false e nao usa parenteses
func validateCommand(command int) bool {
	return command >= 0 && command <= 4
}

func handleCommand(command int) {
	switch command {
	case Exit:
		fmt.Println("Saindo do programa...")
		os.Exit(0) // Finaliza o programa com sucesso, para erro é -1
	case StartMonitoring:
		startMonitoring()
	case ShowLogs:
		fmt.Println("Exibindo logs...")
	case ShowSamples:
		fmt.Println("Exibindo exemplos	...")
		handleSamples()
	case HandleFiles:
		fmt.Println("Lendo arquivos...")
		handleFiles()
	default:
		fmt.Println("Comando não reconhecido")
		os.Exit(-1)
	}
}

func startMonitoring() {
	fmt.Println("Iniciando monitoramento...")
	site := "https://www.google.com"
	response, err := http.Get(site)
	if err != nil {
		fmt.Println("Erro ao acessar o site", site, "Error -> ", err)
	}
	fmt.Println("Monitoramento finalizado..., response", response)
}

func getPerson() (string, int) {
	name := "Lemon"
	age := 22

	return name, age
}
func handleSamples() {
	sliceSample()
	addElement("Mickey")
	pokerSamples()
	forSample()
	forWithRangeSample()
	delaySample()
}

func sliceSample() {
	// Usamos Slice pq o tamanho da lista pode mudar e se
	//usarmos um array o valor limit deve se informado  exemplo varr sites[4]string
	// tudo é um array no Go
	fmt.Println("slice sample")
	slice := []string{"Anderson", "Lemon"}
	fmt.Println(slice)
	fmt.Println("Tamanho", len(slice))
	fmt.Println("Capacidade", cap(slice))
}

func addElement(name string) {
	fmt.Println("slice sample : add element")
	slice := []string{"Anderson", "Lemon", name}
	fmt.Println(slice)
	slice = append(slice, name)
	fmt.Println("Slice atualizado com o novo elemento", name)
	fmt.Println(slice)
}

func pokerSamples() {
	fmt.Println("slice sample : capacity")
	pontosPlanningPoker := []int{1, 2, 3, 5, 8, 13, 21}
	pontosPlanningPoker = append(pontosPlanningPoker, 40)
	fmt.Println(cap(pontosPlanningPoker))
}

func forSample() {
	fmt.Println("for sample")
	for i := 0; i < MAX_NUMBER_OF_SITES; i++ {
		fmt.Println(i)
	}
}

func forWithRangeSample() {
	fmt.Println("for with range sample")
	pontosPlanningPoker := []int{1, 2, 3, 5, 8, 13, 21}
	for i, ponto := range pontosPlanningPoker {
		fmt.Println("Indice", i, "Ponto", ponto)
	}
}

func delaySample() {
	fmt.Println("delay sample")
	delay := DELAY_IN_SECONDS * time.Second
	time.Sleep(delay)
	fmt.Println("delay sample", delay, "seconds")
}

func handleFiles() {
	fmt.Println("handle files")
	fileCreated := createFile("testes", ".md")
	if fileCreated != nil {
		openFileByName("testes.md")
	}
}

func createFile(name string, format string) *os.File {
	fmt.Println("create file")
	fileName := name + format
	fileCreated, errFileCreated := os.Create(fileName)
	if errFileCreated != nil {
		fmt.Println("Erro ao criar o arquivo", fileName, "Error -> ", errFileCreated)
		return nil
	}
	fmt.Println("Arquivo criado com sucesso", fileName)
	return fileCreated
}

func openFileByName(name string) **os.File {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo", name, "Error -> ", err)
		return nil
	}
	fmt.Println("Arquivo aberto com sucesso", name, file)
	return &file
}
