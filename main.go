package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	reader *bufio.Reader
)

func main() {
	//inicializa o reader
	reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("********************************")
		fmt.Println("* Seja bem vindo à Calculadora *")
		fmt.Println("********************************")
		fmt.Println("\nSelecione uma operação:")
		fmt.Println("1-Adição\n2-Subtração\n3-Multiplicação\n4-Divisão\n5-Sair")

		option := readLineString()

		if option == "5" {
			fmt.Println("Fim :)")
			os.Exit(0)
		}

		fmt.Println("Insira o primeiro valor:")
		value1, error := strconv.ParseFloat(readLineString(), 64)
		if error != nil {
			log.Printf("Ocorreu um erro ao capturar o primeiro valor")
			os.Exit(1)
		}

		fmt.Println("Insira o segundo valor:")
		value2, error := strconv.ParseFloat(readLineString(), 64)
		if error != nil {
			log.Printf("Ocorreu um erro ao capturar o segundo valor")
			os.Exit(1)
		}

		var result float64

		switch option {
		case "1":
			{
				println("Adicao:")
				result = value1 + value2
			}
		case "2":
			{
				println("Subtração:")
				result = value1 - value2
			}
		case "3":
			{
				println("Multiplicação:")
				result = value1 * value2
			}
		default:
			{
				println("Divisão:")
				if value2 == 0 {
					fmt.Println("Divisão por 0, não permitida")
					continue
				}
				result = value1 / value2
			}
		}

		fmt.Println("Resultado", result)
		fmt.Println("\nPrecione 'ENTER' para continuar... :)")
		readLineString()
	}

	fmt.Println("Fim :)")
}

func readLineString() string {
	line, error := reader.ReadString('\n')
	if error != nil {
		log.Printf("Ocorreu um erro ao capturar opção")
		os.Exit(1)
	}

	return strings.Replace(line, "\n", "", 1)
}
