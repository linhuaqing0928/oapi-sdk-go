// Package image code generated by lark suite oapi sdk gen
package image

import (
	"io"

	"github.com/larksuite/oapi-sdk-go/v2"
)

type Image struct {
	ImageKey *string `json:"image_key,omitempty"`
}

type ImageGetReq struct {
	ImageKey *string `query:"image_key"`
}

type ImageGetResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
	File     io.Reader `json:"-"`
	FileName string    `json:"-"`
}

type ImagePutReqBody struct {
	Image     io.Reader `json:"image,omitempty"`
	ImageType *string   `json:"image_type,omitempty"`
}

type ImagePutReq struct {
	Body *ImagePutReqBody `body:""`
}

type ImagePutResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
	Data *Image `json:"data"`
}