package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SecurityListRequest msg type = x.
type SecurityListRequest struct {
	message.Message
}

//SecurityListRequestBuilder builds SecurityListRequest messages.
type SecurityListRequestBuilder struct {
	message.MessageBuilder
}

//CreateSecurityListRequestBuilder returns an initialized SecurityListRequestBuilder with specified required fields.
func CreateSecurityListRequestBuilder(
	securityreqid field.SecurityReqID,
	securitylistrequesttype field.SecurityListRequestType) SecurityListRequestBuilder {
	var builder SecurityListRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(securityreqid)
	builder.Body.Set(securitylistrequesttype)
	return builder
}

//SecurityReqID is a required field for SecurityListRequest.
func (m SecurityListRequest) SecurityReqID() (field.SecurityReqID, errors.MessageRejectError) {
	var f field.SecurityReqID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityListRequestType is a required field for SecurityListRequest.
func (m SecurityListRequest) SecurityListRequestType() (field.SecurityListRequestType, errors.MessageRejectError) {
	var f field.SecurityListRequestType
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for SecurityListRequest.
func (m SecurityListRequest) Symbol() (field.Symbol, errors.MessageRejectError) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for SecurityListRequest.
func (m SecurityListRequest) SymbolSfx() (field.SymbolSfx, errors.MessageRejectError) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for SecurityListRequest.
func (m SecurityListRequest) SecurityID() (field.SecurityID, errors.MessageRejectError) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for SecurityListRequest.
func (m SecurityListRequest) SecurityIDSource() (field.SecurityIDSource, errors.MessageRejectError) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for SecurityListRequest.
func (m SecurityListRequest) NoSecurityAltID() (field.NoSecurityAltID, errors.MessageRejectError) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for SecurityListRequest.
func (m SecurityListRequest) Product() (field.Product, errors.MessageRejectError) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for SecurityListRequest.
func (m SecurityListRequest) CFICode() (field.CFICode, errors.MessageRejectError) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for SecurityListRequest.
func (m SecurityListRequest) SecurityType() (field.SecurityType, errors.MessageRejectError) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for SecurityListRequest.
func (m SecurityListRequest) SecuritySubType() (field.SecuritySubType, errors.MessageRejectError) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for SecurityListRequest.
func (m SecurityListRequest) MaturityMonthYear() (field.MaturityMonthYear, errors.MessageRejectError) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for SecurityListRequest.
func (m SecurityListRequest) MaturityDate() (field.MaturityDate, errors.MessageRejectError) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for SecurityListRequest.
func (m SecurityListRequest) CouponPaymentDate() (field.CouponPaymentDate, errors.MessageRejectError) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for SecurityListRequest.
func (m SecurityListRequest) IssueDate() (field.IssueDate, errors.MessageRejectError) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for SecurityListRequest.
func (m SecurityListRequest) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, errors.MessageRejectError) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for SecurityListRequest.
func (m SecurityListRequest) RepurchaseTerm() (field.RepurchaseTerm, errors.MessageRejectError) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for SecurityListRequest.
func (m SecurityListRequest) RepurchaseRate() (field.RepurchaseRate, errors.MessageRejectError) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for SecurityListRequest.
func (m SecurityListRequest) Factor() (field.Factor, errors.MessageRejectError) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for SecurityListRequest.
func (m SecurityListRequest) CreditRating() (field.CreditRating, errors.MessageRejectError) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for SecurityListRequest.
func (m SecurityListRequest) InstrRegistry() (field.InstrRegistry, errors.MessageRejectError) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for SecurityListRequest.
func (m SecurityListRequest) CountryOfIssue() (field.CountryOfIssue, errors.MessageRejectError) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for SecurityListRequest.
func (m SecurityListRequest) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for SecurityListRequest.
func (m SecurityListRequest) LocaleOfIssue() (field.LocaleOfIssue, errors.MessageRejectError) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for SecurityListRequest.
func (m SecurityListRequest) RedemptionDate() (field.RedemptionDate, errors.MessageRejectError) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for SecurityListRequest.
func (m SecurityListRequest) StrikePrice() (field.StrikePrice, errors.MessageRejectError) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for SecurityListRequest.
func (m SecurityListRequest) StrikeCurrency() (field.StrikeCurrency, errors.MessageRejectError) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for SecurityListRequest.
func (m SecurityListRequest) OptAttribute() (field.OptAttribute, errors.MessageRejectError) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for SecurityListRequest.
func (m SecurityListRequest) ContractMultiplier() (field.ContractMultiplier, errors.MessageRejectError) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for SecurityListRequest.
func (m SecurityListRequest) CouponRate() (field.CouponRate, errors.MessageRejectError) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for SecurityListRequest.
func (m SecurityListRequest) SecurityExchange() (field.SecurityExchange, errors.MessageRejectError) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for SecurityListRequest.
func (m SecurityListRequest) Issuer() (field.Issuer, errors.MessageRejectError) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for SecurityListRequest.
func (m SecurityListRequest) EncodedIssuerLen() (field.EncodedIssuerLen, errors.MessageRejectError) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for SecurityListRequest.
func (m SecurityListRequest) EncodedIssuer() (field.EncodedIssuer, errors.MessageRejectError) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for SecurityListRequest.
func (m SecurityListRequest) SecurityDesc() (field.SecurityDesc, errors.MessageRejectError) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for SecurityListRequest.
func (m SecurityListRequest) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, errors.MessageRejectError) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for SecurityListRequest.
func (m SecurityListRequest) EncodedSecurityDesc() (field.EncodedSecurityDesc, errors.MessageRejectError) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for SecurityListRequest.
func (m SecurityListRequest) Pool() (field.Pool, errors.MessageRejectError) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for SecurityListRequest.
func (m SecurityListRequest) ContractSettlMonth() (field.ContractSettlMonth, errors.MessageRejectError) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for SecurityListRequest.
func (m SecurityListRequest) CPProgram() (field.CPProgram, errors.MessageRejectError) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for SecurityListRequest.
func (m SecurityListRequest) CPRegType() (field.CPRegType, errors.MessageRejectError) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for SecurityListRequest.
func (m SecurityListRequest) NoEvents() (field.NoEvents, errors.MessageRejectError) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for SecurityListRequest.
func (m SecurityListRequest) DatedDate() (field.DatedDate, errors.MessageRejectError) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for SecurityListRequest.
func (m SecurityListRequest) InterestAccrualDate() (field.InterestAccrualDate, errors.MessageRejectError) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for SecurityListRequest.
func (m SecurityListRequest) SecurityStatus() (field.SecurityStatus, errors.MessageRejectError) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for SecurityListRequest.
func (m SecurityListRequest) SettleOnOpenFlag() (field.SettleOnOpenFlag, errors.MessageRejectError) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for SecurityListRequest.
func (m SecurityListRequest) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for SecurityListRequest.
func (m SecurityListRequest) StrikeMultiplier() (field.StrikeMultiplier, errors.MessageRejectError) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for SecurityListRequest.
func (m SecurityListRequest) StrikeValue() (field.StrikeValue, errors.MessageRejectError) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for SecurityListRequest.
func (m SecurityListRequest) MinPriceIncrement() (field.MinPriceIncrement, errors.MessageRejectError) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for SecurityListRequest.
func (m SecurityListRequest) PositionLimit() (field.PositionLimit, errors.MessageRejectError) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for SecurityListRequest.
func (m SecurityListRequest) NTPositionLimit() (field.NTPositionLimit, errors.MessageRejectError) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for SecurityListRequest.
func (m SecurityListRequest) NoInstrumentParties() (field.NoInstrumentParties, errors.MessageRejectError) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for SecurityListRequest.
func (m SecurityListRequest) UnitOfMeasure() (field.UnitOfMeasure, errors.MessageRejectError) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for SecurityListRequest.
func (m SecurityListRequest) TimeUnit() (field.TimeUnit, errors.MessageRejectError) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for SecurityListRequest.
func (m SecurityListRequest) MaturityTime() (field.MaturityTime, errors.MessageRejectError) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//SecurityGroup is a non-required field for SecurityListRequest.
func (m SecurityListRequest) SecurityGroup() (field.SecurityGroup, errors.MessageRejectError) {
	var f field.SecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for SecurityListRequest.
func (m SecurityListRequest) MinPriceIncrementAmount() (field.MinPriceIncrementAmount, errors.MessageRejectError) {
	var f field.MinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for SecurityListRequest.
func (m SecurityListRequest) UnitOfMeasureQty() (field.UnitOfMeasureQty, errors.MessageRejectError) {
	var f field.UnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLLen is a non-required field for SecurityListRequest.
func (m SecurityListRequest) SecurityXMLLen() (field.SecurityXMLLen, errors.MessageRejectError) {
	var f field.SecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXML is a non-required field for SecurityListRequest.
func (m SecurityListRequest) SecurityXML() (field.SecurityXML, errors.MessageRejectError) {
	var f field.SecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLSchema is a non-required field for SecurityListRequest.
func (m SecurityListRequest) SecurityXMLSchema() (field.SecurityXMLSchema, errors.MessageRejectError) {
	var f field.SecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//ProductComplex is a non-required field for SecurityListRequest.
func (m SecurityListRequest) ProductComplex() (field.ProductComplex, errors.MessageRejectError) {
	var f field.ProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for SecurityListRequest.
func (m SecurityListRequest) PriceUnitOfMeasure() (field.PriceUnitOfMeasure, errors.MessageRejectError) {
	var f field.PriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for SecurityListRequest.
func (m SecurityListRequest) PriceUnitOfMeasureQty() (field.PriceUnitOfMeasureQty, errors.MessageRejectError) {
	var f field.PriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SettlMethod is a non-required field for SecurityListRequest.
func (m SecurityListRequest) SettlMethod() (field.SettlMethod, errors.MessageRejectError) {
	var f field.SettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//ExerciseStyle is a non-required field for SecurityListRequest.
func (m SecurityListRequest) ExerciseStyle() (field.ExerciseStyle, errors.MessageRejectError) {
	var f field.ExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutAmount is a non-required field for SecurityListRequest.
func (m SecurityListRequest) OptPayoutAmount() (field.OptPayoutAmount, errors.MessageRejectError) {
	var f field.OptPayoutAmount
	err := m.Body.Get(&f)
	return f, err
}

//PriceQuoteMethod is a non-required field for SecurityListRequest.
func (m SecurityListRequest) PriceQuoteMethod() (field.PriceQuoteMethod, errors.MessageRejectError) {
	var f field.PriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//ListMethod is a non-required field for SecurityListRequest.
func (m SecurityListRequest) ListMethod() (field.ListMethod, errors.MessageRejectError) {
	var f field.ListMethod
	err := m.Body.Get(&f)
	return f, err
}

//CapPrice is a non-required field for SecurityListRequest.
func (m SecurityListRequest) CapPrice() (field.CapPrice, errors.MessageRejectError) {
	var f field.CapPrice
	err := m.Body.Get(&f)
	return f, err
}

//FloorPrice is a non-required field for SecurityListRequest.
func (m SecurityListRequest) FloorPrice() (field.FloorPrice, errors.MessageRejectError) {
	var f field.FloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for SecurityListRequest.
func (m SecurityListRequest) PutOrCall() (field.PutOrCall, errors.MessageRejectError) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//FlexibleIndicator is a non-required field for SecurityListRequest.
func (m SecurityListRequest) FlexibleIndicator() (field.FlexibleIndicator, errors.MessageRejectError) {
	var f field.FlexibleIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for SecurityListRequest.
func (m SecurityListRequest) FlexProductEligibilityIndicator() (field.FlexProductEligibilityIndicator, errors.MessageRejectError) {
	var f field.FlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//ValuationMethod is a non-required field for SecurityListRequest.
func (m SecurityListRequest) ValuationMethod() (field.ValuationMethod, errors.MessageRejectError) {
	var f field.ValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplierUnit is a non-required field for SecurityListRequest.
func (m SecurityListRequest) ContractMultiplierUnit() (field.ContractMultiplierUnit, errors.MessageRejectError) {
	var f field.ContractMultiplierUnit
	err := m.Body.Get(&f)
	return f, err
}

//FlowScheduleType is a non-required field for SecurityListRequest.
func (m SecurityListRequest) FlowScheduleType() (field.FlowScheduleType, errors.MessageRejectError) {
	var f field.FlowScheduleType
	err := m.Body.Get(&f)
	return f, err
}

//RestructuringType is a non-required field for SecurityListRequest.
func (m SecurityListRequest) RestructuringType() (field.RestructuringType, errors.MessageRejectError) {
	var f field.RestructuringType
	err := m.Body.Get(&f)
	return f, err
}

//Seniority is a non-required field for SecurityListRequest.
func (m SecurityListRequest) Seniority() (field.Seniority, errors.MessageRejectError) {
	var f field.Seniority
	err := m.Body.Get(&f)
	return f, err
}

//NotionalPercentageOutstanding is a non-required field for SecurityListRequest.
func (m SecurityListRequest) NotionalPercentageOutstanding() (field.NotionalPercentageOutstanding, errors.MessageRejectError) {
	var f field.NotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//OriginalNotionalPercentageOutstanding is a non-required field for SecurityListRequest.
func (m SecurityListRequest) OriginalNotionalPercentageOutstanding() (field.OriginalNotionalPercentageOutstanding, errors.MessageRejectError) {
	var f field.OriginalNotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//AttachmentPoint is a non-required field for SecurityListRequest.
func (m SecurityListRequest) AttachmentPoint() (field.AttachmentPoint, errors.MessageRejectError) {
	var f field.AttachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//DetachmentPoint is a non-required field for SecurityListRequest.
func (m SecurityListRequest) DetachmentPoint() (field.DetachmentPoint, errors.MessageRejectError) {
	var f field.DetachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceDeterminationMethod is a non-required field for SecurityListRequest.
func (m SecurityListRequest) StrikePriceDeterminationMethod() (field.StrikePriceDeterminationMethod, errors.MessageRejectError) {
	var f field.StrikePriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryMethod is a non-required field for SecurityListRequest.
func (m SecurityListRequest) StrikePriceBoundaryMethod() (field.StrikePriceBoundaryMethod, errors.MessageRejectError) {
	var f field.StrikePriceBoundaryMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryPrecision is a non-required field for SecurityListRequest.
func (m SecurityListRequest) StrikePriceBoundaryPrecision() (field.StrikePriceBoundaryPrecision, errors.MessageRejectError) {
	var f field.StrikePriceBoundaryPrecision
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceDeterminationMethod is a non-required field for SecurityListRequest.
func (m SecurityListRequest) UnderlyingPriceDeterminationMethod() (field.UnderlyingPriceDeterminationMethod, errors.MessageRejectError) {
	var f field.UnderlyingPriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutType is a non-required field for SecurityListRequest.
func (m SecurityListRequest) OptPayoutType() (field.OptPayoutType, errors.MessageRejectError) {
	var f field.OptPayoutType
	err := m.Body.Get(&f)
	return f, err
}

//NoComplexEvents is a non-required field for SecurityListRequest.
func (m SecurityListRequest) NoComplexEvents() (field.NoComplexEvents, errors.MessageRejectError) {
	var f field.NoComplexEvents
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryForm is a non-required field for SecurityListRequest.
func (m SecurityListRequest) DeliveryForm() (field.DeliveryForm, errors.MessageRejectError) {
	var f field.DeliveryForm
	err := m.Body.Get(&f)
	return f, err
}

//PctAtRisk is a non-required field for SecurityListRequest.
func (m SecurityListRequest) PctAtRisk() (field.PctAtRisk, errors.MessageRejectError) {
	var f field.PctAtRisk
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrAttrib is a non-required field for SecurityListRequest.
func (m SecurityListRequest) NoInstrAttrib() (field.NoInstrAttrib, errors.MessageRejectError) {
	var f field.NoInstrAttrib
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDesc is a non-required field for SecurityListRequest.
func (m SecurityListRequest) AgreementDesc() (field.AgreementDesc, errors.MessageRejectError) {
	var f field.AgreementDesc
	err := m.Body.Get(&f)
	return f, err
}

//AgreementID is a non-required field for SecurityListRequest.
func (m SecurityListRequest) AgreementID() (field.AgreementID, errors.MessageRejectError) {
	var f field.AgreementID
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDate is a non-required field for SecurityListRequest.
func (m SecurityListRequest) AgreementDate() (field.AgreementDate, errors.MessageRejectError) {
	var f field.AgreementDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementCurrency is a non-required field for SecurityListRequest.
func (m SecurityListRequest) AgreementCurrency() (field.AgreementCurrency, errors.MessageRejectError) {
	var f field.AgreementCurrency
	err := m.Body.Get(&f)
	return f, err
}

//TerminationType is a non-required field for SecurityListRequest.
func (m SecurityListRequest) TerminationType() (field.TerminationType, errors.MessageRejectError) {
	var f field.TerminationType
	err := m.Body.Get(&f)
	return f, err
}

//StartDate is a non-required field for SecurityListRequest.
func (m SecurityListRequest) StartDate() (field.StartDate, errors.MessageRejectError) {
	var f field.StartDate
	err := m.Body.Get(&f)
	return f, err
}

//EndDate is a non-required field for SecurityListRequest.
func (m SecurityListRequest) EndDate() (field.EndDate, errors.MessageRejectError) {
	var f field.EndDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryType is a non-required field for SecurityListRequest.
func (m SecurityListRequest) DeliveryType() (field.DeliveryType, errors.MessageRejectError) {
	var f field.DeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//MarginRatio is a non-required field for SecurityListRequest.
func (m SecurityListRequest) MarginRatio() (field.MarginRatio, errors.MessageRejectError) {
	var f field.MarginRatio
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for SecurityListRequest.
func (m SecurityListRequest) NoUnderlyings() (field.NoUnderlyings, errors.MessageRejectError) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for SecurityListRequest.
func (m SecurityListRequest) NoLegs() (field.NoLegs, errors.MessageRejectError) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for SecurityListRequest.
func (m SecurityListRequest) Currency() (field.Currency, errors.MessageRejectError) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for SecurityListRequest.
func (m SecurityListRequest) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for SecurityListRequest.
func (m SecurityListRequest) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for SecurityListRequest.
func (m SecurityListRequest) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for SecurityListRequest.
func (m SecurityListRequest) TradingSessionID() (field.TradingSessionID, errors.MessageRejectError) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for SecurityListRequest.
func (m SecurityListRequest) TradingSessionSubID() (field.TradingSessionSubID, errors.MessageRejectError) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a non-required field for SecurityListRequest.
func (m SecurityListRequest) SubscriptionRequestType() (field.SubscriptionRequestType, errors.MessageRejectError) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}

//MarketID is a non-required field for SecurityListRequest.
func (m SecurityListRequest) MarketID() (field.MarketID, errors.MessageRejectError) {
	var f field.MarketID
	err := m.Body.Get(&f)
	return f, err
}

//MarketSegmentID is a non-required field for SecurityListRequest.
func (m SecurityListRequest) MarketSegmentID() (field.MarketSegmentID, errors.MessageRejectError) {
	var f field.MarketSegmentID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityListID is a non-required field for SecurityListRequest.
func (m SecurityListRequest) SecurityListID() (field.SecurityListID, errors.MessageRejectError) {
	var f field.SecurityListID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityListType is a non-required field for SecurityListRequest.
func (m SecurityListRequest) SecurityListType() (field.SecurityListType, errors.MessageRejectError) {
	var f field.SecurityListType
	err := m.Body.Get(&f)
	return f, err
}

//SecurityListTypeSource is a non-required field for SecurityListRequest.
func (m SecurityListRequest) SecurityListTypeSource() (field.SecurityListTypeSource, errors.MessageRejectError) {
	var f field.SecurityListTypeSource
	err := m.Body.Get(&f)
	return f, err
}
