package inutil

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
	gc *gin.Context
}

func convertContextFromGin(gc *gin.Context) *Context {
	return &Context{gc, gc}
}

func convertContextToGin(c *Context) *gin.Context {
	return c.gc
}

func (c *Context) HandleError(err error) bool {
	if err != nil {
		LogErrorF("HandleError: %v", err)
		c.Error(err)
		return true
	}
	return false
}

func (c *Context) Body(output any) error {
	ct := c.Request.Header.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		switch mediaType {
		case ApplicationJSON:
			dec := json.NewDecoder(c.Request.Body)
			err := dec.Decode(output)
			return err
		}
	} else {
		return errors.New(Error_ContentTypeNotSet)
	}
	logInternal("body")
	return nil
}

func JSON[T any](payload Return[any], c *Context) {
	c.gc.JSON(payload.Status, payload)
}
