package main

func main() {

	limit := 4000000
	ef1 := 0
	ef2 := 2
	sum := ef1 + ef2
	ef3 := limit

	for ef2 <= limit {
		ef3 = 4*ef2 + ef1
		if ef3 > limit {
			break
		}
		ef1 = ef2
		ef2 = ef3
		sum = sum + ef2
	}
	println(sum)
}
