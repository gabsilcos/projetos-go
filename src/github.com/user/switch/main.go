package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	numero := 3
	fmt.Print("O número se escreve assim: ")
	switch numero {
	case 1:
		fmt.Println("um.")
	case 2:
		fmt.Println("dois.")
	case 3:
		fmt.Println("tres.")
	}

	fmt.Println("Vc é da família do Unix?")
	switch runtime.GOOS {
	case "windows":
		fmt.Println("Blerg")
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println("Sim, sou hipster!")
	default:
		fmt.Println("Sim sou rWindows")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Vai dormir!")
	default:
		fmt.Println("Provavelmente vc está atrasado :D")
	}

	numero = 9
	fmt.Println("Só tem um dígito?")
	switch {
	case numero < 10:
		fmt.Println("Sim")
	case numero >= 10 && numero < 100:
		fmt.Println("Não, serve dois dígitos?")
	case numero >= 100:
		fmt.Println("Vai a merda :D")

	}
}
