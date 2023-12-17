package models

type Role string

const (
	RoleUser  Role = "USER"
	RoleModel Role = "MODEL"
	RoleNil   Role = ""
)

func (r Role) ToOpenAIRole() string {
	switch r {
	case RoleUser:
		return "user"
	case RoleModel:
		return "assistant"
	default:
		return ""
	}
}

type MimeType string

const (
	MimeImagePng    MimeType = "image/png"
	MimeImageJpeg   MimeType = "image/jpeg"
	MimeImageWebP   MimeType = "image/webp"
	MimeImageHEIC   MimeType = "image/heic"
	MimeImageHEIF   MimeType = "image/heif"
	MimeVideoMov    MimeType = "video/mov"
	MimeVideoMpeg   MimeType = "video/mpeg"
	MimeVideoMp4    MimeType = "video/mp4"
	MimeVideoMpg    MimeType = "video/mpg"
	MimeVideoAvi    MimeType = "video/avi"
	MimeVideoWmv    MimeType = "video/wmv"
	MimeVideoMpegps MimeType = "video/mpegps"
	MimeVideoFlv    MimeType = "video/flv"
)
