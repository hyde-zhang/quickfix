package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//ListStrikePrice msg type = m.
type ListStrikePrice struct {
	message.Message
}

//ListStrikePriceBuilder builds ListStrikePrice messages.
type ListStrikePriceBuilder struct {
	message.MessageBuilder
}

//CreateListStrikePriceBuilder returns an initialized ListStrikePriceBuilder with specified required fields.
func CreateListStrikePriceBuilder(
	listid field.ListID,
	totnostrikes field.TotNoStrikes,
	nostrikes field.NoStrikes) ListStrikePriceBuilder {
	var builder ListStrikePriceBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(listid)
	builder.Body.Set(totnostrikes)
	builder.Body.Set(nostrikes)
	return builder
}

//ListID is a required field for ListStrikePrice.
func (m ListStrikePrice) ListID() (field.ListID, errors.MessageRejectError) {
	var f field.ListID
	err := m.Body.Get(&f)
	return f, err
}

//TotNoStrikes is a required field for ListStrikePrice.
func (m ListStrikePrice) TotNoStrikes() (field.TotNoStrikes, errors.MessageRejectError) {
	var f field.TotNoStrikes
	err := m.Body.Get(&f)
	return f, err
}

//NoStrikes is a required field for ListStrikePrice.
func (m ListStrikePrice) NoStrikes() (field.NoStrikes, errors.MessageRejectError) {
	var f field.NoStrikes
	err := m.Body.Get(&f)
	return f, err
}
