package modRamDeviceAuth

// DeviceAuthRequestCt 发起设备授权请求
type DeviceAuthRequestCt struct {
	ClientId string `json:"clientId" form:"clientId" label:"客户端ID" binding:"required"`
	Scope    string `json:"scope" form:"scope" label:"授权范围"`
}

// DeviceLoginCt 设备轮询 / 用户授权 / 取消 / 完成请求
type DevicePollCt struct {
	DeviceCode string `json:"deviceCode" form:"deviceCode" label:"设备码" binding:"required"`
}

// DeviceApproveCt 用户在浏览器中授权
type DeviceApproveCt struct {
	UserCode string `json:"userCode" form:"userCode" label:"用户码" binding:"required"`
}

// DeviceCancelCt 取消设备授权
type DeviceCancelCt struct {
	UserCode    string `json:"userCode" form:"userCode" label:"用户码" binding:"required"`
	CancelToken string `json:"cancelToken" form:"cancelToken" label:"取消令牌" binding:"required"`
}

// DeviceCompleteCt 完成设备授权
type DeviceCompleteCt struct {
	DeviceCode string `json:"deviceCode" form:"deviceCode" label:"设备码" binding:"required"`
}
