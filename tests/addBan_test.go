/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package tests_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"testing"
)

//---------------------------------------------------------

// TestAddBan01 is a test function to ban a new user in the
// Sibyl System. If Sibyl System returns that the user is already banned
// we should ignore the error.
func TestAddBan01(t *testing.T) {
	decideToRun()
	t.Cleanup(closeServer)
	ownerToken := getOwnerToken()
	doTestAddBan01(t, ownerToken, "addBan", http.MethodPost, 1234)
	doTestAddBan01(t, ownerToken, "AddBan", http.MethodPost, 12345)
	doTestAddBan01(t, ownerToken, "BanUser", http.MethodGet, 123456)
	doTestAddBan01(t, ownerToken, "ban", http.MethodGet, 1234567)
}

func doTestAddBan01(t *testing.T, ownerToken, path, method string, id int64) {
	// create a new token
	req, err := http.NewRequest(http.MethodPost, getBaseUrl()+path, nil)
	if err != nil {
		t.Fatal(err)
		return
	}

	req.Header.Set("token", ownerToken)
	req.Header.Set("user-id", strconv.FormatInt(id, 10))
	req.Header.Set("reason", "TestAddBan01")
	req.Header.Set("message", "t.meow/google/123456")
	req.Header.Set("srcUrl", "t.meow/base_chat/123")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusAccepted {
			// already banned with the same parameters.
			return
		}
		strValue, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
			return
		}
		t.Fatal("got unexpected status code: ", res.StatusCode, string(strValue))
		return
	}

	strValue, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
		return
	}

	valueMap := make(map[string]interface{})
	err = json.Unmarshal(strValue, &valueMap)
	if err != nil {
		t.Fatal(err)
		return
	}

	log.Println(string(strValue))
}

//---------------------------------------------------------
