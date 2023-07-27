package oss_model

type OSSType string

// Oss 不同的Oss渠道商
type Oss struct {
	OssType OSSType    `json:"ossType"`
	Local   LocalOSS   `json:"local"`
	Aliyun  AliyunOSS  `json:"aliyun"`
	Huawei  HuaweiOSS  `json:"huawei"`
	Tencent TencentOSS `json:"tencent"`
	Qiniu   QiniuOSS   `json:"qiniu"`
	Minio   MinioOSS   `json:"minio"`
}

type LocalOSS struct {
	Path string `json:"path"` // 本地文件路径
}

type MinioOSS struct {
	Id       string `json:"id"`
	Path     string `json:"path"`
	Token    string `json:"token"`
	Bucket   string `json:"bucket"`
	Secret   string `json:"secret"`
	Endpoint string `json:"endpoint"`
	UseSsl   bool   `json:"useSsl"`
}

type HuaweiOSS struct {
	Path      string `json:"path"`
	Bucket    string `json:"bucket"`
	Endpoint  string `json:"endpoint"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

type AliyunOSS struct {
	Endpoint        string `json:"endpoint"`        // 地域节点 oss-cn-shenzhen.aliyuncs.com
	BasePath        string `json:"basePath"`        // 域名：https://xxx.oss-cn-shenzhen.aliyuncs.com/ (可以自绑定)
	BucketUrl       string `json:"bucketUrl"`       // Bucket 域名 xxxx.oss-cn-shenzhen.aliyuncs.com
	BucketName      string `json:"bucketName"`      // 桶名：xxxx-bucket
	AccessKeyId     string `json:"accessKeyId"`     // 身份认证ID
	AccessKeySecret string `json:"accessKeySecret"` // 身份认证秘钥
}

type TencentOSS struct {
	Bucket     string `json:"bucket"`
	Region     string `json:"region"`
	BaseURL    string `json:"baseURL"`
	SecretID   string `json:"secretID"`
	SecretKey  string `json:"secretKey"`
	PathPrefix string `json:"pathPrefix"`
}

type QiniuOSS struct {
	Zone          string `json:"zone"`          // 存储区域
	Bucket        string `json:"bucket"`        // 空间名称
	ImgPath       string `json:"imgPath"`       // CDN加速域名
	UseHTTPS      bool   `json:"useHTTPS"`      // 是否使用https
	AccessKey     string `json:"accessKey"`     // 秘钥AK
	SecretKey     string `json:"secretKey"`     // 秘钥SK
	UseCdnDomains bool   `json:"useCdnDomains"` // 上传是否使用CDN上传加速
}
