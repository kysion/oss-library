package oss_api

import (
	"github.com/kysion/oss-library/oss_model"
)

type UploadFileReq struct {
	oss_model.PutObject
}

type PartPutUploadReq struct {
	oss_model.PartPutObject
}

type MultipartPartUploadReq struct {
	oss_model.MultipartPartPutObject
}

type DownloadFileReq struct {
	oss_model.DownLoadFile
}

type PartPutDownloadReq struct {
	oss_model.PartPutObject
}

type GetFileReq struct {
	oss_model.GetFile
}

type DeleteFileReq struct {
	oss_model.DeleteFile
}

type DeleteFileListReq struct {
	oss_model.DeleteFileList
}

type QueryFilesReq struct {
	oss_model.QueryFileListReq
}

type GetFileSingURLReq struct {
	oss_model.GetFileSingURL
}

type GetObjectToFileWithURLReq struct {
	oss_model.GetObjectToFileWithURL
}

type ByteRes []byte
