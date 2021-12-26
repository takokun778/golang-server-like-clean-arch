package xxx

import (
	"context"
	"fmt"

	pbXxx "xxx/proto/xxx"
)

type (
	CreateKey  string
	UpdateKey  string
	ReadKey    string
	ReadAllKey string
	DeleteKey  string
)

const (
	CreateCtxKey  CreateKey  = "create"
	UpdateCtxKey  UpdateKey  = "update"
	ReadCtxKey    ReadKey    = "read"
	ReadAllCtxKey ReadAllKey = "readAll"
	DeleteCtxKey  DeleteKey  = "delete"
)

func setCreateResponse(src context.Context, res *pbXxx.CreateResponse) context.Context {
	return context.WithValue(src, CreateCtxKey, res)
}

func getCreateResponse(ctx context.Context) (*pbXxx.CreateResponse, error) {
	v := ctx.Value(CreateCtxKey)

	res, ok := v.(*pbXxx.CreateResponse)

	if !ok {
		return nil, fmt.Errorf("response not found")
	}

	return res, nil
}

func setUpdateResponse(src context.Context, res *pbXxx.UpdateResponse) context.Context {
	return context.WithValue(src, UpdateCtxKey, res)
}

func getUpdateResponse(ctx context.Context) (*pbXxx.UpdateResponse, error) {
	v := ctx.Value(UpdateCtxKey)

	res, ok := v.(*pbXxx.UpdateResponse)

	if !ok {
		return nil, fmt.Errorf("response not found")
	}

	return res, nil
}

func setReadResponse(src context.Context, res *pbXxx.ReadResponse) context.Context {
	return context.WithValue(src, ReadCtxKey, res)
}

func getReadResponse(ctx context.Context) (*pbXxx.ReadResponse, error) {
	v := ctx.Value(ReadCtxKey)

	res, ok := v.(*pbXxx.ReadResponse)

	if !ok {
		return nil, fmt.Errorf("response not found")
	}

	return res, nil
}

func setReadAllResponse(src context.Context, res *pbXxx.ReadAllResponse) context.Context {
	return context.WithValue(src, ReadAllCtxKey, res)
}

func getReadAllResponse(ctx context.Context) (*pbXxx.ReadAllResponse, error) {
	v := ctx.Value(ReadAllCtxKey)

	res, ok := v.(*pbXxx.ReadAllResponse)

	if !ok {
		return nil, fmt.Errorf("response not found")
	}

	return res, nil
}

func setDeleteResponse(src context.Context, res *pbXxx.DeleteResponse) context.Context {
	return context.WithValue(src, DeleteCtxKey, res)
}

func getDeleteResponse(ctx context.Context) (*pbXxx.DeleteResponse, error) {
	v := ctx.Value(DeleteCtxKey)

	res, ok := v.(*pbXxx.DeleteResponse)

	if !ok {
		return nil, fmt.Errorf("response not found")
	}

	return res, nil
}
