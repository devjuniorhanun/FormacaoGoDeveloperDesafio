/* A primeira linha define o pacote principal (package main) do programa */

package main

/* Importa o pacote "fmt" que é utilizado para realizar formatação de saída e
leitura de entrada no terminal. */
import "fmt"

// O código a seguir, apresenta um programa simples que calcula a média
// dos elementos de um array de números de ponto flutuante.
func main() {
	// Declaração de um array de números de ponto flutuante com tamanho 5
	var x [5]float64

	// Inicialização dos elementos do array com valores específicos
	x[0] = 5.3
	x[1] = 8
	x[2] = 4.2
	x[3] = 2.1
	x[4] = 7.8

	// Declaração da variável "total" que armazenará a soma dos elementos do array
	var total float64 = 0

	// Loop para percorrer todos os elementos do array e somá-los à variável "total"
	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	// Cálculo da média dos elementos do array e exibição do resultado
	fmt.Println(total / float64(len(x)))
}

/*
O programa começa declarando e inicializando um array chamado x com tamanho 5.
Em seguida, inicializa uma variável chamada total com o valor zero.
Essa variável será usada para armazenar a soma de todos os elementos do array.

Depois disso, cria um loop for para percorrer todos os elementos do array x,
e cada valor é adicionado à variável total.

Por fim calcula o total da soma dividido pela quantidade de itens do array,

Em resumo, esse código calcula a média dos valores contidos no array x e imprime o resultado na tela. */
