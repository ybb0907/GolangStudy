package main

func Pic(dx, dy int) [][]uint8 {
	if dx <= 0 || dy <= 0 {
		return nil
	}
	res := make([][]uint8, dy)

	for i := range res {
		res[i] = make([]uint8, dx)
	}
	return res
}

func main() {

}
