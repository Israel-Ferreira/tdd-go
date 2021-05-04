package arrays

func Soma(numeros []int) (soma int) {

	for _, num := range numeros {
		soma += num
	}

	return 
}


func GenerateSlice(limit int) []int {

	slice := []int{}

	for i := 1; i <= limit; i++ {
		slice =  append(slice, i)
	}

	return slice
}