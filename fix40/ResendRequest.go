package fix40

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//ResendRequest msg type = 2.
type ResendRequest struct {
	message.Message
}

//ResendRequestBuilder builds ResendRequest messages.
type ResendRequestBuilder struct {
	message.MessageBuilder
}

//CreateResendRequestBuilder returns an initialized ResendRequestBuilder with specified required fields.
func CreateResendRequestBuilder(
	beginseqno field.BeginSeqNo,
	endseqno field.EndSeqNo) ResendRequestBuilder {
	var builder ResendRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(beginseqno)
	builder.Body.Set(endseqno)
	return builder
}

//BeginSeqNo is a required field for ResendRequest.
func (m ResendRequest) BeginSeqNo() (field.BeginSeqNo, errors.MessageRejectError) {
	var f field.BeginSeqNo
	err := m.Body.Get(&f)
	return f, err
}

//EndSeqNo is a required field for ResendRequest.
func (m ResendRequest) EndSeqNo() (field.EndSeqNo, errors.MessageRejectError) {
	var f field.EndSeqNo
	err := m.Body.Get(&f)
	return f, err
}
