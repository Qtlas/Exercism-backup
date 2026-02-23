package luhn

import (
	"fmt"
	"strings"
    "strconv"
)

func Valid(id string) bool {
	n_id := strings.ReplaceAll(id, " ", "")
    if len(n_id) <= 1 {
        return false
    } 
	somme, pas := 0, 0

	for i := len(n_id)-1; i >= 0; i-- {
		val := n_id[i]
		n, err := strconv.Atoi(string(val))

		if err != nil {
			fmt.Println("Error : ", err)
			return false
		}

		if  pas % 2 == 1{	
			n *= 2
			if n >= 10 {
				n -= 9
			}
		}

		somme += n
		pas++
	}
	return somme % 10 == 0
}