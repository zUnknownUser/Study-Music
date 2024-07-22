package main

import (
	"fmt"
	"os"
	"time"
)

func main(){
	fmt.Println("Seja bem ao Study Music. Aplicativo feito para o seu desenvolvimento pessoal, musical. Um simples aplicativo de erro e acerto.")	
	introducao()
	
}

// func VerificandoNocao(){
// 	var perguntaUser string
// 	fmt.Print("Você teve contato com algum tipo de aplicação como esse?")
// 	fmt.Scan(&perguntaUser)

// 	if perguntaUser == "Sim" {
//         fmt.Println("Ótimo... Estamos validando sua informação. E seguiremos adiante. Espero que possa aproveitar e aprimorar seus conhecimentos")
// 	} else  {
// 		fmt.Println("Certo, esperamos que sua primeira experiência seja boa! Aproveite.")
// 	}
	
// 	introducao()
// }

func introducao(){
	fmt.Println("1- Escalas")
	fmt.Println("2- Campo Harmônico")
	fmt.Println("3- Escala Menores")
	fmt.Println("4- Escalas Maiores")
	fmt.Println("5- Escala Pentatônica")
	fmt.Println("0- Pressione 0 para Sair")

	var escolhaUser int
	fmt.Println("Qual você deseja acessar?")
	fmt.Scanf("%d", &escolhaUser)

	switch escolhaUser{
	case 1:
		fmt.Print("1- Inicializando Quiz sobre Escalas...")
	case 2:
		fmt.Print("2- Inicializando Quiz Campo Harmônico...")
	case 3:
		fmt.Print("3- Inicializando Quiz sobre Escalas Menores...")
	case 4:
		fmt.Print("4- Inicializando Quiz sobre Escalas Maiores...")
		case 5:
		fmt.Print("5- Inicializando Quiz sobre Escala Pentatônica...")
		case 0:
			fmt.Println("0 - Saindo do Aplicação...")
			time.Sleep(2* time.Second)
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido ")	
			os.Exit(-1)
	}

}