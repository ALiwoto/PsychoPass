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
	"time"

	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
)

//---------------------------------------------------------

// TestCreateToken01 test to create a token using the owner token with http
// headers. The token should be created successfully.
func TestCreateToken01(t *testing.T) {
	decideToRun()
	t.Cleanup(closeServer)
	ownerToken := getOwnerToken()
	doTestCreateToken01(t, ownerToken, "create", http.MethodPost)
	doTestCreateToken01(t, ownerToken, "createToken", http.MethodPost)
	doTestCreateToken01(t, ownerToken, "create", http.MethodGet)
	doTestCreateToken01(t, ownerToken, "createToken", http.MethodGet)
}

func doTestCreateToken01(t *testing.T, ownerToken, path, method string) {
	// create a new token
	req, err := http.NewRequest(http.MethodPost, getBaseUrl()+path, nil)
	if err != nil {
		t.Fatal(err)
		return
	}

	req.Header.Set("token", ownerToken)
	req.Header.Set("user-id", userId01)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
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

// TestCreateToken02 is a test function to create a token using the owner token
/// with http url encoded variables. The token should be created successfully.
func TestCreateToken02(t *testing.T) {
	decideToRun()
	t.Cleanup(closeServer)

	ownerToken := getOwnerToken()
	doTestCreateToken02(t, ownerToken, "create", http.MethodPost)
	doTestCreateToken02(t, ownerToken, "createToken", http.MethodPost)
	doTestCreateToken02(t, ownerToken, "create", http.MethodGet)
	doTestCreateToken02(t, ownerToken, "createToken", http.MethodGet)
}

func TestCreateToken02Crazy(t *testing.T) {
	decideToRun()
	t.Cleanup(closeServer)
	t.Log("hello")
	ownerToken := getOwnerToken()
	doTestCreateToken02(t, ownerToken, "create", http.MethodPost)
	doTestCreateToken02(t, ownerToken, "createToken", http.MethodPost)

	start := time.Now()
	go func() {
		for i := 0; i < 10; i++ {
			go doTestCreateToken02CrazyPart(t, ownerToken)
			//log.Println(i)
		}
	}()
	for mm < 10 {
		time.Sleep(time.Second)
	}

	log.Println("done in:", time.Since(start).Microseconds())
	t.Log("done in:", time.Since(start).Microseconds())
}

var mm int

func doTestCreateToken02CrazyPart(t *testing.T, ownerToken string) {
	for i := 0; i < 5; i++ {
		doTestCreateToken02b(t, ownerToken, "create", http.MethodGet, false)
		doTestCreateToken02b(t, ownerToken, "createToken", http.MethodGet, false)
		//log.Println(i)
	}
	mm++
}

func doTestCreateToken02(t *testing.T, ownerToken, path, method string) {
	doTestCreateToken02b(t, ownerToken, path, method, true)
}

func doTestCreateToken02b(t *testing.T, ownerToken, path, method string, l bool) {
	url := getBaseUrl() + path + "?user-id=" + userId01 + "&token=" + ownerToken
	// create a new token
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		t.Fatal(err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		strValue, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
			return
		}
		t.Fatal("got unexpected status code: ", res.StatusCode,
			"with content", string(strValue))
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

	/* Expected result should be something like this:
		4
	2021/10/25 20:19:40 {"result":{"id":1341091260,"hash":"1341091260:MiDvOMnpj8Tf7HK6OLVqVMIJg7F4on9Tyr6mRFhtVpesncgMidjc8BbN6etulbfq","permission":0,"created_at":"2021-10-25T19:54:51.194525003Z"},"success":true,"error":null}
	*/

	if l {
		log.Println(string(strValue))
	}
}

//---------------------------------------------------------

func TestCreateToken03(t *testing.T) {
	decideToRun()
	t.Cleanup(closeServer)

	ownerToken := getOwnerToken()
	doTestCreateToken03(t, ownerToken, "create", http.MethodPost)
	doTestCreateToken03Wrong(t, user01TokenTmp, "create", http.MethodPost)
	doTestCreateToken03Wrong(t, user01TokenTmp, "create", http.MethodGet)
	doTestCreateToken03Wrong(t, user01TokenTmp, "createToken", http.MethodPost)
}

func doTestCreateToken03(t *testing.T, ownerToken, path, method string) {
	url := getBaseUrl() + path + "?user-id=" + userId01 + "&token=" + ownerToken
	// create a new token
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		t.Fatal(err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatal("got unexpected status code: ", res.StatusCode)
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

	if len(valueMap) < 1 {
		t.Fatal("got unexpected result: ", valueMap)
		return
	}

	token, _ := valueMap["result"].(map[string]interface{})
	if token == nil {
		t.Fatal("token cannot be nil")
	}
	user01TokenTmp, _ = token["hash"].(string)
	/* Expected result should be something like this:
		4
	2021/10/25 20:19:40 {"result":{"id":1341091260,"hash":"1341091260:MiDvOMnpj8Tf7HK6OLVqVMIJg7F4on9Tyr6mRFhtVpesncgMidjc8BbN6etulbfq","permission":0,"created_at":"2021-10-25T19:54:51.194525003Z"},"success":true,"error":null}
	*/

	log.Println(string(strValue))
}

func doTestCreateToken03Wrong(t *testing.T, ownerToken, path, method string) {
	url := getBaseUrl() + path + "?user-id=" + userId01 + "&token=" + ownerToken
	// create a new token
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		t.Fatal(err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		t.Fatal("got unexpected status code: ", res.StatusCode)
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

	/* Expected result should be something like this:
	2021/10/25 20:48:37 {"result":null,"success":false,"error":{"code":502,"message":"Permission Denied","origin":"CreateToken"}}
	*/

	log.Println(string(strValue))
}

//---------------------------------------------------------

func TestCreateToken04(t *testing.T) {
	decideToRun()
	t.Cleanup(closeServer)

	ownerToken := getOwnerToken()
	doTestCreateToken04(t, ownerToken, "create", http.MethodPost)
	doTestCreateToken04(t, user03TokenTmp, "create", http.MethodPost)
	doTestCreateToken04(t, user03TokenTmp, "create", http.MethodGet)
	doTestCreateToken04(t, user03TokenTmp, "createToken", http.MethodPost)
}

func doTestCreateToken04(t *testing.T, ownerToken, path, method string) {
	url := getBaseUrl() + path + "?user-id=" + userId03 + "&token=" + ownerToken +
		"&permission=" + strconv.Itoa(int(sibylValues.Owner))
	// create a new token
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		t.Fatal(err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Log("got unexpected status code: ", res.StatusCode)
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

	if len(valueMap) < 1 {
		t.Fatal("got unexpected result: ", valueMap)
		return
	}

	token, _ := valueMap["result"].(map[string]interface{})
	if token == nil {
		t.Fatal("token cannot be nil")
	}
	user03TokenTmp, _ = token["hash"].(string)

	/* Expected result should be something like this:
	2021/10/25 21:43:31 {"result":{"id":895373440,"hash":"895373440:WXXwxWYzr4-Gp8NNo7N7-Rx1cs3xmEWPQk1n_rHplfwbkUssvTGlHbmKnK8T7eWc","permission":3,"created_at":"2021-10-25T21:41:33.695577585Z"},"success":true,"error":null}
	*/

	log.Println(string(strValue))
}

//---------------------------------------------------------

// TestCreateToken01 test to create a token using the owner token with http
// headers. The token should be created successfully.
func TestCreateToken05(t *testing.T) {
	decideToRun()
	t.Cleanup(closeServer)

	ownerToken := getOwnerToken()
	doTestCreateToken05(t, ownerToken, "create", http.MethodPost)
	doTestCreateToken05(t, ownerToken, "createToken", http.MethodPost)
	doTestCreateToken05(t, ownerToken, "create", http.MethodGet)
	doTestCreateToken05(t, ownerToken, "createToken", http.MethodGet)
}

func doTestCreateToken05(t *testing.T, ownerToken, path, method string) {
	// create a new token
	req, err := http.NewRequest(http.MethodPost, getBaseUrl()+path, nil)
	if err != nil {
		t.Fatal(err)
		return
	}

	req.Header.Set("token", ownerToken)
	req.Header.Set("user-id", userId04)
	req.Header.Set("perm", strconv.Itoa(int(sibylValues.Owner)))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatal("got unexpected status code: ", res.StatusCode)
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

	if len(valueMap) < 1 {
		t.Fatal("got unexpected result: ", valueMap)
		return
	}

	token, _ := valueMap["result"].(map[string]interface{})
	if token == nil {
		t.Fatal("token cannot be nil")
	}
	user04TokenTmp, _ = token["hash"].(string)

	/* Expected result should be something like this:
	2021/10/25 22:00:53 {"result":{"id":792109647,"hash":"792109647:i2g0Nw-FVPw50HCASCSNeHgTaMnqdYFE8m1ohdTInH_qrsgOcWEdDxRa7ocMuz0w","permission":3,"created_at":"2021-10-25T22:00:53.478450044Z"},"success":true,"error":null}
	*/

	log.Println(string(strValue))
}

//---------------------------------------------------------

func TestCreateToken06(t *testing.T) {
	decideToRun()
	t.Cleanup(closeServer)

	ownerToken := getOwnerToken()
	doTestCreateToken06(t, ownerToken, "create", http.MethodPost)
	doTestCreateToken06Wrong(t, user05TokenTmp, "create", http.MethodPost)
	doTestCreateToken06Wrong(t, user05TokenTmp, "create", http.MethodGet)
	doTestCreateToken06Wrong(t, user05TokenTmp, "createToken", http.MethodPost)
}

func doTestCreateToken06(t *testing.T, ownerToken, path, method string) {
	url := getBaseUrl() + path + "?user-id=" + userId01 + "&token=" + ownerToken +
		"&permission=" + strconv.Itoa(int(sibylValues.Inspector))
	// create a new token
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		t.Fatal(err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		strValue, _ := ioutil.ReadAll(res.Body)
		t.Fatal("got unexpected status code: ", res.StatusCode,
			"and this value: ", strValue)
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

	if len(valueMap) < 1 {
		t.Fatal("got unexpected result: ", valueMap)
		return
	}

	token, _ := valueMap["result"].(map[string]interface{})
	if token == nil {
		t.Fatal("token cannot be nil")
	}
	user05TokenTmp, _ = token["hash"].(string)
	/* Expected result should be something like this:
		4
	2021/10/25 20:19:40 {"result":{"id":1341091260,"hash":"1341091260:MiDvOMnpj8Tf7HK6OLVqVMIJg7F4on9Tyr6mRFhtVpesncgMidjc8BbN6etulbfq","permission":0,"created_at":"2021-10-25T19:54:51.194525003Z"},"success":true,"error":null}
	*/

	log.Println(string(strValue))
}

func doTestCreateToken06Wrong(t *testing.T, ownerToken, path, method string) {
	url := getBaseUrl() + path + "?user-id=" + userId01 + "&token=" + ownerToken
	// create a new token
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		t.Fatal(err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		t.Fatal("got unexpected status code: ", res.StatusCode)
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

	/* Expected result should be something like this:
	2021/10/25 20:48:37 {"result":null,"success":false,"error":{"code":502,"message":"Permission Denied","origin":"CreateToken"}}
	*/

	log.Println(string(strValue))
}

//---------------------------------------------------------

func TestCreateToken07(t *testing.T) {
}

//---------------------------------------------------------
