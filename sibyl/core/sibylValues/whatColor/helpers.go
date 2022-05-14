/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package whatColor

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/logging"
)

func GetHueColor(hue int) HueColor {
	if len(hueColorMap) == 0 {
		return ""
	}

	if hue > maxCoefficient {
		return hueColorMap[maxCoefficient]
	}

	return hueColorMap[hue]
}

func loadColorsFromFile() {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		logging.Warn(err)
		return
	}

	collection := make(HueCollection, 0)
	err = json.Unmarshal(b, &collection)
	if err != nil {
		logging.Warn(err)
		return
	}

	loadValues(collection)
}

func LoadColors() {
	resp, err := http.Get(endPoint)
	if err != nil || resp == nil {
		loadColorsFromFile()
		return
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		loadColorsFromFile()
		return
	}

	collection := make(HueCollection, 0)
	err = json.Unmarshal(b, &collection)
	if err != nil {
		loadColorsFromFile()
		return
	}

	saveCollectionToFile(b)
	loadValues(collection)
}

func saveCollectionToFile(data []byte) {
	err := ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		logging.Warn(err)
	}
}

func loadValues(collection HueCollection) {
	for _, current := range collection {
		if current.Coefficient > maxCoefficient {
			maxCoefficient = current.Coefficient
		}

		hueColorMap[current.Coefficient] = HueColor(current.Color)
	}
}
