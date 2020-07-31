package main

import (
	"fmt"

	"gitlab.com/go-programming/007-documentacion/myMathBrunidas"
	// "github.com/GoesToEleven/go-programming/code_samples/007-documentation/01/myMathBrunidas"
)

func main() {
	fmt.Println("2 + 3 =", myMathBrunidas.Sum(2, 3))
	fmt.Println("4 + 7 =", myMathBrunidas.Sum(4, 7))
	fmt.Println("5 + 9 =", myMathBrunidas.Sum(5, 9))
}
