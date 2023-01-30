package main

import (
	"fmt"
	"github.com/altair2503/myModules/mymath"
	"github.com/altair2503/myModules/myprinting"
)

func main() {
	myprinting.Hello()

	fmt.Println("The factorial of 3 is", mymath.Factorial(3))
	fmt.Println("Absolute value of -2 is", mymath.Abs(-2))
	fmt.Println("Absolute value of 3 is", mymath.Abs(3))

	myprinting.Bye()
}
