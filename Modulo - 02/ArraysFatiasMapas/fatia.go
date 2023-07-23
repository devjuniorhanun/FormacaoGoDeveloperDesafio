/* A primeira linha define o pacote principal (package main) do programa */

package main

/* Importa o pacote "fmt" que é utilizado para realizar formatação de saída e
leitura de entrada no terminal. */
import "fmt"

// O código a seguir, cria dois slice
func main() {

	// Cria uma fatia (slice) de inteiros chamada fatia1 é criada e inicializada com os valores 1, 2 e 3.
	fatia1 := []int{1, 2, 3}
	// Outra fatia de inteiros chamada fatia2 é criada usando a função make. Essa fatia tem tamanho inicial 1 e, portanto, pode armazenar apenas um elemento.
	fatia2 := make([]int, 1)
	// A função copy é usada para copiar elementos da fatia1 para a fatia2. No entanto, como a fatia2 tem capacidade para apenas um elemento, apenas o primeiro elemento de fatia1 (que é o valor 1) é copiado para fatia2.
	copy(fatia2, fatia1)
	// Por fim, a função Println do pacote fmt é utilizada para imprimir o conteúdo das duas fatias fatia1 e fatia2 no terminal.
	fmt.Println(fatia1, fatia2)
	// O resultado da execução desse código será a impressão das duas fatias:
	// [1 2 3] [1]

}

/* Observe que o elemento 3 não foi copiado para fatia2 devido à sua capacidade limitada.
Se a capacidade de fatia2 fosse maior ou igual a 3, todos os elementos de fatia1 seriam
copiados para fatia2. */
