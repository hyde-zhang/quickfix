package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//QuoteRequest msg type = R.
type QuoteRequest struct {
	message.Message
}

//QuoteRequestBuilder builds QuoteRequest messages.
type QuoteRequestBuilder struct {
	message.MessageBuilder
}

//CreateQuoteRequestBuilder returns an initialized QuoteRequestBuilder with specified required fields.
func CreateQuoteRequestBuilder(
	quotereqid field.QuoteReqID,
	norelatedsym field.NoRelatedSym) QuoteRequestBuilder {
	var builder QuoteRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(quotereqid)
	builder.Body.Set(norelatedsym)
	return builder
}

//QuoteReqID is a required field for QuoteRequest.
func (m QuoteRequest) QuoteReqID() (field.QuoteReqID, errors.MessageRejectError) {
	var f field.QuoteReqID
	err := m.Body.Get(&f)
	return f, err
}

//NoRelatedSym is a required field for QuoteRequest.
func (m QuoteRequest) NoRelatedSym() (field.NoRelatedSym, errors.MessageRejectError) {
	var f field.NoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}
