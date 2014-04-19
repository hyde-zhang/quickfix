package fix40

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//OrderCancelRequest msg type = F.
type OrderCancelRequest struct {
	message.Message
}

//OrderCancelRequestBuilder builds OrderCancelRequest messages.
type OrderCancelRequestBuilder struct {
	message.MessageBuilder
}

//CreateOrderCancelRequestBuilder returns an initialized OrderCancelRequestBuilder with specified required fields.
func CreateOrderCancelRequestBuilder(
	origclordid field.OrigClOrdID,
	clordid field.ClOrdID,
	cxltype field.CxlType,
	symbol field.Symbol,
	side field.Side,
	orderqty field.OrderQty) OrderCancelRequestBuilder {
	var builder OrderCancelRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(origclordid)
	builder.Body.Set(clordid)
	builder.Body.Set(cxltype)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	builder.Body.Set(orderqty)
	return builder
}

//OrigClOrdID is a required field for OrderCancelRequest.
func (m OrderCancelRequest) OrigClOrdID() (field.OrigClOrdID, errors.MessageRejectError) {
	var f field.OrigClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//OrderID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) OrderID() (field.OrderID, errors.MessageRejectError) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a required field for OrderCancelRequest.
func (m OrderCancelRequest) ClOrdID() (field.ClOrdID, errors.MessageRejectError) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//ListID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) ListID() (field.ListID, errors.MessageRejectError) {
	var f field.ListID
	err := m.Body.Get(&f)
	return f, err
}

//CxlType is a required field for OrderCancelRequest.
func (m OrderCancelRequest) CxlType() (field.CxlType, errors.MessageRejectError) {
	var f field.CxlType
	err := m.Body.Get(&f)
	return f, err
}

//ClientID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) ClientID() (field.ClientID, errors.MessageRejectError) {
	var f field.ClientID
	err := m.Body.Get(&f)
	return f, err
}

//ExecBroker is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) ExecBroker() (field.ExecBroker, errors.MessageRejectError) {
	var f field.ExecBroker
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a required field for OrderCancelRequest.
func (m OrderCancelRequest) Symbol() (field.Symbol, errors.MessageRejectError) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SymbolSfx() (field.SymbolSfx, errors.MessageRejectError) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecurityID() (field.SecurityID, errors.MessageRejectError) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//IDSource is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) IDSource() (field.IDSource, errors.MessageRejectError) {
	var f field.IDSource
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) Issuer() (field.Issuer, errors.MessageRejectError) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecurityDesc() (field.SecurityDesc, errors.MessageRejectError) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Side is a required field for OrderCancelRequest.
func (m OrderCancelRequest) Side() (field.Side, errors.MessageRejectError) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty is a required field for OrderCancelRequest.
func (m OrderCancelRequest) OrderQty() (field.OrderQty, errors.MessageRejectError) {
	var f field.OrderQty
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}
