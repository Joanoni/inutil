package context

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

func convertContextFromGin(gc *gin.Context) *model.Context {
	return &model.Context{gc, gc}
}

func convertContextToGin(c *model.Context) *gin.Context {
	return c.gc
}

func HandleError(err error) bool {
	if err != nil {
		PrintErrorF("HandleError: %v", err)
		return true
	}
	return false
}

func (c *model.Context) JSON(output Caio) {
	c.gc.JSON(output.GetStatusCode(), output.GetData())
}

func (c *model.Context) Body(output any) (outerr ReturnStructError) {
	defer PrintInternalFunction()()
	return parseBody(c.Request.Body, c.Request.Header[HeaderContentType], output)
}

func parseBody(input io.ReadCloser, contentType []string, output any) (outerr ReturnStructError) {
	if len(contentType) > 0 {
		switch contentType[0] {
		case ApplicationJSON:
			dec := json.NewDecoder(input)
			err := dec.Decode(output)
			outerr = ReturnInternalServerError(err.Error())
			return
		default:
			outerr = ReturnInternalServerError(errs.ContentTypeNotSet)
			return
		}
	} else {
		outerr = ReturnInternalServerError(errs.ContentTypeNotSet)
		return
	}
}
