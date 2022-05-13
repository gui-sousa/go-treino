/*
- Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
    1. 42
    2. "James Bond"
    3. true
- Agora demonstre os valores nestas variáveis, com:
    1. Uma única declaração print
    2. Múltiplas declarações print
*/
package main

func main() {

	var1 := 42
	var2 := "James Bond"
	var3 := true

	/*1. Uma única declaração print*/
	println("Valor das variáveis com um Print \n\nVAR1 =", var1, "\nVAR2 =", var2, "\nVAR3 =", var3, "\n")

	/*2. Múltiplas declarações print*/
	println("Valor com vários Prints:\n")
	println("Variável 1 =", var1)
	println("Variável 2 =", var2)
	println("Variável 3 =", var3)
}
