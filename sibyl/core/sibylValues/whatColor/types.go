package whatColor

type HueColor string

type hueValue struct {
	Color       string
	Hex         string
	Coefficient int
}

type hueCollection []hueValue
