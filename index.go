/* package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello there %s", "dude!")
	})

	srv := http.Server{
		Addr: ":8080",

	}

	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}
} */

//! Go Tour

/* package main

import (
	"fmt"
) */

/* func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Ahora tenemos %g problemas.\n", math.Sqrt(7))
	var piNumber float64 = math.Pi
	fmt.Println(piNumber)
} */

/* func add(number1 int, number2 int) int {
	return number1 + number2
}

func main() {
	fmt.Println(add(30000, 60000))
} */

/* func add(x, y int) int {
	return x * y
}

func main() {
	fmt.Println(add(42, 13))
} */

/* func swap(first_word, second_word string) (string, string) {
	return first_word, second_word
}

func main() {

	a, b := swap("hello", "world")
	fmt.Println(a, b)
} */

/* func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
} */

package main

import (
	"net/http"
)

func main() {

	//routes

	http.HandleFunc("/", homeHandler)

	//inicio del server
	http.ListenAndServe(":5000", nil)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello Golang"))

}
