package main

import "fmt"

func xpto(a *int) int {
	*a = 100
	return *a
}

func main() {

	x := 10         // X está fazendo um apontamento na memória
	fmt.Println(&x) // 0x1400001a0f8 -> O valor impresso é o endereçamento de memória do valor X

	y := &x        //  Go apomnta o Y para o mesmo endereço de memória onde está o X (referência)
	fmt.Println(y) // 0x1400001a0f8 -> AO invés de imprimir o X irá imprimri o endereço de X

	fmt.Println(*y) // 10 -> Usar reference para poder printar o valor de Y

	*y = 20        // Mudança se aplica ao local de memória, logo X terá mesmo valor
	fmt.Println(x) //Eu quero pegar o ponteiro de Y e mudar para 20;

	var z *int = &x
	fmt.Println(*z)

	b := 10
	fmt.Println(xpto(&b))
	fmt.Println(b)
}
