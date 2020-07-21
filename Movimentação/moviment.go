package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Função que limpa a tela
func limpaTela() {
	if runtime.GOOS == "windows" {
		clear := exec.Command("cmd", "/c", "cls")
		clear.Stdout = os.Stdout
		clear.Run()
	} else if runtime.GOOS == "linux" {
		clear := exec.Command("clear")
		clear.Stdout = os.Stdout
		clear.Run()
	}
}

func TelaLogin() {
	//fmt.Printf("\033[47m")
	//fmt.Printf("\033[41m")
	fmt.Print("\a \v")
	fmt.Print("\t __________________________________________\n")
	fmt.Print("\t|                                          |\n")
	fmt.Print("\t|    -----  --   __   ___ __  *            |\n")
	fmt.Print("\t|      |   |--  | _|  |  |  | |  |\\  |     |\n")
	fmt.Print("\t|      |    --  | \\   |  |  | |  |  \\|     |\n")
	fmt.Print("\t|                                          |\n")
	fmt.Print("\t|__________________________________________|\n")
	fmt.Print("\t| [$número da conta$]>> ")
}

func Menu(nome string) {
	//fmt.Printf("\033[47m")
	//fmt.Printf("\033[41m")
	fmt.Print("\a \v")
	fmt.Print("\t __________________________________________\n")
	fmt.Print("\t|                                          |\n")
	fmt.Print("\t|    -----  --   __   ___ __  *            |\n")
	fmt.Print("\t|      |   |--  | _|  |  |  | |  |\\  |     |\n")
	fmt.Print("\t|      |    --  | \\   |  |  | |  |  \\|     |\n")
	fmt.Print("\t|                                          |\n")
	fmt.Print("\t|__________________________________________|\n")
	fmt.Print("\t                                       \n")
	fmt.Printf("\t  Seja Bem-Vindo %s ao terminal! \n", nome)
	fmt.Print("\t                                       \n")
	fmt.Print("\t|==========================================|\n")
	fmt.Print("\t|                                          |\n")
	fmt.Print("\t|   1 - Saldo       2 - Saque              |\n")
	fmt.Print("\t|                                          |\n")
	fmt.Print("\t|   3 - Depósito    4 - Pagamento          |\n")
	fmt.Print("\t|                                          |\n")
	fmt.Print("\t|   5 - Sair                               |\n")
	fmt.Print("\t|                                          |\n")
	fmt.Print("\t|__________________________________________|\n")
	fmt.Print("\t|[++]-> ")
}

func LeArquivo() ([]string, error) {
	file, err := os.Open("db.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

type Cliente struct {
	Nconta   string
	Nome     string
	Saldo    float32
	Email    string
	Telefone string
	Pass     string
}

func (c *Cliente) Verifica(Nconta string, Pass string) (er error) {
	arquivo, err := LeArquivo()
	if err != nil {
		panic(err)
	}
	for _, item := range arquivo {
		campo := strings.Split(item, ",")
		if Nconta != campo[0] || Pass != campo[5] {
			continue
		} else {
			c.Nconta = campo[0]
			c.Nome = campo[1]
			time.Sleep(time.Second + 2)
			if saldo, err := strconv.ParseFloat(campo[2], 32); err == nil {
				sal := fmt.Sprintf("%.2f", saldo)
				//fmt.Println(sal)
				time.Sleep(time.Second + 2)
				if saldo, err = strconv.ParseFloat(sal, 32); err == nil {
					c.Saldo = float32(saldo)
				}
			}
			c.Email = campo[3]
			c.Telefone = campo[4]
			c.Pass = campo[5]
		}
	}
	if c.Email == "" {
		return errors.New("Não achamos sua conta :(")
	}
	return nil
}

func (c Cliente) VerSaldo() {
	fmt.Printf("\a \v")
	fmt.Println("\t ________________________________")
	fmt.Println("\t|                                |")
	fmt.Println("\t|             SALDO              |")
	fmt.Println("\t|                                |")
	fmt.Println("\t|================================|")
	fmt.Println("\t| \a")
	fmt.Println("\t|  Nome ->", c.Nome)
	fmt.Println("\t|  Email ->", c.Email)
	//fmt.Println("\t|  Saldo ->", c.Saldo)
	fmt.Printf("\t|  Saldo -> %.2f\n", c.Saldo)
	fmt.Println("\t| \a")
	fmt.Println("\t|________________________________")
}

func main() {
	// Conferir se já existe no banco
	_, err := LeArquivo()
	if err != nil {
		panic(err)
	}
	//time.Sleep(time.Second + 3)

	// Menu infinito
	for {
		/**********************************/
		/* 		AUTENTICANDO USUÁRIO	  */
		/**********************************/
		var esc string
		c := Cliente{}
		limpaTela()
		TelaLogin()
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Print("\a")
			c.Nconta = scanner.Text()
			break
		}
		if scanner.Err() != nil {
			panic(scanner.Err())
		}
		fmt.Print("\t| [$password$]>> ")
		scanner = bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Print("\a")
			c.Pass = scanner.Text()
			break
		}
		if scanner.Err() != nil {
			panic(scanner.Err())
		}
		c.Verifica(c.Nconta, c.Pass)
		if err != nil {
			panic(err)
		}
		if c.Email == "" {
			fmt.Println("\tErro: Você digitou algo errado!!")
			time.Sleep(time.Second + 2)

			/**********************************/
			/*		Abrindo uma sessão	      */
			/**********************************/
		} else {
			for {
				limpaTela()
				Menu(c.Nome)
				scanner := bufio.NewScanner(os.Stdin)
				for scanner.Scan() {
					fmt.Printf("\a")
					esc = scanner.Text()
					break
				}
				if scanner.Err() != nil {
					panic(scanner.Err())
				}
				if es, err := strconv.Atoi(esc); err != nil || es < 1 || es > 5 {
					fmt.Println("Erro: Número Inválido!")
					time.Sleep(time.Second + 3)
					limpaTela()
				} else if es == 5 {
					fmt.Println("Obrigado por utilizar o terminal :D")
					fmt.Println("Volte Sempre!!")
					break
				} else {
					/****************************/
					/* 		TELA DE ESCOLHA 	*/
					/****************************/
					limpaTela()
					switch es {
					case 1:
						c.VerSaldo()
						fmt.Println()
						fmt.Print("\tTecle para continuar...")
						fmt.Scanln(&esc)
					case 2:
						fmt.Println("Saque")
					case 3:
						fmt.Println("Depósito")
					case 4:
						fmt.Println("Pagamento")
					}
					time.Sleep(time.Second + 3)
				}
			}
			break
		}
	}
}
