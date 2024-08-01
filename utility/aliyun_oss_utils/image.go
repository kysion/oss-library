package aliyun_oss_utils

import (
	"github.com/gogf/gf/v2/encoding/gbase64"
	"strings"
)

// WatermarkEncode 水印编码
func WatermarkEncode(text string) string {
	/*
		水印编码
		在添加水印操作中，文字水印的文字内容、文字字体、图片水印的水印图片名称等参数需要进行URL安全的Base64编码。编码步骤如下：

		1、将内容编码成Base64。

		2、 将结果中的部分编码替换。

			将结果中的加号（+）替换成短划线（-）。

			将结果中的正斜线（/）替换成下划线（_）。

			将结果中尾部的所有等号（=）省略。

			推荐通过base64url encoder对文字水印的文字内容、文字颜色、文字字体、图片水印的水印图片名称等参数进行编码。
	*/
	srcBase64 := string(gbase64.Encode([]byte(text)))

	srcBase64 = strings.Replace(srcBase64, "+", "-", -1)
	srcBase64 = strings.Replace(srcBase64, "/", "_", -1)
	srcBase64 = strings.Replace(srcBase64, "=", "", -1)

	return srcBase64
}
