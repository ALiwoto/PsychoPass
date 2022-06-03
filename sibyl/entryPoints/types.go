/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package entryPoints

type EndpointResponse struct {
	Success bool           `json:"success"`
	Result  any            `json:"result"`
	Error   *EndpointError `json:"error"`
}

type EndpointError struct {
	ErrorCode int    `json:"code"`
	Message   string `json:"message"`
	Origin    string `json:"origin"`
	Date      string `json:"date"`
}
