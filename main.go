package main

import (
	"fmt"
	"os"
	"time"
	"github.com/fatih/color"
)

func main(){
	black := color.New(color.FgBlack).Add(color.Bold)
	black.Println("Seja bem ao Study Music. Aplicativo feito para o seu desenvolvimento pessoal, musical. Um simples aplicativo de erro e acerto.")	
	introducao()	
	// Scale()	
}


func introducao(){
	black := color.New(color.FgBlack).Add(color.Bold) // 
	black.Println("1- Escalas") //Criar essa função primeiro
	black.Println("2- Campo Harmônico")
	black.Println("3- Escala Menores")
	black.Println("4- Escalas Maiores")
	black.Println("5- Escala Pentatônica")
	black.Println("0- Pressione 0 para Sair")

	var escolhaUser int
	black.Println("Qual você deseja acessar?")
	fmt.Scanf("%d", &escolhaUser)

	switch escolhaUser{
	case 1:
		fmt.Println("1- Inicializando Quiz sobre Escalas...")
		Scale()
		time.Sleep(2* time.Second)
	case 2:
		fmt.Println("2- Inicializando Quiz Campo Harmônico...") // Substituir isso pela função ( A criar )
		time.Sleep(2* time.Second)
	case 3:
		fmt.Println("3- Inicializando Quiz sobre Escalas Menores...") // Substituir isso pela função ( A criar )
		time.Sleep(2* time.Second)
	case 4:
		fmt.Println("4- Inicializando Quiz sobre Escalas Maiores...") // Substituir isso pela função ( A criar )
		time.Sleep(2* time.Second)
		case 5:
		fmt.Println("5- Inicializando Quiz sobre Escala Pentatônica...") // Substituir isso pela função ( A criar )
		time.Sleep(2* time.Second)
		case 0:
			fmt.Println("0 - Saindo do Aplicação...")
			time.Sleep(2* time.Second)
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido ")	
			os.Exit(-1)
	}
}


func Scale(){

	fmt.Println("Sabemos que Escala musical nada mais é do que uma sequência de notas de nomes diferentes organizadas dentro de uma mesma oitava, separadas por intervalos musicais.")
	fmt.Println("Nesse modo, trabalharemos com as escalas maiores e menores, de forma sortida.")
	fmt.Println("Sabendo que existem 8 notas naturais, sendo elas : C - D - E - F - G - A - B - C")

	fmt.Println("Por qual nota você gostaria de começar?")

	var user_escolhe_nota string
	fmt.Scan(&user_escolhe_nota)

	switch user_escolhe_nota{
	case "C":
		C()// Terminar de fazer a função ( A criar)
	case "D":
		fmt.Println("Começaremos com Ré...") // Substituir isso pela função ( A criar)
	case "E":
		fmt.Println("Começaremos com Mi...") // Substituir isso pela função ( A criar)
	case "F":
		fmt.Println("Começaremos com Fá...") // Substituir isso pela função ( A criar)
	case "G":
		fmt.Println("Começaremos com Sol...") // Substituir isso pela função ( A criar)
	case "A":
		fmt.Println("Começaremos com Lá...") // Substituir isso pela função ( A criar)
	case "B":
		fmt.Println("Começaremos com Si...") // Substituir isso pela função ( A criar)
	default:
		fmt.Println("Ops... Não reconheci essa nota... Tente novamente")
		os.Exit(0)
	}	
}

func C() {
	
	fmt.Println("Qual é o quinto grau de C?")

	var chooseC string
	fmt.Scan(&chooseC)

	if chooseC == "G" {
		fmt.Println("Você acertou!!")
		time.Sleep(1* time.Second)		
	} else  {
		fmt.Println("Oh, parece que essa resposta não está certa :(")
		os.Exit(-0)
	}

}
