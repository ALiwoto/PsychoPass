/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package whatColor

var (
	// read-only map for converting coefficients to string, do NOT edit.
	// do NOT use set operations on this map during runtime;
	// as it's not using any mutex for protecting itself.
	// it should be loaded from json data received from github or
	// loaded from local file (see `endPoint` and `fileName` constants).
	hueColorMap    = map[int]HueColor{}
	maxCoefficient = 0 /* will be set dynamically */
)
