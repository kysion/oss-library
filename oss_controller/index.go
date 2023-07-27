package oss_controller

import (
	"github.com/kysion/oss-library/oss_controller/internal"
	"github.com/kysion/oss-library/oss_interface/i_controller"
)

type (
	OssController     internal.OssController
	OssFileController internal.OssFileController
)

type CoController struct {
	Oss     i_controller.IOss
	OssFile i_controller.IOssFile
}

var (
	Oss     = internal.Oss
	OssFile = internal.OssFile
)
