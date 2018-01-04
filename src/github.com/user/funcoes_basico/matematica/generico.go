package matematica

/*
Calculo é uma função que recebe três parâmetros - uma função e dois números - e
retorna um número, executando qqr tipo de cálculo.
*/
func Calculo(funcao func(int, int) int, a int, b int) int {
	return funcao(a,b)
}

//Multiplicador que... multiplica :D
func Multiplicador(x int, y int) int {
	return x*y
}