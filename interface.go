package figuras
import "fmt"

type Figura interface {
	Area() float64
	Perimetro() float64
}


func Medidas(f Figura){
	fmt.Println(f)
	fmt.Println("Area: ", f.Area())
	fmt.Println("Perimetro: ", f.Perimetro())
}

