package piscine

func Swap(a *int, b *int) {
	aa := *b
	*b = *a
	*a = aa
}
