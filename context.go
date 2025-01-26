package inutil

import (
	"encoding/json"
	"io"

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
		PrintErrsF("HandleError: %v", err)
		c.Error(err)
		return true
	}
	return false
}

func (c *Context) JSON(output Caio) {
	c.gc.JSON(output.GetStatusCode(), output.GetData())
}

func (c *Context) Body(output any) (outerr ReturnStructError) {
	defer PrintInternalFunction()()
	return parseBody(c.Request.Body, c.Request.Header[HeaderContentType], output)
}
func parseBody(input io.ReadCloser, contentType []string, output any) (outerr ReturnStructError) {
	if len(contentType) > 0 {
		switch contentType[0] {
		case ApplicationJSON:
			dec := json.NewDecoder(input)
			err := dec.Decode(output)
			outerr = ReturnInternalServerError(ErrsFromError(err))
			return
		default:
			outerr = ReturnInternalServerError(ErrsContentTypeNotSet)
			return
		}
	} else {
		outerr = ReturnInternalServerError(ErrsContentTypeNotSet)
		return
	}
}
