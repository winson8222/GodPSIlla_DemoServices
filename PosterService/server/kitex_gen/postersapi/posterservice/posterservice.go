// Code generated by Kitex v0.5.2. DO NOT EDIT.

package posterservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	postersapi "posterservice/kitex_gen/postersapi"
)

func serviceInfo() *kitex.ServiceInfo {
	return posterServiceServiceInfo
}

var posterServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "PosterService"
	handlerType := (*postersapi.PosterService)(nil)
	methods := map[string]kitex.MethodInfo{
		"getuniqueusernames": kitex.NewMethodInfo(getuniqueusernamesHandler, newPosterServiceGetuniqueusernamesArgs, newPosterServiceGetuniqueusernamesResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "postersapi",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.2",
		Extra:           extra,
	}
	return svcInfo
}

func getuniqueusernamesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {

	realResult := result.(*postersapi.PosterServiceGetuniqueusernamesResult)
	success, err := handler.(postersapi.PosterService).Getuniqueusernames(ctx)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPosterServiceGetuniqueusernamesArgs() interface{} {
	return postersapi.NewPosterServiceGetuniqueusernamesArgs()
}

func newPosterServiceGetuniqueusernamesResult() interface{} {
	return postersapi.NewPosterServiceGetuniqueusernamesResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Getuniqueusernames(ctx context.Context) (r *postersapi.Response, err error) {
	var _args postersapi.PosterServiceGetuniqueusernamesArgs
	var _result postersapi.PosterServiceGetuniqueusernamesResult
	if err = p.c.Call(ctx, "getuniqueusernames", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
