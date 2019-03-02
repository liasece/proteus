package categories

// CategoryOptions is not marked for generations, but it is used in another
// structs, so it will be generated because of it.
type CategoryOptions struct {
	ShowPrices bool
	CanBuy     bool
	// The next field was deleted.
	// Field3 int

	// 4th field, use `proteus:",4"` forward compatible
	Field4 int `proteus:",4"`
}
