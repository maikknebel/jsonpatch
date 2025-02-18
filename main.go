// The entrypoint for the package.
package main

func main() {
	print("Hello World")

	// size := 3
	src := make([]int, 3)
	// foo
	dst := make([]int, 3)

	copy(dst, src)
	println(src)
}
