package whatColor

/*

| Color  | Name  | Hex Code | Coefficient |
| ------------ | ------------ | ------------ | ------------ |
| ![DarkRed]  | DarkRed | 8B0000 | 600+ |
| ![DarkBlue]  | DarkBlue | 00008B | 551-600 |
| ![Indigo] | Indigo | 4B0082 | 501-550 |
| ![DarkGreen] | DarkGreen | 006400 | 451-500 |
| ![Green] | Green | 008000 | 401-450 |
| ![Red] | Red | FF0000 | 351-400 |
| ![Purple] | Purple | 800080 | 301-350 |
| ![Navy] | Navy | 000080 | 251-300 |
| ![Blue] | Blue | 0000FF | 201-250 |
| ![Magenta] | Magenta | FF00FF | 151-200 |
| ![FireBrick] | FireBrick | B22222 | 101-150 |
| Restored user range  | ------------ | ------------ | ------------ |
| ![Violet] | Violet | EE82EE | 96-100 |
| ![Crimson] | Crimson | DC143C | 91-95 |
| ![Salmon] | Salmon | FA8072 | 86-90 |
| ![SlateBlue] | SlateBlue | 6A5ACD | 81-85 |
| Below 80 Range begins  | ------------ | ------------ | ------------ |
| ![Lime] | Lime | 00FF00 | 75-80 |
| ![Teal] | Teal | 008080 | 71-75 |
| ![LightPink]  | LightPink | FFB6C1 | 66-70 |
| ![Cyan] | Cyan | 00FFFF | 61-65 |
| ![Lavender] | Lavender | E6E6FA | 56-60 |
| ![LightBlue] | LightBlue | ADD8E6 | 51-55 |
| ![SkyBlue] | SkyBlue | 87CEEB | 46-50 |
| ![Aquamarine] | Aquamarine | 7FFFD4 | 41-45 |
| ![HotPink] | HotPink | FF69B4 | 36-40 |
| ![Aqua] | Aqua | 00FFFF | 31-35 |
| ![Pink] | Pink | FFC0CB | 25-30 |
| ------------ | ------------ | ------------ | ------------ |
*/

var (
	DarkRedRange    = &HueRange{600, unlimited}
	DarkBlueRange   = &HueRange{551, 600}
	IndigoRange     = &HueRange{501, 550}
	DarkGreenRange  = &HueRange{451, 500}
	GreenRange      = &HueRange{401, 450}
	RedRange        = &HueRange{351, 400}
	PurpleRange     = &HueRange{301, 350}
	NavyRange       = &HueRange{251, 300}
	BlueRange       = &HueRange{201, 250}
	MagentaRange    = &HueRange{151, 200}
	FireBrickRange  = &HueRange{101, 150}
	VioletRange     = &HueRange{96, 100}
	CrimsonRange    = &HueRange{91, 95}
	SalmonRange     = &HueRange{86, 90}
	SlateBlueRange  = &HueRange{81, 85}
	LimeRange       = &HueRange{75, 80}
	TealRange       = &HueRange{71, 75}
	LightPinkRange  = &HueRange{66, 70}
	CyanRange       = &HueRange{61, 65}
	LavenderRange   = &HueRange{56, 60}
	LightBlueRange  = &HueRange{51, 55}
	SkyBlueRange    = &HueRange{46, 50}
	AquamarineRange = &HueRange{41, 45}
	HotPinkRange    = &HueRange{36, 40}
	AquaRange       = &HueRange{31, 35}
	PinkRange       = &HueRange{25, 30}
)

// read-only map for converting ranges to string, do NOT edit.
var (
	hueRangeMap = map[*HueRange]HueColor{
		DarkRedRange:    "DarkRed",
		DarkBlueRange:   "DarkBlue",
		IndigoRange:     "Indigo",
		DarkGreenRange:  "DarkGreen",
		GreenRange:      "Green",
		RedRange:        "Red",
		PurpleRange:     "Purple",
		NavyRange:       "Navy",
		BlueRange:       "Blue",
		MagentaRange:    "Magenta",
		FireBrickRange:  "FireBrick",
		VioletRange:     "Violet",
		CrimsonRange:    "Crimson",
		SalmonRange:     "Salmon",
		SlateBlueRange:  "SlateBlue",
		LimeRange:       "Lime",
		TealRange:       "Teal",
		LightPinkRange:  "LightPink",
		CyanRange:       "Cyan",
		LavenderRange:   "Lavender",
		LightBlueRange:  "LightBlue",
		SkyBlueRange:    "SkyBlue",
		AquamarineRange: "Aquamarine",
		HotPinkRange:    "HotPink",
		AquaRange:       "Aqua",
		PinkRange:       "Pink",
	}
)
