package piscine

func UltimateDivMod(a *int, b *int) {

	alo := *a / *b
	nn := *a % *b
	*a = alo
	*b = nn
}
