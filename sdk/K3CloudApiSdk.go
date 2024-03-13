package sdk

import (
	"github.com/kingdee-go/k3cloud/sdk/utils"
)

var (
	hostUrl string            = ""
	config  map[string]string = make(map[string]string)
)

type K3CloudApiSdk struct {
	InitFunc              func()
	LoginForPasswordFunc  func()
	LoginForSecretFunc    func()
	GetHeadersFunc        func()
	ViewFunc              func()
	ExecuteBillQueryFunc  func()
	BillQueryFunc         func()
	QueryBusinessInfoFunc func()
	GetDataCenterListFunc func()
	SaveFunc              func()
	BatchSaveFunc         func()
	AuditFunc             func()
	UnAuditFunc           func()
	SubmitFunc            func()
	OperationFunc         func()
	PushFunc              func()
	DraftFunc             func()
	DeleteFunc            func()
	AllocateFunc          func()
	CancelAllocateFunc    func()
	FlexSaveFunc          func()
	SendMsgFunc           func()
	GroupSaveFunc         func()
	DisassemblyFunc       func()
	WorkflowAuditFunc     func()
	QueryGroupInfoFunc    func()
	GroupDeleteFunc       func()
	GetSysReportDataFunc  func()
}

func Init(conf map[string]string) *K3CloudApiSdk {
	inst := new(K3CloudApiSdk)
	hostUrl = conf["host_url"]
	config = conf
	return inst
}

// 登录: 用户名+密码
func LoginForPassword() string {
	url := hostUrl + utils.LOGIN_API
	postData := map[string]interface{}{
		"acctid":   config["acct_id"],
		"username": config["username"],
		"password": config["password"],
		"lcid":     config["lcid"],
	}
	resp_str := utils.Execute(url, make(map[string]interface{}), postData)
	return resp_str
}

// 登录: 第三方授权应用ID+应用密钥
func LoginForSecret() string {
	url := hostUrl + utils.LOGIN_API_APP_SECRET
	postData := map[string]interface{}{
		"acctid":    config["acct_id"],
		"username":  config["username"],
		"appid":     config["appid"],
		"appsecret": config["appsecret"],
		"lcid":      config["lcid"],
	}
	resp_str := utils.Execute(url, make(map[string]interface{}), postData)
	return resp_str
}

// 登录: 签名
func GetHeaders(url string) map[string]interface{} {
	// cookie
	if config["auth_type"] == utils.USER_ID_PASSWORD {
		if utils.Cookie == "" {
			LoginForPassword()
		}
	}
	if config["auth_type"] == utils.APP_ID_SECRET {
		if utils.Cookie == "" {
			LoginForSecret()
		}
	}
	// headers
	if config["auth_type"] == utils.API_SIGNATURE {
		return utils.BuildHeader(url, config)
	} else {
		return map[string]interface{}{"Cookie": utils.Cookie}
	}
	// return make(map[string]interface{})
}

// 详情
func View(formId string, data map[string]string) string {
	url := hostUrl + utils.VIEW_API
	postData := map[string]interface{}{
		"formid": formId,
		"data":   data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}

// 单据查询
func ExecuteBillQuery(data map[string]string) string {
	url := hostUrl + utils.EXECUTEBILLQUERY_API
	postData := map[string]interface{}{
		"data": data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}

// 单据查询(json) （官方在2023.9.4新增此接口）
func BillQuery(data map[string]string) string {
	url := hostUrl + utils.BILLQUERY_API
	postData := map[string]interface{}{
		"data": data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}

// 元数据查询
func QueryBusinessInfo(data map[string]string) string {
	url := hostUrl + utils.QUERYBUSINESSINFO_API
	postData := map[string]interface{}{
		"data": data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}

// 获取数据中心列表
func GetDataCenterList() string {
	url := hostUrl + utils.GETDATACENTERLIST_API
	resp_str := utils.Execute(url, GetHeaders(url), make(map[string]interface{}))
	return resp_str
}

// 保存
func Save(formId string, data map[string]string) string {
	url := hostUrl + utils.SAVE_API
	postData := map[string]interface{}{
		"formid": formId,
		"data":   data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}

// 批量保存
func BatchSave(formId string, data map[string]string) string {
	url := hostUrl + utils.BATCHSAVE_API
	postData := map[string]interface{}{
		"formid": formId,
		"data":   data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}

// 审核
func Audit(formId string, data map[string]string) string {
	url := hostUrl + utils.AUDIT_API
	postData := map[string]interface{}{
		"formid": formId,
		"data":   data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}

// 反审核
func UnAudit(formId string, data map[string]string) string {
	url := hostUrl + utils.UNAUDIT_API
	postData := map[string]interface{}{
		"formid": formId,
		"data":   data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}

// 提交
func Submit(formId string, data map[string]string) string {
	url := hostUrl + utils.SUBMIT_API
	postData := map[string]interface{}{
		"formid": formId,
		"data":   data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}

// 操作
func Operation(formId string, opNumber string, data map[string]string) string {
	url := hostUrl + utils.EXCUTEOPERATION_API
	postData := map[string]interface{}{
		"formid":   formId,
		"opNumber": opNumber,
		"data":     data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}

// 下推
func Push(formId string, data map[string]string) string {
	url := hostUrl + utils.PUSH_API
	postData := map[string]interface{}{
		"formid": formId,
		"data":   data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}

// 暂存
func Draft(formId string, data map[string]string) string {
	url := hostUrl + utils.DRAFT_API
	postData := map[string]interface{}{
		"formid": formId,
		"data":   data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}

// 删除
func Delete(formId string, data map[string]string) string {
	url := hostUrl + utils.DELETE_API
	postData := map[string]interface{}{
		"formid": formId,
		"data":   data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}

// 分配
func Allocate(formId string, data map[string]string) string {
	url := hostUrl + utils.ALLOCATE_API
	postData := map[string]interface{}{
		"formid": formId,
		"data":   data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}

// 取消分配
func CancelAllocate(formId string, data map[string]string) string {
	url := hostUrl + utils.CANCEL_ALLOCATE_API
	postData := map[string]interface{}{
		"formid": formId,
		"data":   data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}

// 弹性域保存
func FlexSave(formId string, data map[string]string) string {
	url := hostUrl + utils.FLEXSAVE_API
	postData := map[string]interface{}{
		"formid": formId,
		"data":   data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}

// 发送消息
func SendMsg(data map[string]string) string {
	url := hostUrl + utils.SENDMSG_API
	postData := map[string]interface{}{
		"data": data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}

// 分组保存
func GroupSave(formId string, data map[string]string) string {
	url := hostUrl + utils.GROUPSAVE_API
	postData := map[string]interface{}{
		"formid": formId,
		"data":   data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}

// 拆单
func Disassembly(formId string, data map[string]string) string {
	url := hostUrl + utils.DISASSEMBLY_API
	postData := map[string]interface{}{
		"formid": formId,
		"data":   data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}

// 工作流审批
func WorkflowAudit(data map[string]string) string {
	url := hostUrl + utils.WORKFLOWAUDIT_API
	postData := map[string]interface{}{
		"data": data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}

// 查询分组信息
func QueryGroupInfo(data map[string]string) string {
	url := hostUrl + utils.QUERYGROUPINFO_API
	postData := map[string]interface{}{
		"data": data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}

// 分组删除
func GroupDelete(data map[string]string) string {
	url := hostUrl + utils.GROUPDELETE_API
	postData := map[string]interface{}{
		"data": data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}

// 查询报表数据
func GetSysReportData(formId string, data map[string]string) string {
	url := hostUrl + utils.GET_SYS_REPORT_DATA_API
	postData := map[string]interface{}{
		"formid": formId,
		"data":   data,
	}
	resp_str := utils.Execute(url, GetHeaders(url), postData)
	return resp_str
}
