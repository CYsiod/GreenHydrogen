package response

import "github.com/CYsiod/GreenHydrogen/server/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
