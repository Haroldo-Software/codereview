package main

func main() {
	slice := []int{1, 2, 3, 4, 5}
	 
	for i := 0; i < len(slice); i++ {
		print(slice[i], " ")
	}
	println()
	slice = slice[:len(slice)-1]
	for i := 0; i < len(slice); i++ {
		print(slice[i], " ")
	}
}
