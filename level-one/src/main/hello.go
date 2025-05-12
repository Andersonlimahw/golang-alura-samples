package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
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
		handleLogsReader()
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
	fileName := "sites"
	format := ".txt"
	fileNameValue := fileName + format

	createFile("logs", ".txt")
	readFileByName(fileNameValue)
	readFileByNameWithBufer(fileNameValue)
	now := time.Now().Format(time.DateTime)
	message := "Hello World, now is " + now
	writeFile("exemplo_escrita.txt", message)
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
	fileCreated.Close()
	return fileCreated
}

func readFileByName(name string) {
	file, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo", name, "Error -> ", err)
		return
	}
	fmt.Println("Arquivo aberto com sucesso", name, string(file))
}

func readFileByNameWithBufer(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo", name, "Error -> ", err)
		return
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err == io.EOF {
			break
		}
		fmt.Println("Linha", line)
	}

	file.Close()
}

func writeFile(name string, content string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo", name, "Error -> ", err)
		return
	}
	file.WriteString(content)
	file.Close()
	fmt.Println("Arquivo escrito com sucesso", name, content)
}

func handleLogsReader() {
	fmt.Println("handle logs reader")
	fileName := "logs"
	format := ".txt"
	createFile(fileName, format)
	writeFile(fileName+format, "Hello World, from handleLogsReader")
	readFileByNameWithBufer(fileName + format)
}
