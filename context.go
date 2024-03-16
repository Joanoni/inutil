package inutil

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
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
