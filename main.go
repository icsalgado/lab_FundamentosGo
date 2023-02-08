package main
import "fmt"

func main(){
	fmt.Println("Conversao de escalas termometricas")

	var H2OCelcius float64
	var H2OKelvin float64

	fmt.Println("Informe a temperatura em graus Celsius da agua")
	fmt.Scanln(&H2OCelcius)

	H2OKelvin = H2OCelcius + 273.15

	fmt.Println("A temperatura da agua em Celcius eh: ", H2OCelcius)
	fmt.Println("A temperatura da agua em Kelvin eh: ", H2OKelvin)
	fmt.Println("O ponto de ebulição da agua é 100 graus Celsius ou ", 100 + 273.15, " Kelvin")


}