//cobtains my solutions to all exercise from golang.org tour
package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	//"github.com/smgolangdev/godev/goexer/exers"
	"github.com/smgolangdev/godev/goexers/exers"
	"strconv"
)

type Vertex struct {
	Lat, Long float64
}

func makeMap(key string, v Vertex) map[string]Vertex {
	m := make(map[string]Vertex)
	m[key] = v
	return m
}

func main() {
	usage := `Usage: goexer run <exernum> <exargs>

Run exernum the exercise number from golang tour.

Options:
  -h --help      Runs any one of the exercises from golang.org. Enter an exercise number.
  --version      Show current version.`

	arguments, _ := docopt.Parse(usage, nil, true, "goexer", false)

	fmt.Println("Running golang exercise number now: ", arguments["<exernum>"])
	fmt.Println(arguments["<exernum>"])
	fmt.Println(arguments["<exargs>"])

	switch arguments["<exernum>"] {
	case "38":
		fmt.Println("Runninge exercise number - 38 ", arguments["<exargs>"])
		st := arguments["<exargs>"]
		ft, _ := strconv.ParseFloat(st.(string), 64)
		fmt.Println("The float is:", ft)
		fmt.Println("Square root of 2 = %v/n", exers.SquareRoot(ft))

	case "43":
		fmt.Println(exers.WordCount("I have a have not cake cake fox"))
	case "46":
		//Run exerise 46
		ff := exers.Fibonacci()
		for j := 0; j <= 20; j++ {
			fmt.Println(ff())
		}
	case "54":
		fmt.Println("Runninge exercise number - 54")
	case "58":
		fmt.Println("Running exercise number - 58")
		ans, err := exers.MyRt(-7)
		fmt.Println(ans, err)
	case "64":
		m := makeMap("myKey", Vertex{40.6844, -74.39967})
		v := map[string]Vertex{
			"Google":    Vertex{89.0987, -76.666},
			"Microsoft": Vertex{89.0987, -76.666},
			"Oracle":    Vertex{89.0987, -76.666},
		}
		fmt.Println(m["myKey"])
		fmt.Println(v)
	default:
		fmt.Println("Incorrect exercise number.")
	}

}
