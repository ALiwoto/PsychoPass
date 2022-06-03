/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package whatColor

type HueColor string

type HueValue struct {
	Color       string `json:"color"`
	Hex         string `json:"hex"`
	Coefficient int    `json:"coefficient"`
}

type HueCollection []HueValue
