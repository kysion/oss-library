package i_controller

import (
	"context"
	"github.com/kysion/oss-library/api/oss_api"
	"github.com/kysion/oss-library/oss_model"
)

type IOssFile interface {
	iModule

	UploadFile(ctx context.Context, req *oss_api.UploadFileReq) (res oss_api.BoolRes, err error)

	PartPutUpload(ctx context.Context, req *oss_api.PartPutUploadReq) (res oss_api.BoolRes, err error)

	DownloadFile(ctx context.Context, req *oss_api.DownloadFileReq) (res oss_api.BoolRes, err error)

	MultipartPartUpload(ctx context.Context, req *oss_api.MultipartPartUploadReq) (res oss_api.BoolRes, err error)

	PartPutDownload(ctx context.Context, req *oss_api.PartPutDownloadReq) (res oss_api.BoolRes, err error)

	GetFile(ctx context.Context, req *oss_api.GetFileReq) (oss_api.ByteRes, error)

	DeleteFile(ctx context.Context, req *oss_api.DeleteFileReq) (res oss_api.BoolRes, err error)

	DeleteFileList(ctx context.Context, req *oss_api.DeleteFileListReq) (res oss_api.BoolRes, err error)

	QueryFiles(ctx context.Context, req *oss_api.QueryFilesReq) (oss_model.ObjectInfoListRes, error)

	CopyFileToPath(ctx context.Context, req *oss_api.CopyFileToPathReq) (res oss_api.BoolRes, err error)

	GetFileSingURL(ctx context.Context, req *oss_api.GetFileSingURLReq) (oss_api.StringRes, error)

	GetObjectToFileWithURL(ctx context.Context, req *oss_api.GetObjectToFileWithURLReq) (oss_api.BoolRes, error)
}
