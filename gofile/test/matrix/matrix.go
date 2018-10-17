package matrix

type Matrix []int

func (a Matrix) Add(b Matrix) (c Matrix) {
	for idx, valx := range a {
		val := valx + b[idx]
		c = append(c, val)
	}
	return
}

// futures used internally
type futureMatrix chan Matrix

// API remains the same
func Inverse(a Matrix) Matrix {
	return <-InverseAsync(promise(a))
}

func Product(a Matrix, b Matrix) Matrix {
	return <-ProductAsync(promise(a), promise(b))
}

// expose async version of the API
func InverseAsync(a futureMatrix) futureMatrix {
	c := make(futureMatrix)
	go func() { c <- inverse(<-a) }()
	return c
}

func ProductAsync(a, b futureMatrix) futureMatrix {
	c := make(futureMatrix)
	go func() { c <- product(<-a, <-b) }()
	return c
}

// actual implementation is the same as before
func product(a Matrix, b Matrix) Matrix {
	//...
	return a.Add(b)
}

func inverse(a Matrix) Matrix {
	//....
	return a
}

// utility fxn: create a futureMatrix from a given matrix
func promise(a Matrix) futureMatrix {
	future := make(futureMatrix, 1)
	future <- a
	return future
}
