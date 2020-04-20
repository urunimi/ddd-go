package dto

const CodeOK = 0

// BaseRequest struct definition
type BaseRequest struct {
	//DeviceId    string `form:"deviceId" validate:"required"`
	DeviceModel string `form:"deviceModel" validate:"required"`
	Locale      string `form:"locale" validate:"required"`
	SdkVersion  int    `form:"sdkVersion" validate:"required"`
	VersionCode int    `form:"versionCode"`
}
