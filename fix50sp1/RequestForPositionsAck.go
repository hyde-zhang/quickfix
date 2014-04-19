package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//RequestForPositionsAck msg type = AO.
type RequestForPositionsAck struct {
	message.Message
}

//RequestForPositionsAckBuilder builds RequestForPositionsAck messages.
type RequestForPositionsAckBuilder struct {
	message.MessageBuilder
}

//CreateRequestForPositionsAckBuilder returns an initialized RequestForPositionsAckBuilder with specified required fields.
func CreateRequestForPositionsAckBuilder(
	posmaintrptid field.PosMaintRptID,
	posreqresult field.PosReqResult,
	posreqstatus field.PosReqStatus) RequestForPositionsAckBuilder {
	var builder RequestForPositionsAckBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(posmaintrptid)
	builder.Body.Set(posreqresult)
	builder.Body.Set(posreqstatus)
	return builder
}

//PosMaintRptID is a required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PosMaintRptID() (field.PosMaintRptID, errors.MessageRejectError) {
	var f field.PosMaintRptID
	err := m.Body.Get(&f)
	return f, err
}

//PosReqID is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PosReqID() (field.PosReqID, errors.MessageRejectError) {
	var f field.PosReqID
	err := m.Body.Get(&f)
	return f, err
}

//TotalNumPosReports is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) TotalNumPosReports() (field.TotalNumPosReports, errors.MessageRejectError) {
	var f field.TotalNumPosReports
	err := m.Body.Get(&f)
	return f, err
}

//UnsolicitedIndicator is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) UnsolicitedIndicator() (field.UnsolicitedIndicator, errors.MessageRejectError) {
	var f field.UnsolicitedIndicator
	err := m.Body.Get(&f)
	return f, err
}

//PosReqResult is a required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PosReqResult() (field.PosReqResult, errors.MessageRejectError) {
	var f field.PosReqResult
	err := m.Body.Get(&f)
	return f, err
}

//PosReqStatus is a required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PosReqStatus() (field.PosReqStatus, errors.MessageRejectError) {
	var f field.PosReqStatus
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoPartyIDs() (field.NoPartyIDs, errors.MessageRejectError) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Account() (field.Account, errors.MessageRejectError) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AcctIDSource is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) AcctIDSource() (field.AcctIDSource, errors.MessageRejectError) {
	var f field.AcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) AccountType() (field.AccountType, errors.MessageRejectError) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Symbol() (field.Symbol, errors.MessageRejectError) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SymbolSfx() (field.SymbolSfx, errors.MessageRejectError) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityID() (field.SecurityID, errors.MessageRejectError) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityIDSource() (field.SecurityIDSource, errors.MessageRejectError) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoSecurityAltID() (field.NoSecurityAltID, errors.MessageRejectError) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Product() (field.Product, errors.MessageRejectError) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CFICode() (field.CFICode, errors.MessageRejectError) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityType() (field.SecurityType, errors.MessageRejectError) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecuritySubType() (field.SecuritySubType, errors.MessageRejectError) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) MaturityMonthYear() (field.MaturityMonthYear, errors.MessageRejectError) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) MaturityDate() (field.MaturityDate, errors.MessageRejectError) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CouponPaymentDate() (field.CouponPaymentDate, errors.MessageRejectError) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) IssueDate() (field.IssueDate, errors.MessageRejectError) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, errors.MessageRejectError) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) RepurchaseTerm() (field.RepurchaseTerm, errors.MessageRejectError) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) RepurchaseRate() (field.RepurchaseRate, errors.MessageRejectError) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Factor() (field.Factor, errors.MessageRejectError) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CreditRating() (field.CreditRating, errors.MessageRejectError) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) InstrRegistry() (field.InstrRegistry, errors.MessageRejectError) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CountryOfIssue() (field.CountryOfIssue, errors.MessageRejectError) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) LocaleOfIssue() (field.LocaleOfIssue, errors.MessageRejectError) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) RedemptionDate() (field.RedemptionDate, errors.MessageRejectError) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) StrikePrice() (field.StrikePrice, errors.MessageRejectError) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) StrikeCurrency() (field.StrikeCurrency, errors.MessageRejectError) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) OptAttribute() (field.OptAttribute, errors.MessageRejectError) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ContractMultiplier() (field.ContractMultiplier, errors.MessageRejectError) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CouponRate() (field.CouponRate, errors.MessageRejectError) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityExchange() (field.SecurityExchange, errors.MessageRejectError) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Issuer() (field.Issuer, errors.MessageRejectError) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedIssuerLen() (field.EncodedIssuerLen, errors.MessageRejectError) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedIssuer() (field.EncodedIssuer, errors.MessageRejectError) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityDesc() (field.SecurityDesc, errors.MessageRejectError) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, errors.MessageRejectError) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedSecurityDesc() (field.EncodedSecurityDesc, errors.MessageRejectError) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Pool() (field.Pool, errors.MessageRejectError) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ContractSettlMonth() (field.ContractSettlMonth, errors.MessageRejectError) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CPProgram() (field.CPProgram, errors.MessageRejectError) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CPRegType() (field.CPRegType, errors.MessageRejectError) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoEvents() (field.NoEvents, errors.MessageRejectError) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) DatedDate() (field.DatedDate, errors.MessageRejectError) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) InterestAccrualDate() (field.InterestAccrualDate, errors.MessageRejectError) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityStatus() (field.SecurityStatus, errors.MessageRejectError) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SettleOnOpenFlag() (field.SettleOnOpenFlag, errors.MessageRejectError) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) StrikeMultiplier() (field.StrikeMultiplier, errors.MessageRejectError) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) StrikeValue() (field.StrikeValue, errors.MessageRejectError) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) MinPriceIncrement() (field.MinPriceIncrement, errors.MessageRejectError) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PositionLimit() (field.PositionLimit, errors.MessageRejectError) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NTPositionLimit() (field.NTPositionLimit, errors.MessageRejectError) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoInstrumentParties() (field.NoInstrumentParties, errors.MessageRejectError) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) UnitOfMeasure() (field.UnitOfMeasure, errors.MessageRejectError) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) TimeUnit() (field.TimeUnit, errors.MessageRejectError) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) MaturityTime() (field.MaturityTime, errors.MessageRejectError) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//SecurityGroup is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityGroup() (field.SecurityGroup, errors.MessageRejectError) {
	var f field.SecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) MinPriceIncrementAmount() (field.MinPriceIncrementAmount, errors.MessageRejectError) {
	var f field.MinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) UnitOfMeasureQty() (field.UnitOfMeasureQty, errors.MessageRejectError) {
	var f field.UnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLLen is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityXMLLen() (field.SecurityXMLLen, errors.MessageRejectError) {
	var f field.SecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXML is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityXML() (field.SecurityXML, errors.MessageRejectError) {
	var f field.SecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLSchema is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityXMLSchema() (field.SecurityXMLSchema, errors.MessageRejectError) {
	var f field.SecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//ProductComplex is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ProductComplex() (field.ProductComplex, errors.MessageRejectError) {
	var f field.ProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PriceUnitOfMeasure() (field.PriceUnitOfMeasure, errors.MessageRejectError) {
	var f field.PriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PriceUnitOfMeasureQty() (field.PriceUnitOfMeasureQty, errors.MessageRejectError) {
	var f field.PriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SettlMethod is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SettlMethod() (field.SettlMethod, errors.MessageRejectError) {
	var f field.SettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//ExerciseStyle is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ExerciseStyle() (field.ExerciseStyle, errors.MessageRejectError) {
	var f field.ExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//OptPayAmount is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) OptPayAmount() (field.OptPayAmount, errors.MessageRejectError) {
	var f field.OptPayAmount
	err := m.Body.Get(&f)
	return f, err
}

//PriceQuoteMethod is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PriceQuoteMethod() (field.PriceQuoteMethod, errors.MessageRejectError) {
	var f field.PriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//ListMethod is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ListMethod() (field.ListMethod, errors.MessageRejectError) {
	var f field.ListMethod
	err := m.Body.Get(&f)
	return f, err
}

//CapPrice is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CapPrice() (field.CapPrice, errors.MessageRejectError) {
	var f field.CapPrice
	err := m.Body.Get(&f)
	return f, err
}

//FloorPrice is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) FloorPrice() (field.FloorPrice, errors.MessageRejectError) {
	var f field.FloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PutOrCall() (field.PutOrCall, errors.MessageRejectError) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//FlexibleIndicator is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) FlexibleIndicator() (field.FlexibleIndicator, errors.MessageRejectError) {
	var f field.FlexibleIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) FlexProductEligibilityIndicator() (field.FlexProductEligibilityIndicator, errors.MessageRejectError) {
	var f field.FlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FuturesValuationMethod is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) FuturesValuationMethod() (field.FuturesValuationMethod, errors.MessageRejectError) {
	var f field.FuturesValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Currency() (field.Currency, errors.MessageRejectError) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoLegs() (field.NoLegs, errors.MessageRejectError) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoUnderlyings() (field.NoUnderlyings, errors.MessageRejectError) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//ResponseTransportType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ResponseTransportType() (field.ResponseTransportType, errors.MessageRejectError) {
	var f field.ResponseTransportType
	err := m.Body.Get(&f)
	return f, err
}

//ResponseDestination is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ResponseDestination() (field.ResponseDestination, errors.MessageRejectError) {
	var f field.ResponseDestination
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//PosReqType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PosReqType() (field.PosReqType, errors.MessageRejectError) {
	var f field.PosReqType
	err := m.Body.Get(&f)
	return f, err
}

//MatchStatus is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) MatchStatus() (field.MatchStatus, errors.MessageRejectError) {
	var f field.MatchStatus
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ClearingBusinessDate() (field.ClearingBusinessDate, errors.MessageRejectError) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SubscriptionRequestType() (field.SubscriptionRequestType, errors.MessageRejectError) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessID is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SettlSessID() (field.SettlSessID, errors.MessageRejectError) {
	var f field.SettlSessID
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessSubID is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SettlSessSubID() (field.SettlSessSubID, errors.MessageRejectError) {
	var f field.SettlSessSubID
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrency is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SettlCurrency() (field.SettlCurrency, errors.MessageRejectError) {
	var f field.SettlCurrency
	err := m.Body.Get(&f)
	return f, err
}
