package piscine

func UltimateDivMod(a *int, b *int) {

	devided := *a / *b
	nn := *a % *b
	*a = devided
	*b = nn
}
