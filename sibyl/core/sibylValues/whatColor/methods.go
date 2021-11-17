package whatColor

func (r *HueRange) IsInRange(hue int) bool {
	if r.IsUnlimited() {
		return true
	}

	if r.IsLeftUnlimited() && hue <= r.end {
		return true
	}

	if r.IsRightUnlimited() && hue >= r.start {
		return true
	}

	return hue >= r.start && hue <= r.end
}

func (r *HueRange) IsUnlimited() bool {
	return r.start == unlimited && r.end == unlimited
}

func (r *HueRange) IsLeftUnlimited() bool {
	return r.start == unlimited
}

func (r *HueRange) IsRightUnlimited() bool {
	return r.end == unlimited
}
