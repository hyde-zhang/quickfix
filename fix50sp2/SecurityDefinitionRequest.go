package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SecurityDefinitionRequest msg type = c.
type SecurityDefinitionRequest struct {
	message.Message
}

//SecurityDefinitionRequestBuilder builds SecurityDefinitionRequest messages.
type SecurityDefinitionRequestBuilder struct {
	message.MessageBuilder
}

//CreateSecurityDefinitionRequestBuilder returns an initialized SecurityDefinitionRequestBuilder with specified required fields.
func CreateSecurityDefinitionRequestBuilder(
	securityreqid field.SecurityReqID,
	securityrequesttype field.SecurityRequestType) SecurityDefinitionRequestBuilder {
	var builder SecurityDefinitionRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(securityreqid)
	builder.Body.Set(securityrequesttype)
	return builder
}

//SecurityReqID is a required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityReqID() (field.SecurityReqID, errors.MessageRejectError) {
	var f field.SecurityReqID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityRequestType is a required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityRequestType() (field.SecurityRequestType, errors.MessageRejectError) {
	var f field.SecurityRequestType
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Symbol() (field.Symbol, errors.MessageRejectError) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SymbolSfx() (field.SymbolSfx, errors.MessageRejectError) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityID() (field.SecurityID, errors.MessageRejectError) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityIDSource() (field.SecurityIDSource, errors.MessageRejectError) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) NoSecurityAltID() (field.NoSecurityAltID, errors.MessageRejectError) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Product() (field.Product, errors.MessageRejectError) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) CFICode() (field.CFICode, errors.MessageRejectError) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityType() (field.SecurityType, errors.MessageRejectError) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecuritySubType() (field.SecuritySubType, errors.MessageRejectError) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) MaturityMonthYear() (field.MaturityMonthYear, errors.MessageRejectError) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) MaturityDate() (field.MaturityDate, errors.MessageRejectError) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) CouponPaymentDate() (field.CouponPaymentDate, errors.MessageRejectError) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) IssueDate() (field.IssueDate, errors.MessageRejectError) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, errors.MessageRejectError) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) RepurchaseTerm() (field.RepurchaseTerm, errors.MessageRejectError) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) RepurchaseRate() (field.RepurchaseRate, errors.MessageRejectError) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Factor() (field.Factor, errors.MessageRejectError) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) CreditRating() (field.CreditRating, errors.MessageRejectError) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) InstrRegistry() (field.InstrRegistry, errors.MessageRejectError) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) CountryOfIssue() (field.CountryOfIssue, errors.MessageRejectError) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) LocaleOfIssue() (field.LocaleOfIssue, errors.MessageRejectError) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) RedemptionDate() (field.RedemptionDate, errors.MessageRejectError) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) StrikePrice() (field.StrikePrice, errors.MessageRejectError) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) StrikeCurrency() (field.StrikeCurrency, errors.MessageRejectError) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) OptAttribute() (field.OptAttribute, errors.MessageRejectError) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) ContractMultiplier() (field.ContractMultiplier, errors.MessageRejectError) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) CouponRate() (field.CouponRate, errors.MessageRejectError) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityExchange() (field.SecurityExchange, errors.MessageRejectError) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Issuer() (field.Issuer, errors.MessageRejectError) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) EncodedIssuerLen() (field.EncodedIssuerLen, errors.MessageRejectError) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) EncodedIssuer() (field.EncodedIssuer, errors.MessageRejectError) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityDesc() (field.SecurityDesc, errors.MessageRejectError) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, errors.MessageRejectError) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) EncodedSecurityDesc() (field.EncodedSecurityDesc, errors.MessageRejectError) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Pool() (field.Pool, errors.MessageRejectError) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) ContractSettlMonth() (field.ContractSettlMonth, errors.MessageRejectError) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) CPProgram() (field.CPProgram, errors.MessageRejectError) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) CPRegType() (field.CPRegType, errors.MessageRejectError) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) NoEvents() (field.NoEvents, errors.MessageRejectError) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) DatedDate() (field.DatedDate, errors.MessageRejectError) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) InterestAccrualDate() (field.InterestAccrualDate, errors.MessageRejectError) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityStatus() (field.SecurityStatus, errors.MessageRejectError) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SettleOnOpenFlag() (field.SettleOnOpenFlag, errors.MessageRejectError) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) StrikeMultiplier() (field.StrikeMultiplier, errors.MessageRejectError) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) StrikeValue() (field.StrikeValue, errors.MessageRejectError) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) MinPriceIncrement() (field.MinPriceIncrement, errors.MessageRejectError) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) PositionLimit() (field.PositionLimit, errors.MessageRejectError) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) NTPositionLimit() (field.NTPositionLimit, errors.MessageRejectError) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) NoInstrumentParties() (field.NoInstrumentParties, errors.MessageRejectError) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) UnitOfMeasure() (field.UnitOfMeasure, errors.MessageRejectError) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) TimeUnit() (field.TimeUnit, errors.MessageRejectError) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) MaturityTime() (field.MaturityTime, errors.MessageRejectError) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//SecurityGroup is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityGroup() (field.SecurityGroup, errors.MessageRejectError) {
	var f field.SecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) MinPriceIncrementAmount() (field.MinPriceIncrementAmount, errors.MessageRejectError) {
	var f field.MinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) UnitOfMeasureQty() (field.UnitOfMeasureQty, errors.MessageRejectError) {
	var f field.UnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLLen is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityXMLLen() (field.SecurityXMLLen, errors.MessageRejectError) {
	var f field.SecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXML is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityXML() (field.SecurityXML, errors.MessageRejectError) {
	var f field.SecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLSchema is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityXMLSchema() (field.SecurityXMLSchema, errors.MessageRejectError) {
	var f field.SecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//ProductComplex is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) ProductComplex() (field.ProductComplex, errors.MessageRejectError) {
	var f field.ProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) PriceUnitOfMeasure() (field.PriceUnitOfMeasure, errors.MessageRejectError) {
	var f field.PriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) PriceUnitOfMeasureQty() (field.PriceUnitOfMeasureQty, errors.MessageRejectError) {
	var f field.PriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SettlMethod is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SettlMethod() (field.SettlMethod, errors.MessageRejectError) {
	var f field.SettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//ExerciseStyle is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) ExerciseStyle() (field.ExerciseStyle, errors.MessageRejectError) {
	var f field.ExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutAmount is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) OptPayoutAmount() (field.OptPayoutAmount, errors.MessageRejectError) {
	var f field.OptPayoutAmount
	err := m.Body.Get(&f)
	return f, err
}

//PriceQuoteMethod is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) PriceQuoteMethod() (field.PriceQuoteMethod, errors.MessageRejectError) {
	var f field.PriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//ListMethod is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) ListMethod() (field.ListMethod, errors.MessageRejectError) {
	var f field.ListMethod
	err := m.Body.Get(&f)
	return f, err
}

//CapPrice is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) CapPrice() (field.CapPrice, errors.MessageRejectError) {
	var f field.CapPrice
	err := m.Body.Get(&f)
	return f, err
}

//FloorPrice is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) FloorPrice() (field.FloorPrice, errors.MessageRejectError) {
	var f field.FloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) PutOrCall() (field.PutOrCall, errors.MessageRejectError) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//FlexibleIndicator is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) FlexibleIndicator() (field.FlexibleIndicator, errors.MessageRejectError) {
	var f field.FlexibleIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) FlexProductEligibilityIndicator() (field.FlexProductEligibilityIndicator, errors.MessageRejectError) {
	var f field.FlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//ValuationMethod is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) ValuationMethod() (field.ValuationMethod, errors.MessageRejectError) {
	var f field.ValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplierUnit is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) ContractMultiplierUnit() (field.ContractMultiplierUnit, errors.MessageRejectError) {
	var f field.ContractMultiplierUnit
	err := m.Body.Get(&f)
	return f, err
}

//FlowScheduleType is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) FlowScheduleType() (field.FlowScheduleType, errors.MessageRejectError) {
	var f field.FlowScheduleType
	err := m.Body.Get(&f)
	return f, err
}

//RestructuringType is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) RestructuringType() (field.RestructuringType, errors.MessageRejectError) {
	var f field.RestructuringType
	err := m.Body.Get(&f)
	return f, err
}

//Seniority is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Seniority() (field.Seniority, errors.MessageRejectError) {
	var f field.Seniority
	err := m.Body.Get(&f)
	return f, err
}

//NotionalPercentageOutstanding is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) NotionalPercentageOutstanding() (field.NotionalPercentageOutstanding, errors.MessageRejectError) {
	var f field.NotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//OriginalNotionalPercentageOutstanding is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) OriginalNotionalPercentageOutstanding() (field.OriginalNotionalPercentageOutstanding, errors.MessageRejectError) {
	var f field.OriginalNotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//AttachmentPoint is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) AttachmentPoint() (field.AttachmentPoint, errors.MessageRejectError) {
	var f field.AttachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//DetachmentPoint is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) DetachmentPoint() (field.DetachmentPoint, errors.MessageRejectError) {
	var f field.DetachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceDeterminationMethod is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) StrikePriceDeterminationMethod() (field.StrikePriceDeterminationMethod, errors.MessageRejectError) {
	var f field.StrikePriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryMethod is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) StrikePriceBoundaryMethod() (field.StrikePriceBoundaryMethod, errors.MessageRejectError) {
	var f field.StrikePriceBoundaryMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryPrecision is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) StrikePriceBoundaryPrecision() (field.StrikePriceBoundaryPrecision, errors.MessageRejectError) {
	var f field.StrikePriceBoundaryPrecision
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceDeterminationMethod is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) UnderlyingPriceDeterminationMethod() (field.UnderlyingPriceDeterminationMethod, errors.MessageRejectError) {
	var f field.UnderlyingPriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutType is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) OptPayoutType() (field.OptPayoutType, errors.MessageRejectError) {
	var f field.OptPayoutType
	err := m.Body.Get(&f)
	return f, err
}

//NoComplexEvents is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) NoComplexEvents() (field.NoComplexEvents, errors.MessageRejectError) {
	var f field.NoComplexEvents
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryForm is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) DeliveryForm() (field.DeliveryForm, errors.MessageRejectError) {
	var f field.DeliveryForm
	err := m.Body.Get(&f)
	return f, err
}

//PctAtRisk is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) PctAtRisk() (field.PctAtRisk, errors.MessageRejectError) {
	var f field.PctAtRisk
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrAttrib is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) NoInstrAttrib() (field.NoInstrAttrib, errors.MessageRejectError) {
	var f field.NoInstrAttrib
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) NoUnderlyings() (field.NoUnderlyings, errors.MessageRejectError) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Currency() (field.Currency, errors.MessageRejectError) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) TradingSessionID() (field.TradingSessionID, errors.MessageRejectError) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) TradingSessionSubID() (field.TradingSessionSubID, errors.MessageRejectError) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) NoLegs() (field.NoLegs, errors.MessageRejectError) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//ExpirationCycle is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) ExpirationCycle() (field.ExpirationCycle, errors.MessageRejectError) {
	var f field.ExpirationCycle
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SubscriptionRequestType() (field.SubscriptionRequestType, errors.MessageRejectError) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}

//NoStipulations is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) NoStipulations() (field.NoStipulations, errors.MessageRejectError) {
	var f field.NoStipulations
	err := m.Body.Get(&f)
	return f, err
}

//Spread is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Spread() (field.Spread, errors.MessageRejectError) {
	var f field.Spread
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) BenchmarkCurveCurrency() (field.BenchmarkCurveCurrency, errors.MessageRejectError) {
	var f field.BenchmarkCurveCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveName is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) BenchmarkCurveName() (field.BenchmarkCurveName, errors.MessageRejectError) {
	var f field.BenchmarkCurveName
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) BenchmarkCurvePoint() (field.BenchmarkCurvePoint, errors.MessageRejectError) {
	var f field.BenchmarkCurvePoint
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPrice is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) BenchmarkPrice() (field.BenchmarkPrice, errors.MessageRejectError) {
	var f field.BenchmarkPrice
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPriceType is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) BenchmarkPriceType() (field.BenchmarkPriceType, errors.MessageRejectError) {
	var f field.BenchmarkPriceType
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) BenchmarkSecurityID() (field.BenchmarkSecurityID, errors.MessageRejectError) {
	var f field.BenchmarkSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) BenchmarkSecurityIDSource() (field.BenchmarkSecurityIDSource, errors.MessageRejectError) {
	var f field.BenchmarkSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//YieldType is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) YieldType() (field.YieldType, errors.MessageRejectError) {
	var f field.YieldType
	err := m.Body.Get(&f)
	return f, err
}

//Yield is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Yield() (field.Yield, errors.MessageRejectError) {
	var f field.Yield
	err := m.Body.Get(&f)
	return f, err
}

//YieldCalcDate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) YieldCalcDate() (field.YieldCalcDate, errors.MessageRejectError) {
	var f field.YieldCalcDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionDate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) YieldRedemptionDate() (field.YieldRedemptionDate, errors.MessageRejectError) {
	var f field.YieldRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) YieldRedemptionPrice() (field.YieldRedemptionPrice, errors.MessageRejectError) {
	var f field.YieldRedemptionPrice
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) YieldRedemptionPriceType() (field.YieldRedemptionPriceType, errors.MessageRejectError) {
	var f field.YieldRedemptionPriceType
	err := m.Body.Get(&f)
	return f, err
}

//MarketID is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) MarketID() (field.MarketID, errors.MessageRejectError) {
	var f field.MarketID
	err := m.Body.Get(&f)
	return f, err
}

//MarketSegmentID is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) MarketSegmentID() (field.MarketSegmentID, errors.MessageRejectError) {
	var f field.MarketSegmentID
	err := m.Body.Get(&f)
	return f, err
}
