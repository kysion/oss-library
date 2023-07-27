package oss_tencent

import (
	"context"
	"github.com/kysion/oss-library/internal/logic/osser"
	"github.com/kysion/oss-library/oss_interface"
	"github.com/kysion/oss-library/oss_model"
	"github.com/kysion/oss-library/oss_model/oss_dao"
	"github.com/kysion/oss-library/oss_model/oss_enum"
)

type sOssTencent struct {
	modules oss_interface.IModules
	dao     *oss_dao.XDao
	osser.OSSer
}

func NewOssTencent(modules oss_interface.IModules) oss_interface.IOssTencent {
	return &sOssTencent{
		modules: modules,
		dao:     modules.Dao(),
	}
}

// CreateProvider 添加渠道商
func (s *sOssTencent) CreateProvider(ctx context.Context, info *oss_model.OssServiceProviderConfig) (*oss_model.OssServiceProviderConfig, error) {
	info.ProviderNo = oss_enum.Oss.Type.Tencent.Code()
	info.ProviderName = oss_enum.Oss.Type.Tencent.Description()
	// 或者还有其他拓展字段...

	provider, err := s.modules.OssServiceProviderConfig().CreateProvider(ctx, info)

	return provider, err
}

// RegisterBucket 添加Bucket存储空间配置
func (s *sOssTencent) RegisterBucket(ctx context.Context, info *oss_model.OssBucketConfig) (bool, error) {
	info.ProviderNo = oss_enum.Oss.Type.Tencent.Code() // aliyun

	result, err := s.modules.OssBucketConfig().CreateBucket(ctx, info)

	return result, err
}

// ---------------------------------------- 文件操作 ---------------------------------------------------
// 创建BucketClient对象

// UploadFile 上传文件
func (s *sOssTencent) UploadFile(ctx context.Context, info *oss_model.PutObject) (bool, error) {
	return false, nil
}

// PartPutUpload 断点续传文件
func (s *sOssTencent) PartPutUpload(ctx context.Context, info *oss_model.PartPutObject) (bool, error) {
	return false, nil
}

// MultipartPartUpload 分片上传，将要上传的较大文件（Object）分成多个分片（Part）来分别上传
func (s *sOssTencent) MultipartPartUpload(ctx context.Context, info *oss_model.MultipartPartPutObject) (bool, error) {
	return false, nil
}

// DownloadFile 下载文件到本地
func (s *sOssTencent) DownloadFile(ctx context.Context, info *oss_model.DownLoadFile) (bool, error) {
	return false, nil
}

// PartPutDownload 断点续传下载
func (s *sOssTencent) PartPutDownload(ctx context.Context, info *oss_model.PartPutObject) (bool, error) {
	return false, nil
}

// GetFile 获取文件
func (s *sOssTencent) GetFile(ctx context.Context, info *oss_model.GetFile) ([]byte, error) {
	return nil, nil
}

// DeleteFile 删除文件
func (s *sOssTencent) DeleteFile(ctx context.Context, info oss_model.DeleteFile) (bool, error) {
	return false, nil
}

// DeleteFileList 批量删除文件
func (s *sOssTencent) DeleteFileList(ctx context.Context, info oss_model.DeleteFileList) (bool, error) {
	return false, nil
}

// QueryFiles 列举文件  列举指定存储空间（Bucket）下的所有文件（Object）、指定前缀的文件、指定目录下的文件和子目录("/")
func (s *sOssTencent) QueryFiles(ctx context.Context, info *oss_model.QueryFileListReq) (res []oss_model.ObjectInfoRes, err error) {
	return nil, nil
}
