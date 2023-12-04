package q1

//Na loja de animais à venda, existem alguns tipos de ração disponíveis para compra, sendo eles:
//
//* Ração para cachorro
//* Ração para gato
//* Ração universal
//
//O estoque dessas rações está representado em um mapa, onde a chave é o nome da ração e o valor correspondente é a
//quantidade disponível em estoque.
//
//Polycarp possui 𝑥 cães e 𝑦 gatos. Gostaríamos de determinar se é possível para ele comprar comida suficiente para todos
//os seus animais na loja. Cada um dos seus cães e gatos deve receber um pacote de ração adequado para sua espécie.

func CanBuyFood(stock map[string]int, dogs, cats int) bool {
	return false
}
package main

import "fmt"

func main() {
    quantCachorros := 2
    quantGatos := 3

    estoque := map[string]int{
        "dog":       2,
        "cat":       2,
        "universal": 2,
    }

    if quantCachorros > estoque["dog"] {
        fmt.Println("Não há ração suficiente para os cachorros.")
        return
    }

    if quantGatos > estoque["cat"] {
        fmt.Println("Não há ração suficiente para os gatos.")
        return
    }

    estoque["dog"] -= quantCachorros
    estoque["cat"] -= quantGatos

    quantRestante := quantCachorros + quantGatos
    if quantRestante > estoque["universal"] {
        fmt.Println("Não há ração universal suficiente para o restante dos animais.")
        return
    }

    fmt.Println("Polycarp pode comprar comida suficiente para todos os seus animais.")
}
