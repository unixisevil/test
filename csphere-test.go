package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var ControllerUrl = "http://localhost:1016"
var ApiKey = "178578adaf483ae43bf402aec27960edae3cac25"
var C2Id = "C2Id=MTQzNzQ0ODIzMHxEdi1CQkFFQ180SUFBUkFCRUFBQU5QLUNBQUVHYzNSeWFXNW5EQVFBQWtsRUJuTjBjbWx1Wnd3YUFCZzFOV0ZqWWpSa1lqTTFPV1ZoTlRBd01UY3dNREF3TURNPXzoVw-Absq0_z7vhNri5uMiGCvAlTw3Y7e1kpYrQXHAJw=="

func handleError(desc string, err error) {
	if err != nil {
		log.Fatalln(desc, err)
	}
}

func UserCreate(email, passwd string) string {
	b := new(bytes.Buffer)
	user := struct {
		Email  string `json:"email"`
		Passwd string `json:"password"`
	}{
		Email:  email,
		Passwd: passwd,
	}
	err := json.NewEncoder(b).Encode(&user)
	handleError("encode to buffer failed: ", err)
	url := ControllerUrl + "/api/users"
	req, err := http.NewRequest("POST", url, b)
	handleError("make req failed: ", err)
	req.Header.Add("Cookie", C2Id)
	resp, err := http.DefaultClient.Do(req)
	handleError("http.Post  /api/users  failed: ", err)
	defer resp.Body.Close()
	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	handleError("json decode  /api/users  failed: ", err)
	fmt.Println(res)
	return res["id"].(string)
}
func UserDel(uid string) bool {
	url := ControllerUrl + "/api/users/" + uid
	req, err := http.NewRequest("DELETE", url, nil)
	handleError("make del user req failed:", err)
	req.Header.Add("Cookie", C2Id)
	resp, err := http.DefaultClient.Do(req)
	handleError("Do  failed: ", err)
	fmt.Println(resp.Status)
	var ok bool
	if resp.StatusCode == 204 {
		ok = true
	}
	return ok
}

func UserExist(uid string) bool {
}
func main() {
	//UserCreate("zhangsan@nicescale.com", "12345678")
	UserDel("55ae21fa21e3f8001e000006")
}
