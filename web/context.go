package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

type Context struct {
	req *http.Request
	rsp http.ResponseWriter

	pathParams map[string]string
	queryVals  url.Values
}

type HandleFunc func(ctx *Context)

func (c *Context) RspJson(statusCode int, val any) error {

	bs, err := json.Marshal(val)
	if err != nil {
		return err
	}

	c.rsp.WriteHeader(statusCode)
	_, err = c.rsp.Write(bs)

	return err
}

func (c *Context) OkJson(val any) error {
	return c.RspJson(http.StatusOK, val)
}

func (c *Context) PathVal(key string) StringVal {
	if val, ok := c.pathParams[key]; ok {
		return StringVal{
			val: val,
			err: nil,
		}
	}

	return StringVal{
		val: "",
		err: errors.New("[ctx] path key not exist"),
	}
}

func (c *Context) QueryVal(key string) StringVal {
	if c.queryVals == nil {
		c.queryVals = c.req.URL.Query()
	}

	if vals, ok := c.queryVals[key]; ok {
		return StringVal{
			val: vals[0],
			err: nil,
		}
	}

	return StringVal{
		val: "",
		err: errors.New("[ctx] query key not exist"),
	}
}

func (c *Context) FormVal(key string) StringVal {
	if err := c.req.ParseForm(); err != nil {
		return StringVal{
			val: "",
			err: err,
		}
	}

	return StringVal{
		val: c.req.FormValue(key),
		err: nil,
	}
}

func (c *Context) BindJson(val any) error {
	if c.req.Body == nil {
		return errors.New("[ctx] request body is nil")
	}

	decoder := json.NewDecoder(c.req.Body)
	return decoder.Decode(val)
}

type StringVal struct {
	val string
	err error
}

func (s StringVal) AsInt() (int, error) {
	if s.err != nil {
		return 0, s.err
	}

	return strconv.Atoi(s.val)
}

func (s StringVal) AsInt64() (int64, error) {
	if s.err != nil {
		return 0, s.err
	}

	return strconv.ParseInt(s.val, 10, 64)
}

func (s StringVal) AsFloat64() (float64, error) {
	if s.err != nil {
		return 0, s.err
	}

	return strconv.ParseFloat(s.val, 64)
}
