package whatColor

type HueColor string

type HueValue struct {
	Color       string `json:"color"`
	Hex         string `json:"hex"`
	Coefficient int    `json:"coefficient"`
}

type HueCollection []HueValue
