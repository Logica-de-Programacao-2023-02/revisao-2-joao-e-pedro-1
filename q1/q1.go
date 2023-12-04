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
