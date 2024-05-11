package oss_router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/kysion/oss-library/oss_controller"
	"github.com/kysion/oss-library/oss_interface"
)

func ModulesGroup(modules oss_interface.IModules, group *ghttp.RouterGroup) *ghttp.RouterGroup {
	OssGroup(modules, group)
	OssFileGroup(modules, group)

	return group
}

func OssGroup(modules oss_interface.IModules, group *ghttp.RouterGroup) *ghttp.RouterGroup {
	routePrefix := modules.GetConfig().RoutePrefix + "/oss"
	controller := oss_controller.Oss(modules)

	group.POST(routePrefix+"/createProvider", controller.CreateProvider)
	group.POST(routePrefix+"/queryProviderByNo", controller.QueryProviderByNo)
	group.POST(routePrefix+"/queryProviderList", controller.QueryProviderList)
	group.POST(routePrefix+"/getProviderById", controller.GetProviderById)

	group.POST(routePrefix+"/getAppConfigById", controller.GetAppConfigById)
	group.POST(routePrefix+"/createAppConfig", controller.CreateAppConfig)
	group.POST(routePrefix+"/getAppAvailableNumber", controller.GetAppAvailableNumber)
	// group.POST(routePrefix+"/updateAppNumber", controller.UpdateAppNumber)

	group.POST(routePrefix+"/registerBucket", controller.RegisterBucket)
	group.POST(routePrefix+"/getBucketById", controller.GetBucketById)

	return group
}

func OssFileGroup(modules oss_interface.IModules, group *ghttp.RouterGroup) *ghttp.RouterGroup {
	routePrefix := modules.GetConfig().RoutePrefix + "/oss/file"
	controller := oss_controller.OssFile(modules)

	group.POST(routePrefix+"/uploadFile", controller.UploadFile)
	group.POST(routePrefix+"/downloadFile", controller.DownloadFile)
	group.POST(routePrefix+"/partPutUploadFile", controller.PartPutUpload)
	group.POST(routePrefix+"/multipartPartUpload", controller.MultipartPartUpload)
	group.POST(routePrefix+"/partPutDownload", controller.PartPutDownload)
	group.POST(routePrefix+"/getFile", controller.GetFile)
	group.POST(routePrefix+"/deleteFile", controller.DeleteFile)
	group.POST(routePrefix+"/deleteFileList", controller.DeleteFileList)
	group.POST(routePrefix+"/queryFiles", controller.QueryFiles)
	group.POST(routePrefix+"/getFileSingURL", controller.GetFileSingURL)
	return group

}
