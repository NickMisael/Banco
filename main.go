package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func limpaTela() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
}

//Struct Conta
type Conta struct {
	NumConta  int
	FirstName string
	LastName  string
	Idade     int
	Email     string
	Saldo     float32
}

func Cadastrar() {
	c := Conta{}
	fmt.Printf("Digite o Primeiro Nome: ")
	fmt.Scanf("%s", &c.FirstName)
	fmt.Printf("Digite o Sobrenome: ")
	fmt.Scanf("%s", &c.LastName)
	fmt.Printf("Digite a Idade: ")
	fmt.Scanf("%d", &c.Idade)
	fmt.Printf("Digite o Email: ")
	fmt.Scanf("%s", &c.Idade)
	limpaTela()
	fmt.Println("Carregando Dados...")
	time.Sleep(time.Second + 3)
}

func Consultar() {

	fmt.Println("Tecle para sair...")
}

func Deletar() {
	limpaTela()
	var id int
	var err error
	for {
		var i string
		fmt.Printf("Digite o ID a ser removido: ")
		fmt.Scanf("%s", &i)
		if id, err = strconv.Atoi(i); err != nil {
			fmt.Println("Digite um número válido")
		} else if id < 1 {
			fmt.Println("Digite um número válido")
		} else {
			break
		}
	}
	limpaTela()
	time.Sleep(time.Second + 2)
}

func main() {
	for {
		var esc string
		limpaTela()
		fmt.Println("1 - Adicionar Cliente")
		fmt.Println("2 - Alterar Cliente")
		fmt.Println("3 - Excluir Cliente")
		fmt.Println("4 - Buscar Cliente")
		fmt.Println("5 - Listar Cliente")
		fmt.Println("6 - Sair do sistema Banco")
		fmt.Printf("-> ")
		fmt.Scanf("%s", &esc)

		if es, er := strconv.Atoi(esc); er != nil {
			fmt.Println("Engraçadão você hein!!")
			fmt.Println("Digite um número válido!!")
			time.Sleep(time.Second + 4)
		} else if es < 1 || es > 6 {
			fmt.Println("Engraçadão você hein!!")
			fmt.Println("Digite um número válido!!")
			time.Sleep(time.Second + 4)
		} else {
			if es == 6 {
				fmt.Println("Obrigado por utilizar!!")
				break
			} else {
				switch es {
				case 1:
					limpaTela()
					Cadastrar()
				case 2:
					limpaTela()
					Consultar()
					for {
						fmt.Scanf("%s", &esc)
						if esc != "" {
							break
						}
					}
					time.Sleep(time.Second + 2)
					limpaTela()
				case 3:
					Deletar()
				}
			}
		}
	}
}
