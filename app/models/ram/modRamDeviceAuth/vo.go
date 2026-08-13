package modRamDeviceAuth

// DeviceAuthVo 设备授权发起响应
type DeviceAuthVo struct {
	DeviceCode      string `json:"deviceCode"`
	UserCode        string `json:"userCode"`
	VerificationUri string `json:"verificationUri"`
	ExpiresIn       int    `json:"expiresIn"`
	Interval        int    `json:"interval"`
	CancelToken     string `json:"cancelToken,omitempty"`
}

// DeviceStatusVo 设备轮询状态响应
type DeviceStatusVo struct {
	Status       string `json:"status"`                 // pending / approved / denied / expired
	AccessToken  string `json:"accessToken,omitempty"`  // 仅在 approved 后签发
	RefreshToken string `json:"refreshToken,omitempty"` // 仅在 approved 后签发
}
