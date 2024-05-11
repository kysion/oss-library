package oss_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/oss-library/api/oss_api"
)

type UploadFileReq struct {
	g.Meta ` method:"post" summary:"上传文件" tags:"Oss文件"`

	oss_api.UploadFileReq
}

type PartPutUploadReq struct {
	g.Meta `method:"post" summary:"断点续传" tags:"Oss文件"`

	oss_api.PartPutUploadReq
}

type MultipartPartUploadReq struct {
	g.Meta `method:"post" summary:"分片上传" tags:"Oss文件"`

	oss_api.MultipartPartUploadReq
}

type DownloadFileReq struct {
	g.Meta `method:"post" summary:"下载文件" tags:"Oss文件"`

	oss_api.DownloadFileReq
}

type PartPutDownloadReq struct {
	g.Meta `method:"post" summary:"断点续传下载文件" tags:"Oss文件"`

	oss_api.PartPutDownloadReq
}

type GetFileReq struct {
	g.Meta `method:"post" summary:"获取文件" tags:"Oss文件"`

	oss_api.GetFileReq
}

type DeleteFileReq struct {
	g.Meta `method:"post" summary:"删除文件" tags:"Oss文件"`

	oss_api.DeleteFileReq
}

type DeleteFileListReq struct {
	g.Meta `method:"post" summary:"批量删除文件" tags:"Oss文件"`

	oss_api.DeleteFileListReq
}

type QueryFilesReq struct {
	g.Meta `method:"post" summary:"列举文件" tags:"Oss文件"`

	oss_api.QueryFilesReq
}

type GetFileSingURLReq struct {
	g.Meta `method:"post" summary:"获取文件访问URL" tags:"Oss文件"`

	oss_api.GetFileSingURLReq
}

type GetObjectToFileWithURLReq struct {
	g.Meta `method:"post" summary:"根据文件访问URL获取文件" tags:"Oss文件"`

	oss_api.GetObjectToFileWithURLReq
}
