package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	fmt.Println("MENU")
	fmt.Println("ecreva P para acrescentar uma nova palavra")
	fmt.Println("escreva J para jogar ")
	fmt.Println("escreva S para finalizar")
	var res string
	fmt.Scanln(&res)
	switch res {
	case "P":
		add()
	case "J":
		gameMenu()
	case "S":
		final()
	default:
		fmt.Println("ESCREVA DIREITO")
		main()
	}
}

func list() []string {
	f, err := os.Open("data.txt")
	words := []string{}
	if err != nil {

		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {

		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return words
}

func add() {
	words := list()
	var newWord string
	fmt.Println("qual palavra vc deseja adicionar ?")
	fmt.Scanln(&newWord)
	if len(newWord) > 0 {

		for _, v := range words {
			if v == newWord {
				fmt.Println("Essa palavra ja existe")
				main()
			}
		}
	} else {
		fmt.Println("DIGITA ALGO ANTES DE ADD")

		add()
	}

	Write(newWord)
}
func Write(str string) {
	file, err := os.OpenFile("data.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	_, err = file.WriteString("\n" + str)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	fmt.Printf("\nA palavra %v foi devidamente adicionada no arquivo", str)
	fmt.Printf(" no arquivo: %s \n", file.Name())
	main()
}

func gameMenu() {
	var vida int
	fmt.Println("em qual dificuldade deseja jogar?")
	fmt.Println("Dificil press 10")
	fmt.Println("Medio press 12")
	fmt.Println("Facil press 24")
	fmt.Scanln(&vida)
	switch vida {
	case 10:
		fmt.Println("que aldacia vc tem")
		game(vida)
	case 12:
		fmt.Println("uma escolha razoavel")
		game(vida)
	case 24:
		fmt.Println("vc gosta de tudo facil")
		game(vida)
	default:
		fmt.Println("DEIXA DE SER BURRO")
		gameMenu()
	}

}

var jogadas = []string{}

func game(v int) {
	var (
		limite     int
		str        string
		res        string
		historico  = []string{}
		words      = list()
		vida       = v
		max        = len(words) - 1
		palavra    = words[randomNumber(max)]
		wordHidden = esconder(palavra)
	)
	if existArray(historico, palavra) {
		palavra = words[randomNumber(max)]
	}

	for vida != 0 {
		fmt.Printf("vc tem %v de vida", vida)
		fmt.Printf("\nA palavra é %v", wordHidden)
		fmt.Printf("\nvc ja jogou as letras %v \n", historico)
		if len(jogadas) > 0 {
			fmt.Printf("vc ja jogou as palaras %v \n", jogadas)
		}
		fmt.Scanln(&res)

		if res == "dica" && limite < 2 {
			limite++
			index := randomNumber(len(palavra) - 1)

			if existArray(wordHidden, string(palavra[index])) {
				for existArray(wordHidden, string((palavra[index]))) {
					index = randomNumber(len(palavra) - 1)
				}
			}

			for _, v := range palavra {
				if string(v) == string(palavra[index]) {
					wordHidden[index] = string(palavra[index])
				}
			}

			if !existArray(historico, string(palavra[index])) {
				historico = append(historico, string(palavra[index]))
			}

		} else if len(res) > 1 {
			fmt.Println("PARE DE TENTAR BUGAR O CODIGO")
		} else {

			if existArray(historico, res) {

				fmt.Println("perde uma vida pra ficar esperto")
				vida--
			} else {

				historico = append(historico, res)
				if strings.Index(palavra, res) == -1 {
					vida--
				} else {

					for k, v := range palavra {

						if res == string(v) {
							wordHidden[k] = string(palavra[k])
						}

					}

					for _, v := range wordHidden {
						str += string(v)
					}
					if str == palavra {
						fmt.Printf("A palavra é %v \n", wordHidden)
						break
					}
					str = ""
				}
			}
		}
	}
	jogadas = append(jogadas, palavra)
	if vida == 0 {
		restart("perdeu", palavra)

	} else {
		restart("ganhou", palavra)

	}

}
func existArray(arr []string, str string) bool {
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}
func restart(x string, y string) {
	var res string
	fmt.Printf("vc %v !! a palavra era %v \n", x, y)
	fmt.Println("Digite M para votar pro menu")
	fmt.Println("Digite S para finalizar o game")
	fmt.Scanln(&res)
	switch res {
	case "M":
		main()
	case "S":
		final()
	default:
		fmt.Println("DIGITA CERTO")
		restart(x, y)
	}
}
func randomNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	min := 0
	return rand.Intn(max-min+1) + min
}
func esconder(s string) []string {
	var wordHidden = []string{}
	for i := 0; i < len(s); i++ {
		wordHidden = append(wordHidden, "-")
	}
	return wordHidden
}
func final() {

	fmt.Println("GG")
}
