//newtorn方を用いて平方根を求める
//

package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {

	z := 1.0;
	for i:=0; i < 100; i++ {
		z = z - (z*z -x) / (2*x)

	}
	return z
}

func main() {
	fmt.Println(Sqrt(9))
}

