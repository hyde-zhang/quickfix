package fix41

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
	symbol field.Symbol) QuoteRequestBuilder {
	var builder QuoteRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(quotereqid)
	builder.Body.Set(symbol)
	return builder
}

//QuoteReqID is a required field for QuoteRequest.
func (m QuoteRequest) QuoteReqID() (field.QuoteReqID, errors.MessageRejectError) {
	var f field.QuoteReqID
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a required field for QuoteRequest.
func (m QuoteRequest) Symbol() (field.Symbol, errors.MessageRejectError) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for QuoteRequest.
func (m QuoteRequest) SymbolSfx() (field.SymbolSfx, errors.MessageRejectError) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for QuoteRequest.
func (m QuoteRequest) SecurityID() (field.SecurityID, errors.MessageRejectError) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//IDSource is a non-required field for QuoteRequest.
func (m QuoteRequest) IDSource() (field.IDSource, errors.MessageRejectError) {
	var f field.IDSource
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for QuoteRequest.
func (m QuoteRequest) SecurityType() (field.SecurityType, errors.MessageRejectError) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for QuoteRequest.
func (m QuoteRequest) MaturityMonthYear() (field.MaturityMonthYear, errors.MessageRejectError) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDay is a non-required field for QuoteRequest.
func (m QuoteRequest) MaturityDay() (field.MaturityDay, errors.MessageRejectError) {
	var f field.MaturityDay
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for QuoteRequest.
func (m QuoteRequest) PutOrCall() (field.PutOrCall, errors.MessageRejectError) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for QuoteRequest.
func (m QuoteRequest) StrikePrice() (field.StrikePrice, errors.MessageRejectError) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for QuoteRequest.
func (m QuoteRequest) OptAttribute() (field.OptAttribute, errors.MessageRejectError) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for QuoteRequest.
func (m QuoteRequest) SecurityExchange() (field.SecurityExchange, errors.MessageRejectError) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for QuoteRequest.
func (m QuoteRequest) Issuer() (field.Issuer, errors.MessageRejectError) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for QuoteRequest.
func (m QuoteRequest) SecurityDesc() (field.SecurityDesc, errors.MessageRejectError) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//PrevClosePx is a non-required field for QuoteRequest.
func (m QuoteRequest) PrevClosePx() (field.PrevClosePx, errors.MessageRejectError) {
	var f field.PrevClosePx
	err := m.Body.Get(&f)
	return f, err
}

//Side is a non-required field for QuoteRequest.
func (m QuoteRequest) Side() (field.Side, errors.MessageRejectError) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty is a non-required field for QuoteRequest.
func (m QuoteRequest) OrderQty() (field.OrderQty, errors.MessageRejectError) {
	var f field.OrderQty
	err := m.Body.Get(&f)
	return f, err
}

//FutSettDate is a non-required field for QuoteRequest.
func (m QuoteRequest) FutSettDate() (field.FutSettDate, errors.MessageRejectError) {
	var f field.FutSettDate
	err := m.Body.Get(&f)
	return f, err
}

//OrdType is a non-required field for QuoteRequest.
func (m QuoteRequest) OrdType() (field.OrdType, errors.MessageRejectError) {
	var f field.OrdType
	err := m.Body.Get(&f)
	return f, err
}

//FutSettDate2 is a non-required field for QuoteRequest.
func (m QuoteRequest) FutSettDate2() (field.FutSettDate2, errors.MessageRejectError) {
	var f field.FutSettDate2
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty2 is a non-required field for QuoteRequest.
func (m QuoteRequest) OrderQty2() (field.OrderQty2, errors.MessageRejectError) {
	var f field.OrderQty2
	err := m.Body.Get(&f)
	return f, err
}
