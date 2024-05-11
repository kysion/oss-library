package controller

import (
	"context"
	"github.com/kysion/oss-library/api/oss_api"
	"github.com/kysion/oss-library/api/oss_v1"
	"github.com/kysion/oss-library/oss_controller"
	"github.com/kysion/oss-library/oss_interface"
	"github.com/kysion/oss-library/oss_interface/i_controller"
	"github.com/kysion/oss-library/oss_model"
)

type OssFileController struct {
	i_controller.IOssFile
}

var OssFile = func(modules oss_interface.IModules) *OssFileController {
	return &OssFileController{
		oss_controller.OssFile(modules),
	}
}

func (c *OssFileController) GetModules() oss_interface.IModules {
	return c.IOssFile.GetModules()
}

// UploadFile 上传文件
func (c *OssFileController) UploadFile(ctx context.Context, req *oss_v1.UploadFileReq) (oss_api.BoolRes, error) {
	return c.IOssFile.UploadFile(ctx, &req.UploadFileReq)
}

// PartPutUpload 断点续传文件
func (c *OssFileController) PartPutUpload(ctx context.Context, req *oss_v1.PartPutUploadReq) (res oss_api.BoolRes, err error) {
	return c.IOssFile.PartPutUpload(ctx, &req.PartPutUploadReq)
}

// DownloadFile 下载文件到本地
func (c *OssFileController) DownloadFile(ctx context.Context, req *oss_v1.DownloadFileReq) (res oss_api.BoolRes, err error) {
	return c.IOssFile.DownloadFile(ctx, &req.DownloadFileReq)
}

// MultipartPartUpload 分片上传，将要上传的较大文件（Object）分成多个分片（Part）来分别上传
func (c *OssFileController) MultipartPartUpload(ctx context.Context, req *oss_v1.MultipartPartUploadReq) (res oss_api.BoolRes, err error) {
	return c.IOssFile.MultipartPartUpload(ctx, &req.MultipartPartUploadReq)
}

// PartPutDownload 断点续传下载
func (c *OssFileController) PartPutDownload(ctx context.Context, req *oss_v1.PartPutDownloadReq) (res oss_api.BoolRes, err error) {
	return c.IOssFile.PartPutDownload(ctx, &req.PartPutDownloadReq)
}

// GetFile 获取文件
func (c *OssFileController) GetFile(ctx context.Context, req *oss_v1.GetFileReq) (res oss_api.ByteRes, err error) {
	return c.IOssFile.GetFile(ctx, &req.GetFileReq)
}

// DeleteFile 删除文件
func (c *OssFileController) DeleteFile(ctx context.Context, req *oss_v1.DeleteFileReq) (res oss_api.BoolRes, err error) {
	return c.IOssFile.DeleteFile(ctx, &req.DeleteFileReq)
}

// DeleteFileList 批量删除文件
func (c *OssFileController) DeleteFileList(ctx context.Context, req *oss_v1.DeleteFileListReq) (res oss_api.BoolRes, err error) {
	return c.IOssFile.DeleteFileList(ctx, &req.DeleteFileListReq)
}

// QueryFiles 列举文件  列举指定存储空间（Bucket）下的所有文件（Object）、指定前缀的文件、指定目录下的文件和子目录("/")
func (c *OssFileController) QueryFiles(ctx context.Context, req *oss_v1.QueryFilesReq) (oss_model.ObjectInfoListRes, error) {
	return c.IOssFile.QueryFiles(ctx, &req.QueryFilesReq)
}

// GetFileSingURL 获取文件的签名访问URL
func (c *OssFileController) GetFileSingURL(ctx context.Context, req *oss_v1.GetFileSingURLReq) (oss_api.StringRes, error) {
	return c.IOssFile.GetFileSingURL(ctx, &req.GetFileSingURLReq)
}

// GetObjectToFileWithURL 根据文件访问URL获取文件
func (c *OssFileController) GetObjectToFileWithURL(ctx context.Context, req *oss_v1.GetObjectToFileWithURLReq) (oss_api.BoolRes, error) {
	return c.IOssFile.GetObjectToFileWithURL(ctx, &req.GetObjectToFileWithURLReq)
}
