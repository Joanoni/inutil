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
	defer PrintInternalFunction()()

	Print("1")
	if len(contentType) > 0 {
		Print("2")
		switch contentType[0] {
		case ApplicationJSON:
			Print("3")
			dec := json.NewDecoder(input)
			Print("4")
			err := dec.Decode(output)
			Print("5")
			if HandleError(err) {
				outerr = ReturnInternalServerError(ErrsFromError(err))
			}
			Print("6")
			return
		default:
			outerr = ReturnInternalServerError(ErrsContentTypeNotSet)
			Print("7")
			return
		}
	} else {
		outerr = ReturnInternalServerError(ErrsContentTypeNotSet)
		Print("8")
		return
	}
}
