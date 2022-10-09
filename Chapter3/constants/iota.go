package main

type weekday int

const (
	Sunday weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type flags uint

const (
	FlagUp flags = 1 << iota
	FlagB
	FlagC
)

func main() {
	println(FlagUp)
	println(FlagB)
	println(FlagC)
}