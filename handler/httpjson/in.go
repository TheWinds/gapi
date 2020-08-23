package httpjson

import (
	"github.com/zhiduoke/gapi"
	"github.com/zhiduoke/gapi/metadata"
	"github.com/zhiduoke/gapi/proto/jtop"
	"github.com/zhiduoke/gapi/proto/kvpb"
	"io/ioutil"
)

func (h *Handler) handleInput(msg *metadata.Message, ctx *gapi.Context) ([]byte, error) {
	contentType := ctx.Request().Header.Get("Content-Type")
	var (
		pb  []byte
		err error
	)

	if contentType == "application/json" {
		// json
		var body []byte
		body, err = ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			return nil, err
		}
		pb, err = jtop.Encode(msg, body)
		if err != nil {
			return nil, err
		}
	} else {
		err = ctx.Request().ParseForm()
		if err != nil {
			return nil, err
		}
		// form
		pb, err = kvpb.Encode(msg, &formKv{ctx: ctx})
		if err != nil {
			return nil, err
		}
	}
	ctxpb, err := kvpb.Encode(msg, ctx)
	if err != nil {
		return nil, err
	}
	pb = append(pb, ctxpb...)
	return pb, nil
}

type formKv struct {
	ctx *gapi.Context
}

func (kv *formKv) Get(key string) (string, bool) {
	v := kv.ctx.Request().FormValue(key)
	return v, len(v) > 0
}
