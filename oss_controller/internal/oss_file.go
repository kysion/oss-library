package internal

import (
	"context"
	"errors"
	"github.com/kysion/oss-library/api/oss_api"
	"github.com/kysion/oss-library/oss_interface"
	"github.com/kysion/oss-library/oss_interface/i_controller"
	"github.com/kysion/oss-library/oss_model"
	"github.com/kysion/oss-library/oss_model/oss_enum"
)

// OssFileController 文件存储
type OssFileController struct {
	i_controller.IOssFile
	modules oss_interface.IModules
}

// OssFile 文件存储服务
var OssFile = func(modules oss_interface.IModules) i_controller.IOssFile {
	return &OssFileController{
		modules: modules,
	}
}

// UploadFile 上传文件(普通上传)
func (c *OssFileController) UploadFile(ctx context.Context, req *oss_api.UploadFileReq) (res oss_api.BoolRes, err error) {
	var ret bool
	// 判断平台
	switch req.ProviderNo {
	case oss_enum.Oss.Type.Aliyun.Code():
		ret, err = c.modules.OssAliyun().UploadFile(ctx, &req.PutObject)
	case oss_enum.Oss.Type.Tencent.Code():
		ret, err = c.modules.OssTencent().UploadFile(ctx, &req.PutObject)

	case oss_enum.Oss.Type.Qiniu.Code():
	case oss_enum.Oss.Type.Minio.Code():
	case oss_enum.Oss.Type.Huawei.Code():
	case oss_enum.Oss.Type.Local.Code():

	default:
		return false, errors.New("抱歉，暂不支持此渠道商！")
	}

	// 实际需要返回上传的一些信息，不能只是简单的布尔值
	return ret == true, err
}

// PartPutUpload 断点续传文件
func (c *OssFileController) PartPutUpload(ctx context.Context, req *oss_api.PartPutUploadReq) (res oss_api.BoolRes, err error) {
	var ret bool
	switch req.ProviderNo {
	case oss_enum.Oss.Type.Aliyun.Code():
		ret, err = c.modules.OssAliyun().PartPutUpload(ctx, &req.PartPutObject)
	case oss_enum.Oss.Type.Tencent.Code():
		ret, err = c.modules.OssTencent().PartPutUpload(ctx, &req.PartPutObject)

	case oss_enum.Oss.Type.Qiniu.Code():
	case oss_enum.Oss.Type.Minio.Code():
	case oss_enum.Oss.Type.Huawei.Code():
	case oss_enum.Oss.Type.Local.Code():

	default:
		return false, errors.New("抱歉，暂不支持此渠道商！")
	}
	return ret == true, err
}

// MultipartPartUpload 分片上传，将要上传的较大文件（Object）分成多个分片（Part）来分别上传
func (c *OssFileController) MultipartPartUpload(ctx context.Context, req *oss_api.MultipartPartUploadReq) (res oss_api.BoolRes, err error) {
	var ret bool
	switch req.ProviderNo {
	case oss_enum.Oss.Type.Aliyun.Code():
		ret, err = c.modules.OssAliyun().MultipartPartUpload(ctx, &req.MultipartPartPutObject)
	case oss_enum.Oss.Type.Tencent.Code():

	case oss_enum.Oss.Type.Qiniu.Code():
	case oss_enum.Oss.Type.Minio.Code():
	case oss_enum.Oss.Type.Huawei.Code():
	case oss_enum.Oss.Type.Local.Code():

	default:
		return false, errors.New("抱歉，暂不支持此渠道商！")
	}
	return ret == true, err

}

// DownloadFile 下载文件到本地
func (c *OssFileController) DownloadFile(ctx context.Context, req *oss_api.DownloadFileReq) (res oss_api.BoolRes, err error) {
	var ret bool
	switch req.ProviderNo {
	case oss_enum.Oss.Type.Aliyun.Code():
		ret, err = c.modules.OssAliyun().DownloadFile(ctx, &req.DownLoadFile)
	case oss_enum.Oss.Type.Tencent.Code():

	case oss_enum.Oss.Type.Qiniu.Code():
	case oss_enum.Oss.Type.Minio.Code():
	case oss_enum.Oss.Type.Huawei.Code():
	case oss_enum.Oss.Type.Local.Code():

	default:
		return false, errors.New("抱歉，暂不支持此渠道商！")
	}
	return ret == true, err
}

// PartPutDownload 断点续传下载   没测
func (c *OssFileController) PartPutDownload(ctx context.Context, req *oss_api.PartPutDownloadReq) (res oss_api.BoolRes, err error) {
	var ret bool
	switch req.ProviderNo {
	case oss_enum.Oss.Type.Aliyun.Code():
		ret, err = c.modules.OssAliyun().PartPutDownload(ctx, &req.PartPutObject)
	case oss_enum.Oss.Type.Tencent.Code():

	case oss_enum.Oss.Type.Qiniu.Code():
	case oss_enum.Oss.Type.Minio.Code():
	case oss_enum.Oss.Type.Huawei.Code():
	case oss_enum.Oss.Type.Local.Code():

	default:
		return false, errors.New("抱歉，暂不支持此渠道商！")
	}
	return ret == true, err
}

// GetFile 获取文件
func (c *OssFileController) GetFile(ctx context.Context, req *oss_api.GetFileReq) (res oss_api.ByteRes, err error) {
	var ret []byte

	switch req.ProviderNo {
	case oss_enum.Oss.Type.Aliyun.Code():
		ret, err = c.modules.OssAliyun().GetFile(ctx, &req.GetFile)
	case oss_enum.Oss.Type.Tencent.Code():

	case oss_enum.Oss.Type.Qiniu.Code():
	case oss_enum.Oss.Type.Minio.Code():
	case oss_enum.Oss.Type.Huawei.Code():
	case oss_enum.Oss.Type.Local.Code():

	default:
		return nil, errors.New("抱歉，暂不支持此渠道商！")
	}

	// 返回字节
	return ret, err
}

// DeleteFile 删除文件
func (c *OssFileController) DeleteFile(ctx context.Context, req *oss_api.DeleteFileReq) (res oss_api.BoolRes, err error) {
	var ret bool
	switch req.ProviderNo {
	case oss_enum.Oss.Type.Aliyun.Code():
		ret, err = c.modules.OssAliyun().DeleteFile(ctx, req.DeleteFile)
	case oss_enum.Oss.Type.Tencent.Code():

	case oss_enum.Oss.Type.Qiniu.Code():

	case oss_enum.Oss.Type.Minio.Code():
	case oss_enum.Oss.Type.Huawei.Code():
	case oss_enum.Oss.Type.Local.Code():

	default:
		return false, errors.New("抱歉，暂不支持此渠道商！")
	}
	return ret == true, err
}

// DeleteFileList 批量删除文件
func (c *OssFileController) DeleteFileList(ctx context.Context, req *oss_api.DeleteFileListReq) (res oss_api.BoolRes, err error) {
	var ret bool
	switch req.ProviderNo {
	case oss_enum.Oss.Type.Aliyun.Code():
		ret, err = c.modules.OssAliyun().DeleteFileList(ctx, req.DeleteFileList)
	case oss_enum.Oss.Type.Tencent.Code():

	case oss_enum.Oss.Type.Qiniu.Code():
	case oss_enum.Oss.Type.Minio.Code():
	case oss_enum.Oss.Type.Huawei.Code():
	case oss_enum.Oss.Type.Local.Code():

	default:
		return false, errors.New("抱歉，暂不支持此渠道商！")
	}
	return ret == true, err
}

// QueryFiles 列举文件  列举指定存储空间（Bucket）下的所有文件（Object）、指定前缀的文件、指定目录下的文件和子目录("/")
func (c *OssFileController) QueryFiles(ctx context.Context, req *oss_api.QueryFilesReq) (res oss_model.ObjectInfoListRes, err error) {
	var ret []oss_model.ObjectInfoRes
	switch req.ProviderNo {
	case oss_enum.Oss.Type.Aliyun.Code():
		ret, err = c.modules.OssAliyun().QueryFiles(ctx, &req.QueryFileListReq)

	case oss_enum.Oss.Type.Tencent.Code():

	case oss_enum.Oss.Type.Qiniu.Code():
	case oss_enum.Oss.Type.Minio.Code():
	case oss_enum.Oss.Type.Huawei.Code():
	case oss_enum.Oss.Type.Local.Code():

	default:
		return nil, errors.New("抱歉，暂不支持此渠道商！")
	}
	return ret, err
}

// CopyFileToPath 将指定文件拷贝到指定位置
func (c *OssFileController) CopyFileToPath(ctx context.Context, req *oss_api.CopyFileToPathReq) (res oss_api.BoolRes, err error) {
	var ret bool
	switch req.ProviderNo {
	case oss_enum.Oss.Type.Aliyun.Code():
		ret, err = c.modules.OssAliyun().CopyFileToPath(ctx, &req.CopyFileToPath)

	case oss_enum.Oss.Type.Tencent.Code():

	case oss_enum.Oss.Type.Qiniu.Code():
	case oss_enum.Oss.Type.Minio.Code():
	case oss_enum.Oss.Type.Huawei.Code():
	case oss_enum.Oss.Type.Local.Code():

	default:
		return false, errors.New("抱歉，暂不支持此渠道商！")
	}
	return ret == true, err
}

// GetFileSingURL 获取文件的签名访问URL
func (c *OssFileController) GetFileSingURL(ctx context.Context, req *oss_api.GetFileSingURLReq) (res oss_api.StringRes, err error) {
	var ret string
	switch req.ProviderNo {
	case oss_enum.Oss.Type.Aliyun.Code():
		ret, err = c.modules.OssAliyun().GetFileSingURL(ctx, &req.GetFileSingURL, req.StyleStr)

	case oss_enum.Oss.Type.Tencent.Code():

	case oss_enum.Oss.Type.Qiniu.Code():
	case oss_enum.Oss.Type.Minio.Code():
	case oss_enum.Oss.Type.Huawei.Code():
	case oss_enum.Oss.Type.Local.Code():

	default:
		return "", errors.New("抱歉，暂不支持此渠道商！")
	}
	return (oss_api.StringRes)(ret), err
}

// GetObjectToFileWithURL 根据URL获取存储对象
func (c *OssFileController) GetObjectToFileWithURL(ctx context.Context, req *oss_api.GetObjectToFileWithURLReq) (oss_api.BoolRes, error) {
	var ret bool
	var err error
	switch req.ProviderNo {
	case oss_enum.Oss.Type.Aliyun.Code():
		ret, err = c.modules.OssAliyun().GetObjectToFileWithURL(ctx, req.GetObjectToFileWithURL)

	case oss_enum.Oss.Type.Tencent.Code():

	case oss_enum.Oss.Type.Qiniu.Code():
	case oss_enum.Oss.Type.Minio.Code():
	case oss_enum.Oss.Type.Huawei.Code():
	case oss_enum.Oss.Type.Local.Code():

	default:
		return false, errors.New("抱歉，暂不支持此渠道商！")
	}
	return ret == true, err
}
