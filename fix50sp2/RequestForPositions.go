package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//RequestForPositions msg type = AN.
type RequestForPositions struct {
	message.Message
}

//RequestForPositionsBuilder builds RequestForPositions messages.
type RequestForPositionsBuilder struct {
	message.MessageBuilder
}

//CreateRequestForPositionsBuilder returns an initialized RequestForPositionsBuilder with specified required fields.
func CreateRequestForPositionsBuilder(
	posreqid field.PosReqID,
	posreqtype field.PosReqType,
	clearingbusinessdate field.ClearingBusinessDate,
	transacttime field.TransactTime) RequestForPositionsBuilder {
	var builder RequestForPositionsBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(posreqid)
	builder.Body.Set(posreqtype)
	builder.Body.Set(clearingbusinessdate)
	builder.Body.Set(transacttime)
	return builder
}

//PosReqID is a required field for RequestForPositions.
func (m RequestForPositions) PosReqID() (field.PosReqID, errors.MessageRejectError) {
	var f field.PosReqID
	err := m.Body.Get(&f)
	return f, err
}

//PosReqType is a required field for RequestForPositions.
func (m RequestForPositions) PosReqType() (field.PosReqType, errors.MessageRejectError) {
	var f field.PosReqType
	err := m.Body.Get(&f)
	return f, err
}

//MatchStatus is a non-required field for RequestForPositions.
func (m RequestForPositions) MatchStatus() (field.MatchStatus, errors.MessageRejectError) {
	var f field.MatchStatus
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a non-required field for RequestForPositions.
func (m RequestForPositions) SubscriptionRequestType() (field.SubscriptionRequestType, errors.MessageRejectError) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for RequestForPositions.
func (m RequestForPositions) NoPartyIDs() (field.NoPartyIDs, errors.MessageRejectError) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for RequestForPositions.
func (m RequestForPositions) Account() (field.Account, errors.MessageRejectError) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AcctIDSource is a non-required field for RequestForPositions.
func (m RequestForPositions) AcctIDSource() (field.AcctIDSource, errors.MessageRejectError) {
	var f field.AcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for RequestForPositions.
func (m RequestForPositions) AccountType() (field.AccountType, errors.MessageRejectError) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for RequestForPositions.
func (m RequestForPositions) Symbol() (field.Symbol, errors.MessageRejectError) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for RequestForPositions.
func (m RequestForPositions) SymbolSfx() (field.SymbolSfx, errors.MessageRejectError) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for RequestForPositions.
func (m RequestForPositions) SecurityID() (field.SecurityID, errors.MessageRejectError) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for RequestForPositions.
func (m RequestForPositions) SecurityIDSource() (field.SecurityIDSource, errors.MessageRejectError) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for RequestForPositions.
func (m RequestForPositions) NoSecurityAltID() (field.NoSecurityAltID, errors.MessageRejectError) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for RequestForPositions.
func (m RequestForPositions) Product() (field.Product, errors.MessageRejectError) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for RequestForPositions.
func (m RequestForPositions) CFICode() (field.CFICode, errors.MessageRejectError) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for RequestForPositions.
func (m RequestForPositions) SecurityType() (field.SecurityType, errors.MessageRejectError) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for RequestForPositions.
func (m RequestForPositions) SecuritySubType() (field.SecuritySubType, errors.MessageRejectError) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for RequestForPositions.
func (m RequestForPositions) MaturityMonthYear() (field.MaturityMonthYear, errors.MessageRejectError) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for RequestForPositions.
func (m RequestForPositions) MaturityDate() (field.MaturityDate, errors.MessageRejectError) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for RequestForPositions.
func (m RequestForPositions) CouponPaymentDate() (field.CouponPaymentDate, errors.MessageRejectError) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for RequestForPositions.
func (m RequestForPositions) IssueDate() (field.IssueDate, errors.MessageRejectError) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for RequestForPositions.
func (m RequestForPositions) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, errors.MessageRejectError) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for RequestForPositions.
func (m RequestForPositions) RepurchaseTerm() (field.RepurchaseTerm, errors.MessageRejectError) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for RequestForPositions.
func (m RequestForPositions) RepurchaseRate() (field.RepurchaseRate, errors.MessageRejectError) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for RequestForPositions.
func (m RequestForPositions) Factor() (field.Factor, errors.MessageRejectError) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for RequestForPositions.
func (m RequestForPositions) CreditRating() (field.CreditRating, errors.MessageRejectError) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for RequestForPositions.
func (m RequestForPositions) InstrRegistry() (field.InstrRegistry, errors.MessageRejectError) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for RequestForPositions.
func (m RequestForPositions) CountryOfIssue() (field.CountryOfIssue, errors.MessageRejectError) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for RequestForPositions.
func (m RequestForPositions) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for RequestForPositions.
func (m RequestForPositions) LocaleOfIssue() (field.LocaleOfIssue, errors.MessageRejectError) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for RequestForPositions.
func (m RequestForPositions) RedemptionDate() (field.RedemptionDate, errors.MessageRejectError) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for RequestForPositions.
func (m RequestForPositions) StrikePrice() (field.StrikePrice, errors.MessageRejectError) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for RequestForPositions.
func (m RequestForPositions) StrikeCurrency() (field.StrikeCurrency, errors.MessageRejectError) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for RequestForPositions.
func (m RequestForPositions) OptAttribute() (field.OptAttribute, errors.MessageRejectError) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for RequestForPositions.
func (m RequestForPositions) ContractMultiplier() (field.ContractMultiplier, errors.MessageRejectError) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for RequestForPositions.
func (m RequestForPositions) CouponRate() (field.CouponRate, errors.MessageRejectError) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for RequestForPositions.
func (m RequestForPositions) SecurityExchange() (field.SecurityExchange, errors.MessageRejectError) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for RequestForPositions.
func (m RequestForPositions) Issuer() (field.Issuer, errors.MessageRejectError) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for RequestForPositions.
func (m RequestForPositions) EncodedIssuerLen() (field.EncodedIssuerLen, errors.MessageRejectError) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for RequestForPositions.
func (m RequestForPositions) EncodedIssuer() (field.EncodedIssuer, errors.MessageRejectError) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for RequestForPositions.
func (m RequestForPositions) SecurityDesc() (field.SecurityDesc, errors.MessageRejectError) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for RequestForPositions.
func (m RequestForPositions) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, errors.MessageRejectError) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for RequestForPositions.
func (m RequestForPositions) EncodedSecurityDesc() (field.EncodedSecurityDesc, errors.MessageRejectError) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for RequestForPositions.
func (m RequestForPositions) Pool() (field.Pool, errors.MessageRejectError) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for RequestForPositions.
func (m RequestForPositions) ContractSettlMonth() (field.ContractSettlMonth, errors.MessageRejectError) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for RequestForPositions.
func (m RequestForPositions) CPProgram() (field.CPProgram, errors.MessageRejectError) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for RequestForPositions.
func (m RequestForPositions) CPRegType() (field.CPRegType, errors.MessageRejectError) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for RequestForPositions.
func (m RequestForPositions) NoEvents() (field.NoEvents, errors.MessageRejectError) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for RequestForPositions.
func (m RequestForPositions) DatedDate() (field.DatedDate, errors.MessageRejectError) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for RequestForPositions.
func (m RequestForPositions) InterestAccrualDate() (field.InterestAccrualDate, errors.MessageRejectError) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for RequestForPositions.
func (m RequestForPositions) SecurityStatus() (field.SecurityStatus, errors.MessageRejectError) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for RequestForPositions.
func (m RequestForPositions) SettleOnOpenFlag() (field.SettleOnOpenFlag, errors.MessageRejectError) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for RequestForPositions.
func (m RequestForPositions) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for RequestForPositions.
func (m RequestForPositions) StrikeMultiplier() (field.StrikeMultiplier, errors.MessageRejectError) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for RequestForPositions.
func (m RequestForPositions) StrikeValue() (field.StrikeValue, errors.MessageRejectError) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for RequestForPositions.
func (m RequestForPositions) MinPriceIncrement() (field.MinPriceIncrement, errors.MessageRejectError) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for RequestForPositions.
func (m RequestForPositions) PositionLimit() (field.PositionLimit, errors.MessageRejectError) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for RequestForPositions.
func (m RequestForPositions) NTPositionLimit() (field.NTPositionLimit, errors.MessageRejectError) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for RequestForPositions.
func (m RequestForPositions) NoInstrumentParties() (field.NoInstrumentParties, errors.MessageRejectError) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for RequestForPositions.
func (m RequestForPositions) UnitOfMeasure() (field.UnitOfMeasure, errors.MessageRejectError) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for RequestForPositions.
func (m RequestForPositions) TimeUnit() (field.TimeUnit, errors.MessageRejectError) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for RequestForPositions.
func (m RequestForPositions) MaturityTime() (field.MaturityTime, errors.MessageRejectError) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//SecurityGroup is a non-required field for RequestForPositions.
func (m RequestForPositions) SecurityGroup() (field.SecurityGroup, errors.MessageRejectError) {
	var f field.SecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for RequestForPositions.
func (m RequestForPositions) MinPriceIncrementAmount() (field.MinPriceIncrementAmount, errors.MessageRejectError) {
	var f field.MinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for RequestForPositions.
func (m RequestForPositions) UnitOfMeasureQty() (field.UnitOfMeasureQty, errors.MessageRejectError) {
	var f field.UnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLLen is a non-required field for RequestForPositions.
func (m RequestForPositions) SecurityXMLLen() (field.SecurityXMLLen, errors.MessageRejectError) {
	var f field.SecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXML is a non-required field for RequestForPositions.
func (m RequestForPositions) SecurityXML() (field.SecurityXML, errors.MessageRejectError) {
	var f field.SecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLSchema is a non-required field for RequestForPositions.
func (m RequestForPositions) SecurityXMLSchema() (field.SecurityXMLSchema, errors.MessageRejectError) {
	var f field.SecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//ProductComplex is a non-required field for RequestForPositions.
func (m RequestForPositions) ProductComplex() (field.ProductComplex, errors.MessageRejectError) {
	var f field.ProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for RequestForPositions.
func (m RequestForPositions) PriceUnitOfMeasure() (field.PriceUnitOfMeasure, errors.MessageRejectError) {
	var f field.PriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for RequestForPositions.
func (m RequestForPositions) PriceUnitOfMeasureQty() (field.PriceUnitOfMeasureQty, errors.MessageRejectError) {
	var f field.PriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SettlMethod is a non-required field for RequestForPositions.
func (m RequestForPositions) SettlMethod() (field.SettlMethod, errors.MessageRejectError) {
	var f field.SettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//ExerciseStyle is a non-required field for RequestForPositions.
func (m RequestForPositions) ExerciseStyle() (field.ExerciseStyle, errors.MessageRejectError) {
	var f field.ExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutAmount is a non-required field for RequestForPositions.
func (m RequestForPositions) OptPayoutAmount() (field.OptPayoutAmount, errors.MessageRejectError) {
	var f field.OptPayoutAmount
	err := m.Body.Get(&f)
	return f, err
}

//PriceQuoteMethod is a non-required field for RequestForPositions.
func (m RequestForPositions) PriceQuoteMethod() (field.PriceQuoteMethod, errors.MessageRejectError) {
	var f field.PriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//ListMethod is a non-required field for RequestForPositions.
func (m RequestForPositions) ListMethod() (field.ListMethod, errors.MessageRejectError) {
	var f field.ListMethod
	err := m.Body.Get(&f)
	return f, err
}

//CapPrice is a non-required field for RequestForPositions.
func (m RequestForPositions) CapPrice() (field.CapPrice, errors.MessageRejectError) {
	var f field.CapPrice
	err := m.Body.Get(&f)
	return f, err
}

//FloorPrice is a non-required field for RequestForPositions.
func (m RequestForPositions) FloorPrice() (field.FloorPrice, errors.MessageRejectError) {
	var f field.FloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for RequestForPositions.
func (m RequestForPositions) PutOrCall() (field.PutOrCall, errors.MessageRejectError) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//FlexibleIndicator is a non-required field for RequestForPositions.
func (m RequestForPositions) FlexibleIndicator() (field.FlexibleIndicator, errors.MessageRejectError) {
	var f field.FlexibleIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for RequestForPositions.
func (m RequestForPositions) FlexProductEligibilityIndicator() (field.FlexProductEligibilityIndicator, errors.MessageRejectError) {
	var f field.FlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//ValuationMethod is a non-required field for RequestForPositions.
func (m RequestForPositions) ValuationMethod() (field.ValuationMethod, errors.MessageRejectError) {
	var f field.ValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplierUnit is a non-required field for RequestForPositions.
func (m RequestForPositions) ContractMultiplierUnit() (field.ContractMultiplierUnit, errors.MessageRejectError) {
	var f field.ContractMultiplierUnit
	err := m.Body.Get(&f)
	return f, err
}

//FlowScheduleType is a non-required field for RequestForPositions.
func (m RequestForPositions) FlowScheduleType() (field.FlowScheduleType, errors.MessageRejectError) {
	var f field.FlowScheduleType
	err := m.Body.Get(&f)
	return f, err
}

//RestructuringType is a non-required field for RequestForPositions.
func (m RequestForPositions) RestructuringType() (field.RestructuringType, errors.MessageRejectError) {
	var f field.RestructuringType
	err := m.Body.Get(&f)
	return f, err
}

//Seniority is a non-required field for RequestForPositions.
func (m RequestForPositions) Seniority() (field.Seniority, errors.MessageRejectError) {
	var f field.Seniority
	err := m.Body.Get(&f)
	return f, err
}

//NotionalPercentageOutstanding is a non-required field for RequestForPositions.
func (m RequestForPositions) NotionalPercentageOutstanding() (field.NotionalPercentageOutstanding, errors.MessageRejectError) {
	var f field.NotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//OriginalNotionalPercentageOutstanding is a non-required field for RequestForPositions.
func (m RequestForPositions) OriginalNotionalPercentageOutstanding() (field.OriginalNotionalPercentageOutstanding, errors.MessageRejectError) {
	var f field.OriginalNotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//AttachmentPoint is a non-required field for RequestForPositions.
func (m RequestForPositions) AttachmentPoint() (field.AttachmentPoint, errors.MessageRejectError) {
	var f field.AttachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//DetachmentPoint is a non-required field for RequestForPositions.
func (m RequestForPositions) DetachmentPoint() (field.DetachmentPoint, errors.MessageRejectError) {
	var f field.DetachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceDeterminationMethod is a non-required field for RequestForPositions.
func (m RequestForPositions) StrikePriceDeterminationMethod() (field.StrikePriceDeterminationMethod, errors.MessageRejectError) {
	var f field.StrikePriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryMethod is a non-required field for RequestForPositions.
func (m RequestForPositions) StrikePriceBoundaryMethod() (field.StrikePriceBoundaryMethod, errors.MessageRejectError) {
	var f field.StrikePriceBoundaryMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryPrecision is a non-required field for RequestForPositions.
func (m RequestForPositions) StrikePriceBoundaryPrecision() (field.StrikePriceBoundaryPrecision, errors.MessageRejectError) {
	var f field.StrikePriceBoundaryPrecision
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceDeterminationMethod is a non-required field for RequestForPositions.
func (m RequestForPositions) UnderlyingPriceDeterminationMethod() (field.UnderlyingPriceDeterminationMethod, errors.MessageRejectError) {
	var f field.UnderlyingPriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutType is a non-required field for RequestForPositions.
func (m RequestForPositions) OptPayoutType() (field.OptPayoutType, errors.MessageRejectError) {
	var f field.OptPayoutType
	err := m.Body.Get(&f)
	return f, err
}

//NoComplexEvents is a non-required field for RequestForPositions.
func (m RequestForPositions) NoComplexEvents() (field.NoComplexEvents, errors.MessageRejectError) {
	var f field.NoComplexEvents
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for RequestForPositions.
func (m RequestForPositions) Currency() (field.Currency, errors.MessageRejectError) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for RequestForPositions.
func (m RequestForPositions) NoLegs() (field.NoLegs, errors.MessageRejectError) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for RequestForPositions.
func (m RequestForPositions) NoUnderlyings() (field.NoUnderlyings, errors.MessageRejectError) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a required field for RequestForPositions.
func (m RequestForPositions) ClearingBusinessDate() (field.ClearingBusinessDate, errors.MessageRejectError) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessID is a non-required field for RequestForPositions.
func (m RequestForPositions) SettlSessID() (field.SettlSessID, errors.MessageRejectError) {
	var f field.SettlSessID
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessSubID is a non-required field for RequestForPositions.
func (m RequestForPositions) SettlSessSubID() (field.SettlSessSubID, errors.MessageRejectError) {
	var f field.SettlSessSubID
	err := m.Body.Get(&f)
	return f, err
}

//NoTradingSessions is a non-required field for RequestForPositions.
func (m RequestForPositions) NoTradingSessions() (field.NoTradingSessions, errors.MessageRejectError) {
	var f field.NoTradingSessions
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for RequestForPositions.
func (m RequestForPositions) TransactTime() (field.TransactTime, errors.MessageRejectError) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//ResponseTransportType is a non-required field for RequestForPositions.
func (m RequestForPositions) ResponseTransportType() (field.ResponseTransportType, errors.MessageRejectError) {
	var f field.ResponseTransportType
	err := m.Body.Get(&f)
	return f, err
}

//ResponseDestination is a non-required field for RequestForPositions.
func (m RequestForPositions) ResponseDestination() (field.ResponseDestination, errors.MessageRejectError) {
	var f field.ResponseDestination
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for RequestForPositions.
func (m RequestForPositions) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for RequestForPositions.
func (m RequestForPositions) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for RequestForPositions.
func (m RequestForPositions) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrency is a non-required field for RequestForPositions.
func (m RequestForPositions) SettlCurrency() (field.SettlCurrency, errors.MessageRejectError) {
	var f field.SettlCurrency
	err := m.Body.Get(&f)
	return f, err
}
