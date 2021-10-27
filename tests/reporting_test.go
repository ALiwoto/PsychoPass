package tests_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestReportUser01(t *testing.T) {
	// run the app in anoter goroutine
	//go runApp()

	//time.Sleep(time.Millisecond * 5600)

	ownerToken := getOwnerToken()
	doTestReportUser01(t, ownerToken, "report", http.MethodPost)
	doTestReportUser01(t, ownerToken, "reportUser", http.MethodPost)
	doTestReportUser01(t, ownerToken, "report", http.MethodGet)
	doTestReportUser01(t, ownerToken, "reportUser", http.MethodGet)
}

func doTestReportUser01(t *testing.T, ownerToken, path, method string) {
	// create a new token
	req, err := http.NewRequest(http.MethodPost, getBaseUrl()+path, nil)
	if err != nil {
		t.Fatal(err)
		return
	}

	req.Header.Set("token", ownerToken)
	req.Header.Set("user-id", userId01)
	req.Header.Set("reason", "owo")
	req.Header.Set("message", "https://t.me/OnePunchDev/101824")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		strValue, _ := ioutil.ReadAll(res.Body)
		log.Println(strValue)
		valueMap := make(map[string]interface{})
		_ = json.Unmarshal(strValue, &valueMap)
		log.Println(strValue)
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

	log.Println(string(strValue))
}
