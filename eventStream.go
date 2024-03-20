package inutil

type EventStreamHandlerInput struct {
	WriteChannel *chan EventStreamMessage
}
type EventStreamMessage struct {
	Data []byte
}

func EventStreamHandler(input EventStreamHandlerInput) HandlerFunc {
	return func(c *Context) {
		c.Writer.Header().Set(HeaderContentType, TextEventStream)
		c.Writer.Header().Set(HeaderCacheControl, NoCache)
		c.Writer.Flush()

		for message := range *input.WriteChannel {
			c.Writer.Write(message.Data)

		}
	}
}
