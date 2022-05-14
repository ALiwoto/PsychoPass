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
	"strings"
	"testing"
	"time"

	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"
)

func TestGetRawMultiBanSampleData(t *testing.T) {
	data := sibylValues.MultiBanRawData{}
	var tmpInfo *sibylValues.MultiBanUserInfo
	for i := 0; i < 10; i++ {
		tmpInfo = &sibylValues.MultiBanUserInfo{
			UserId:     int64(i + 100),
			Reason:     "spam and raid",
			Message:    "https://t.me/telegram/505050",
			Source:     "https://t.me/src/5123",
			TargetType: sibylValues.EntityType(i % 2),
		}
		data.Users = append(data.Users, *tmpInfo)
	}

	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
		return
	}

	myStr := string(b)

	log.Println(myStr)
}

func GetRawMultiBanSampleData(t *testing.T) []byte {
	data := sibylValues.MultiBanRawData{}
	var tmpInfo *sibylValues.MultiBanUserInfo
	for i := 0; i < 10; i++ {
		tmpInfo = &sibylValues.MultiBanUserInfo{
			UserId:     int64(i + 100),
			Reason:     "spam and raid",
			Message:    "https://t.me/telegram/505050",
			Source:     "https://t.me/src/5123",
			TargetType: sibylValues.EntityType(i % 2),
		}
		data.Users = append(data.Users, *tmpInfo)
	}

	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
		return nil
	}

	return b
}

func TestMultiBan01(t *testing.T) {
	data := GetRawMultiBanSampleData(t)
	decideToRun()
	t.Cleanup(closeServer)
	ownerToken := getOwnerToken()
	body := strings.NewReader(string(data))

	time.Sleep(900 * time.Millisecond)

	req, err := http.NewRequest(http.MethodPost, getBaseUrl()+"multiBan", body)
	if err != nil {
		t.Fatal(err)
		return
	}

	req.Header.Add("token", ownerToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return
	}

	defer resp.Body.Close()

	var b []byte

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return
	}

	log.Println(string(b))

	time.Sleep(2900 * time.Millisecond)

	req, err = http.NewRequest(http.MethodGet, getBaseUrl()+"getInfo", body)
	if err != nil {
		t.Fatal(err)
		return
	}

	req.Header.Add("token", ownerToken)
	req.Header.Add("user-id", "109")

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return
	}

	defer resp.Body.Close()

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return
	}

	log.Println(string(b))

	log.Println("multiBan: done")
}
