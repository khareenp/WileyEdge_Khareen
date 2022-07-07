package notes

import ("fmt"
"math"
)
//creates a class of function
//if i create a variable of type Operator, it will be a function
type Operator func(x float64) float64

// Map applies op to each element of a.
//takes 2 parameters one is Operator type other is float slice
//return slice of floats
func Map(op Operator, a []float64) []float64 {
    //make creates a slice (type, size)
    res := make([]float64, len(a))
    //for every element in slice a
    for i, x := range a {
        res[i] = op(x)
    }
    return res
}

func main() {
    op := math.Abs
    a := []float64{1, -2}
    b := Map(op, a)
    fmt.Println(b) // [1 2]

    c := Map(func(x float64) float64 { return 10 * x }, b)
    fmt.Println(c) // [10, 20]
}