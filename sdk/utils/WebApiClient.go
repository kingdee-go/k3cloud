package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var (
	Cookie         string = ""
	defaultHeaders        = map[string]string{
		"Accept":         "application/json",
		"Content-Type":   "application/json",
		"Accept-Charset": "utf-8",
		"User-Agent":     "Kingdee/Golang WebApi SDK (compatible: K3Cloud 7.3+)",
	}
)

type WebApiClient struct {
	ExecuteFunc     func(url string, headers map[string]string, postData map[string]string, format string)
	BuildHeaderFunc func(service_url string, config map[string]string)
	// Parse_strToInterfaceFunc func(resp_str string)
}

func Execute(url string, headers map[string]interface{}, postData map[string]interface{}) string {
	json_data, err := json.Marshal(postData)
	if err != nil {
		fmt.Println(err)
		return "err"
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json_data))
	if err != nil {
		panic(err)
	}

	for k, v := range defaultHeaders {
		req.Header.Add(k, v)
	}
	for k, v := range headers {
		str := fmt.Sprintf("%v", v)
		req.Header.Add(k, str)
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println(err)
		return "err"

	} else {
		body, _ := io.ReadAll(resp.Body)
		// fmt.Println(resp.Header)
		Cookie = strings.Join(resp.Header["Set-Cookie"], ";")
		return string(body)
	}
}

func BuildHeader(service_url string, config map[string]string) map[string]interface{} {
	path_url := strings.Replace(service_url, "/", "%2F", -1)
	time_stamp := _int64ToString(time.Now().UTC().Unix())
	nonce := time_stamp
	arr := strings.Split(config["appid"], "_")
	client_id := arr[0]
	client_sec := _decodeAppSecret(arr[1])
	api_sign := "POST\n" + path_url + "\n\nx-api-nonce:" + nonce + "\nx-api-timestamp:" + time_stamp + "\n"
	app_data := config["acct_id"] + "," + config["username"] + "," + _getMapValue_str(config, "lcid", "2052") + "," + _getMapValue_str(config, "org_num", "0")

	header := map[string]interface{}{
		"X-Api-Auth-Version": "2.0",
		"X-Api-ClientID":     client_id,
		"x-api-timestamp":    time_stamp,
		"x-api-nonce":        nonce,
		"x-api-signheaders":  "x-api-timestamp,x-api-nonce",
		"X-Api-Signature":    _kdHmacSHA256(api_sign, client_sec),
		"X-Kd-Appkey":        config["appid"],
		"X-Kd-Appdata":       base64.StdEncoding.EncodeToString([]byte(app_data)),
		"X-Kd-Signature":     _kdHmacSHA256(config["appid"]+app_data, config["appsecret"]),
	}
	return header
	// return make(map[string]interface{})
}

// func Parse_strToInterface(resp_str string) interface{} {
// 	var resp_parsed interface{}
// 	err := json.Unmarshal([]byte(resp_str), &resp_parsed)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return resp_parsed
// }

func _kdHmacSHA256(content string, sign_key string) string {
	mac := hmac.New(sha256.New, []byte(sign_key))
	_, _ = mac.Write([]byte(content))
	hexData := mac.Sum(nil)
	hexStr := hex.EncodeToString(hexData)
	return base64.StdEncoding.EncodeToString([]byte(hexStr))
}

func _decodeAppSecret(secret string) string {
	if len(secret) != 32 {
		fmt.Println("sec:", secret, "is not 32 char")
		return secret
	}
	// xor_code := "0054s3974c62343787b09ca7d32e5debce72" // example from official Python SDK
	xor_code := "0054f397c6234378b09ca7d3e5debce7" // example from official Java SDK
	base64_xor := _xorCode(secret, xor_code)
	return base64.StdEncoding.EncodeToString([]byte(base64_xor))
}

func _xorCode(str string, key string) string {
	xor_code := []byte(key)
	base64_decode := []byte(str)
	base64_xor := make([]byte, len(base64_decode))
	for i := 0; i < len(base64_decode); i++ {
		base64_xor[i] = base64_decode[i] ^ xor_code[i]
	}
	return string(base64_xor)
}

func _getMapValue_str(m map[string]string, key string, defaultVaue string) string {
	if m[key] == "" {
		return defaultVaue
	} else {
		return m[key]
	}
}

func _int64ToString(value int64) string {
	return fmt.Sprintf("%v", value)
}
