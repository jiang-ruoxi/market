package cos

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// ObjectGetACLResult is the result of GetObjectACL
type ObjectGetACLResult = ACLXml

// GetACL Get Object ACL接口实现使用API读取Object的ACL表，只有所有者有权操作。
//
// https://www.qcloud.com/document/product/436/7744
func (s *ObjectService) GetACL(ctx context.Context, name string, id ...string) (*ObjectGetACLResult, *Response, error) {
	var u string
	if len(id) == 1 {
		u = fmt.Sprintf("/%s?acl&versionId=%s", encodeURIComponent(name), id[0])
	} else if len(id) == 0 {
		u = fmt.Sprintf("/%s?acl", encodeURIComponent(name))
	} else {
		return nil, nil, errors.New("wrong params")
	}
	var res ObjectGetACLResult
	sendOpt := sendOptions{
		baseURL: s.client.BaseURL.BucketURL,
		uri:     u,
		method:  http.MethodGet,
		result:  &res,
	}
	resp, err := s.client.doRetry(ctx, &sendOpt)
	if err == nil {
		decodeACL(resp, &res)
	}
	return &res, resp, err
}

// ObjectPutACLOptions the options of put object acl
type ObjectPutACLOptions struct {
	Header *ACLHeaderOptions `url:"-" xml:"-"`
	Body   *ACLXml           `url:"-" header:"-"`
}

// PutACL 使用API写入Object的ACL表，您可以通过Header："x-cos-acl", "x-cos-grant-read" ,
// "x-cos-grant-write" ,"x-cos-grant-full-control"传入ACL信息，
// 也可以通过body以XML格式传入ACL信息，但是只能选择Header和Body其中一种，否则，返回冲突。
//
// Put Object ACL是一个覆盖操作，传入新的ACL将覆盖原有ACL。只有所有者有权操作。
//
// "x-cos-acl"：枚举值为public-read，private；public-read意味这个Object有公有读私有写的权限，
// private意味这个Object有私有读写的权限。
//
// "x-cos-grant-read"：意味被赋予权限的用户拥有该Object的读权限
//
// "x-cos-grant-write"：意味被赋予权限的用户拥有该Object的写权限
//
// "x-cos-grant-full-control"：意味被赋予权限的用户拥有该Object的读写权限
//
// https://www.qcloud.com/document/product/436/7748
func (s *ObjectService) PutACL(ctx context.Context, name string, opt *ObjectPutACLOptions, id ...string) (*Response, error) {
	var u string
	if len(id) == 1 {
		u = fmt.Sprintf("/%s?acl&versionId=%s", encodeURIComponent(name), id[0])
	} else if len(id) == 0 {
		u = fmt.Sprintf("/%s?acl", encodeURIComponent(name))
	} else {
		return nil, errors.New("wrong params")
	}
	header := opt.Header
	body := opt.Body
	if body != nil {
		header = nil
	}
	sendOpt := sendOptions{
		baseURL:   s.client.BaseURL.BucketURL,
		uri:       u,
		method:    http.MethodPut,
		optHeader: header,
		body:      body,
	}
	resp, err := s.client.doRetry(ctx, &sendOpt)
	return resp, err
}
