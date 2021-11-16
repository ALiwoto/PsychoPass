package whatColor

func GetHueColor(hue int) HueColor {
	for key, value := range hueRangeMap {
		if key.IsInRange(hue) {
			return HueColor(value)
		}
	}

	return ""
}
