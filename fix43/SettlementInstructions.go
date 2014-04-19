package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SettlementInstructions msg type = T.
type SettlementInstructions struct {
	message.Message
}

//SettlementInstructionsBuilder builds SettlementInstructions messages.
type SettlementInstructionsBuilder struct {
	message.MessageBuilder
}

//CreateSettlementInstructionsBuilder returns an initialized SettlementInstructionsBuilder with specified required fields.
func CreateSettlementInstructionsBuilder(
	settlinstid field.SettlInstID,
	settlinsttranstype field.SettlInstTransType,
	settlinstrefid field.SettlInstRefID,
	settlinstmode field.SettlInstMode,
	settlinstsource field.SettlInstSource,
	allocaccount field.AllocAccount,
	transacttime field.TransactTime) SettlementInstructionsBuilder {
	var builder SettlementInstructionsBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(settlinstid)
	builder.Body.Set(settlinsttranstype)
	builder.Body.Set(settlinstrefid)
	builder.Body.Set(settlinstmode)
	builder.Body.Set(settlinstsource)
	builder.Body.Set(allocaccount)
	builder.Body.Set(transacttime)
	return builder
}

//SettlInstID is a required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstID() (field.SettlInstID, errors.MessageRejectError) {
	var f field.SettlInstID
	err := m.Body.Get(&f)
	return f, err
}

//SettlInstTransType is a required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstTransType() (field.SettlInstTransType, errors.MessageRejectError) {
	var f field.SettlInstTransType
	err := m.Body.Get(&f)
	return f, err
}

//SettlInstRefID is a required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstRefID() (field.SettlInstRefID, errors.MessageRejectError) {
	var f field.SettlInstRefID
	err := m.Body.Get(&f)
	return f, err
}

//SettlInstMode is a required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstMode() (field.SettlInstMode, errors.MessageRejectError) {
	var f field.SettlInstMode
	err := m.Body.Get(&f)
	return f, err
}

//SettlInstSource is a required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstSource() (field.SettlInstSource, errors.MessageRejectError) {
	var f field.SettlInstSource
	err := m.Body.Get(&f)
	return f, err
}

//AllocAccount is a required field for SettlementInstructions.
func (m SettlementInstructions) AllocAccount() (field.AllocAccount, errors.MessageRejectError) {
	var f field.AllocAccount
	err := m.Body.Get(&f)
	return f, err
}

//IndividualAllocID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) IndividualAllocID() (field.IndividualAllocID, errors.MessageRejectError) {
	var f field.IndividualAllocID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) ClOrdID() (field.ClOrdID, errors.MessageRejectError) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a non-required field for SettlementInstructions.
func (m SettlementInstructions) TradeDate() (field.TradeDate, errors.MessageRejectError) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//AllocID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) AllocID() (field.AllocID, errors.MessageRejectError) {
	var f field.AllocID
	err := m.Body.Get(&f)
	return f, err
}

//LastMkt is a non-required field for SettlementInstructions.
func (m SettlementInstructions) LastMkt() (field.LastMkt, errors.MessageRejectError) {
	var f field.LastMkt
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) TradingSessionID() (field.TradingSessionID, errors.MessageRejectError) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) TradingSessionSubID() (field.TradingSessionSubID, errors.MessageRejectError) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//Side is a non-required field for SettlementInstructions.
func (m SettlementInstructions) Side() (field.Side, errors.MessageRejectError) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SecurityType() (field.SecurityType, errors.MessageRejectError) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//EffectiveTime is a non-required field for SettlementInstructions.
func (m SettlementInstructions) EffectiveTime() (field.EffectiveTime, errors.MessageRejectError) {
	var f field.EffectiveTime
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for SettlementInstructions.
func (m SettlementInstructions) TransactTime() (field.TransactTime, errors.MessageRejectError) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for SettlementInstructions.
func (m SettlementInstructions) NoPartyIDs() (field.NoPartyIDs, errors.MessageRejectError) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//StandInstDbType is a non-required field for SettlementInstructions.
func (m SettlementInstructions) StandInstDbType() (field.StandInstDbType, errors.MessageRejectError) {
	var f field.StandInstDbType
	err := m.Body.Get(&f)
	return f, err
}

//StandInstDbName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) StandInstDbName() (field.StandInstDbName, errors.MessageRejectError) {
	var f field.StandInstDbName
	err := m.Body.Get(&f)
	return f, err
}

//StandInstDbID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) StandInstDbID() (field.StandInstDbID, errors.MessageRejectError) {
	var f field.StandInstDbID
	err := m.Body.Get(&f)
	return f, err
}

//SettlDeliveryType is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SettlDeliveryType() (field.SettlDeliveryType, errors.MessageRejectError) {
	var f field.SettlDeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//SettlDepositoryCode is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SettlDepositoryCode() (field.SettlDepositoryCode, errors.MessageRejectError) {
	var f field.SettlDepositoryCode
	err := m.Body.Get(&f)
	return f, err
}

//SettlBrkrCode is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SettlBrkrCode() (field.SettlBrkrCode, errors.MessageRejectError) {
	var f field.SettlBrkrCode
	err := m.Body.Get(&f)
	return f, err
}

//SettlInstCode is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstCode() (field.SettlInstCode, errors.MessageRejectError) {
	var f field.SettlInstCode
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySettlAgentName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SecuritySettlAgentName() (field.SecuritySettlAgentName, errors.MessageRejectError) {
	var f field.SecuritySettlAgentName
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySettlAgentCode is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SecuritySettlAgentCode() (field.SecuritySettlAgentCode, errors.MessageRejectError) {
	var f field.SecuritySettlAgentCode
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySettlAgentAcctNum is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SecuritySettlAgentAcctNum() (field.SecuritySettlAgentAcctNum, errors.MessageRejectError) {
	var f field.SecuritySettlAgentAcctNum
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySettlAgentAcctName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SecuritySettlAgentAcctName() (field.SecuritySettlAgentAcctName, errors.MessageRejectError) {
	var f field.SecuritySettlAgentAcctName
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySettlAgentContactName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SecuritySettlAgentContactName() (field.SecuritySettlAgentContactName, errors.MessageRejectError) {
	var f field.SecuritySettlAgentContactName
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySettlAgentContactPhone is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SecuritySettlAgentContactPhone() (field.SecuritySettlAgentContactPhone, errors.MessageRejectError) {
	var f field.SecuritySettlAgentContactPhone
	err := m.Body.Get(&f)
	return f, err
}

//CashSettlAgentName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CashSettlAgentName() (field.CashSettlAgentName, errors.MessageRejectError) {
	var f field.CashSettlAgentName
	err := m.Body.Get(&f)
	return f, err
}

//CashSettlAgentCode is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CashSettlAgentCode() (field.CashSettlAgentCode, errors.MessageRejectError) {
	var f field.CashSettlAgentCode
	err := m.Body.Get(&f)
	return f, err
}

//CashSettlAgentAcctNum is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CashSettlAgentAcctNum() (field.CashSettlAgentAcctNum, errors.MessageRejectError) {
	var f field.CashSettlAgentAcctNum
	err := m.Body.Get(&f)
	return f, err
}

//CashSettlAgentAcctName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CashSettlAgentAcctName() (field.CashSettlAgentAcctName, errors.MessageRejectError) {
	var f field.CashSettlAgentAcctName
	err := m.Body.Get(&f)
	return f, err
}

//CashSettlAgentContactName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CashSettlAgentContactName() (field.CashSettlAgentContactName, errors.MessageRejectError) {
	var f field.CashSettlAgentContactName
	err := m.Body.Get(&f)
	return f, err
}

//CashSettlAgentContactPhone is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CashSettlAgentContactPhone() (field.CashSettlAgentContactPhone, errors.MessageRejectError) {
	var f field.CashSettlAgentContactPhone
	err := m.Body.Get(&f)
	return f, err
}

//PaymentMethod is a non-required field for SettlementInstructions.
func (m SettlementInstructions) PaymentMethod() (field.PaymentMethod, errors.MessageRejectError) {
	var f field.PaymentMethod
	err := m.Body.Get(&f)
	return f, err
}

//PaymentRef is a non-required field for SettlementInstructions.
func (m SettlementInstructions) PaymentRef() (field.PaymentRef, errors.MessageRejectError) {
	var f field.PaymentRef
	err := m.Body.Get(&f)
	return f, err
}

//CardHolderName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CardHolderName() (field.CardHolderName, errors.MessageRejectError) {
	var f field.CardHolderName
	err := m.Body.Get(&f)
	return f, err
}

//CardNumber is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CardNumber() (field.CardNumber, errors.MessageRejectError) {
	var f field.CardNumber
	err := m.Body.Get(&f)
	return f, err
}

//CardStartDate is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CardStartDate() (field.CardStartDate, errors.MessageRejectError) {
	var f field.CardStartDate
	err := m.Body.Get(&f)
	return f, err
}

//CardExpDate is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CardExpDate() (field.CardExpDate, errors.MessageRejectError) {
	var f field.CardExpDate
	err := m.Body.Get(&f)
	return f, err
}

//CardIssNo is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CardIssNo() (field.CardIssNo, errors.MessageRejectError) {
	var f field.CardIssNo
	err := m.Body.Get(&f)
	return f, err
}

//PaymentDate is a non-required field for SettlementInstructions.
func (m SettlementInstructions) PaymentDate() (field.PaymentDate, errors.MessageRejectError) {
	var f field.PaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//PaymentRemitterID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) PaymentRemitterID() (field.PaymentRemitterID, errors.MessageRejectError) {
	var f field.PaymentRemitterID
	err := m.Body.Get(&f)
	return f, err
}
