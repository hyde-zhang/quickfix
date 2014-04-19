package fix41

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
	settlinstmode field.SettlInstMode,
	settlinstsource field.SettlInstSource,
	allocaccount field.AllocAccount,
	transacttime field.TransactTime) SettlementInstructionsBuilder {
	var builder SettlementInstructionsBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(settlinstid)
	builder.Body.Set(settlinsttranstype)
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

//SettlLocation is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SettlLocation() (field.SettlLocation, errors.MessageRejectError) {
	var f field.SettlLocation
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

//ClientID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) ClientID() (field.ClientID, errors.MessageRejectError) {
	var f field.ClientID
	err := m.Body.Get(&f)
	return f, err
}

//ExecBroker is a non-required field for SettlementInstructions.
func (m SettlementInstructions) ExecBroker() (field.ExecBroker, errors.MessageRejectError) {
	var f field.ExecBroker
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
