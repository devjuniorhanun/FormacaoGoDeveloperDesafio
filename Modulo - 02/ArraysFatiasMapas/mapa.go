/* A primeira linha define o pacote principal (package main) do programa */

package main

/* Importa o pacote "fmt" que é utilizado para realizar formatação de saída e
leitura de entrada no terminal. */
import "fmt"

// O código a seguir, cria dois slice, e copia os elementos de um para o outro
func main() {
	// Cria uma variável chamada elemento e inicializada ela como um mapa (map) de strings para strings.
	elemento := make(map[string]string)
	// Cria Três elementosm que são adicionados ao mapa.
	elemento["H"] = "Hidrogênio" // A chave "H" tem o valor "Hidrogênio"
	elemento["He"] = "Hélio"     // A chave "He" tem o valor "Hélio"
	elemento["Li"] = "Lítio"     // A chave "Li" tem o valor "Lítio".
	// A função Println é utilizada para imprimir o valor associado à chave "Li" no mapa elemento. Neste caso, o valor é "Lítio".
	fmt.Println(elemento["Li"])
	// O resultado da execução desse código será a impressão da string "Lítio", que é o valor associado à chave "Li" no mapa elemento.
}

//coleção ordenada (lista) pares chave-valor
//var x map[string]int
//x é um mapa de strings para ints
