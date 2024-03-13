package main

import (
	"fmt"

	"github.com/kingdee-go/k3cloud/sdk"
	"github.com/tidwall/gjson"
)

func main() {
	fmt.Println("kingdee-k3cloud-go ...")

	// ------ init
	config := map[string]string{
		"auth_type": "3",                                        // 授权类型：1 用户名+密码；2 第三方授权应用ID+应用密钥；3 签名；
		"host_url":  "http||https://xxxxxxxxxxxxxxxxx/k3cloud/", // 金蝶授权请求地址
		"acct_id":   "xxxxxxxxxx",                               // 账户ID
		"username":  "xxxxxxxxxx",                               // 用户名（授权类型为1时必须）
		"password":  "xxxxxxxxxx",                               // 密码（授权类型为1时必须）
		"appid":     "xxxxxxxxxx",                               // 应用ID（授权类型为2或3时必须）
		"appsecret": "xxxxxxxxxx",                               // 应用Secret（授权类型为2或3时必须）
		"lcid":      "2052",                                     // 账套语系，默认2052
	}
	sdk.Init(config)

	// ------ send request (基础管理->基础资料->物料->单据查询)
	postData := map[string]string{
		"FormId":       "BD_MATERIAL",
		"FieldKeys":    "FMATERIALID,FNumber,FName",
		"FilterString": "FDocumentStatus = 'C'",
		"OrderString":  "",
		"TopRowCount":  "0",
		"StartRow":     "0",
		"Limit":        "2",
		"SubSystemId":  "",
	}
	resp_str := sdk.ExecuteBillQuery(postData)
	fmt.Printf("resopnse = %v\n", resp_str)

	// ------ parse response with 3rd part module: gjson, 【Recommended】
	fmt.Printf("response.hasError = %v\n", gjson.Get(resp_str, "0.0.Result.ResponseStatus.ErrorCode").Exists())
	fmt.Printf("response.errorMsg = %v\n", gjson.Get(resp_str, "0.0.Result.ResponseStatus.Errors.0.Message"))
	fmt.Printf("len(response) = %v\n", gjson.Get(resp_str, "#"))
	fmt.Printf("response[0][2] = %v\n", gjson.Get(resp_str, "0.2"))

	// ------ or parse response with built-in module: json, 【Not recommended】
	/*
		var resp_parsed interface{}
		err := json.Unmarshal([]byte(resp_str), &resp_parsed)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp_parsed)
	*/
}
