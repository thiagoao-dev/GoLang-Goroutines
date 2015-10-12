package functions

import "fmt"
import "time"

func PrintMe(n int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", n)
		time.Sleep(200 * time.Millisecond)
	}
}
