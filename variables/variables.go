package main

import "fmt"

func main() {

    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "apple"
    fmt.Println(f)

    var g, h, i = 3, a, 'z'
    fmt.Println(g, h, i)

    const j = 0.12321412141219
    fmt.Println(j)

    const k = (2+3)
    fmt.Println(k)
}