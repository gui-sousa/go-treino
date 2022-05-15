/*
- Crie um tipo. O tipo subjacente deve ser int.
- Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
- Na função main:
    1. Demonstre o valor da variável "x"
    2. Demonstre o tipo da variável "x"
    3. Atribua 42 à variável "x" utilizando o operador "="
    4. Demonstre o valor da variável "x"
*/

package main

import (
	"fmt"
)

type quaqua int

var x quaqua

func main() {
	fmt.Println("Valor da variavel X: ", x)
	fmt.Printf("Tipo da variavel X: %T\n", x)

	x = 42

	fmt.Println("Novo valor de X: ", x)
}
