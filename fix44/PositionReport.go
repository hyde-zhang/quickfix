package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//PositionReport msg type = AP.
type PositionReport struct {
	message.Message
}

//PositionReportBuilder builds PositionReport messages.
type PositionReportBuilder struct {
	message.MessageBuilder
}

//CreatePositionReportBuilder returns an initialized PositionReportBuilder with specified required fields.
func CreatePositionReportBuilder(
	posmaintrptid field.PosMaintRptID,
	posreqresult field.PosReqResult,
	clearingbusinessdate field.ClearingBusinessDate,
	account field.Account,
	accounttype field.AccountType,
	settlprice field.SettlPrice,
	settlpricetype field.SettlPriceType,
	priorsettlprice field.PriorSettlPrice) PositionReportBuilder {
	var builder PositionReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(posmaintrptid)
	builder.Body.Set(posreqresult)
	builder.Body.Set(clearingbusinessdate)
	builder.Body.Set(account)
	builder.Body.Set(accounttype)
	builder.Body.Set(settlprice)
	builder.Body.Set(settlpricetype)
	builder.Body.Set(priorsettlprice)
	return builder
}

//PosMaintRptID is a required field for PositionReport.
func (m PositionReport) PosMaintRptID() (field.PosMaintRptID, errors.MessageRejectError) {
	var f field.PosMaintRptID
	err := m.Body.Get(&f)
	return f, err
}

//PosReqID is a non-required field for PositionReport.
func (m PositionReport) PosReqID() (field.PosReqID, errors.MessageRejectError) {
	var f field.PosReqID
	err := m.Body.Get(&f)
	return f, err
}

//PosReqType is a non-required field for PositionReport.
func (m PositionReport) PosReqType() (field.PosReqType, errors.MessageRejectError) {
	var f field.PosReqType
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a non-required field for PositionReport.
func (m PositionReport) SubscriptionRequestType() (field.SubscriptionRequestType, errors.MessageRejectError) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}

//TotalNumPosReports is a non-required field for PositionReport.
func (m PositionReport) TotalNumPosReports() (field.TotalNumPosReports, errors.MessageRejectError) {
	var f field.TotalNumPosReports
	err := m.Body.Get(&f)
	return f, err
}

//UnsolicitedIndicator is a non-required field for PositionReport.
func (m PositionReport) UnsolicitedIndicator() (field.UnsolicitedIndicator, errors.MessageRejectError) {
	var f field.UnsolicitedIndicator
	err := m.Body.Get(&f)
	return f, err
}

//PosReqResult is a required field for PositionReport.
func (m PositionReport) PosReqResult() (field.PosReqResult, errors.MessageRejectError) {
	var f field.PosReqResult
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a required field for PositionReport.
func (m PositionReport) ClearingBusinessDate() (field.ClearingBusinessDate, errors.MessageRejectError) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessID is a non-required field for PositionReport.
func (m PositionReport) SettlSessID() (field.SettlSessID, errors.MessageRejectError) {
	var f field.SettlSessID
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessSubID is a non-required field for PositionReport.
func (m PositionReport) SettlSessSubID() (field.SettlSessSubID, errors.MessageRejectError) {
	var f field.SettlSessSubID
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for PositionReport.
func (m PositionReport) NoPartyIDs() (field.NoPartyIDs, errors.MessageRejectError) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//Account is a required field for PositionReport.
func (m PositionReport) Account() (field.Account, errors.MessageRejectError) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AcctIDSource is a non-required field for PositionReport.
func (m PositionReport) AcctIDSource() (field.AcctIDSource, errors.MessageRejectError) {
	var f field.AcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a required field for PositionReport.
func (m PositionReport) AccountType() (field.AccountType, errors.MessageRejectError) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for PositionReport.
func (m PositionReport) Symbol() (field.Symbol, errors.MessageRejectError) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for PositionReport.
func (m PositionReport) SymbolSfx() (field.SymbolSfx, errors.MessageRejectError) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for PositionReport.
func (m PositionReport) SecurityID() (field.SecurityID, errors.MessageRejectError) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for PositionReport.
func (m PositionReport) SecurityIDSource() (field.SecurityIDSource, errors.MessageRejectError) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for PositionReport.
func (m PositionReport) NoSecurityAltID() (field.NoSecurityAltID, errors.MessageRejectError) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for PositionReport.
func (m PositionReport) Product() (field.Product, errors.MessageRejectError) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for PositionReport.
func (m PositionReport) CFICode() (field.CFICode, errors.MessageRejectError) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for PositionReport.
func (m PositionReport) SecurityType() (field.SecurityType, errors.MessageRejectError) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for PositionReport.
func (m PositionReport) SecuritySubType() (field.SecuritySubType, errors.MessageRejectError) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for PositionReport.
func (m PositionReport) MaturityMonthYear() (field.MaturityMonthYear, errors.MessageRejectError) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for PositionReport.
func (m PositionReport) MaturityDate() (field.MaturityDate, errors.MessageRejectError) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for PositionReport.
func (m PositionReport) CouponPaymentDate() (field.CouponPaymentDate, errors.MessageRejectError) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for PositionReport.
func (m PositionReport) IssueDate() (field.IssueDate, errors.MessageRejectError) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for PositionReport.
func (m PositionReport) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, errors.MessageRejectError) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for PositionReport.
func (m PositionReport) RepurchaseTerm() (field.RepurchaseTerm, errors.MessageRejectError) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for PositionReport.
func (m PositionReport) RepurchaseRate() (field.RepurchaseRate, errors.MessageRejectError) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for PositionReport.
func (m PositionReport) Factor() (field.Factor, errors.MessageRejectError) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for PositionReport.
func (m PositionReport) CreditRating() (field.CreditRating, errors.MessageRejectError) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for PositionReport.
func (m PositionReport) InstrRegistry() (field.InstrRegistry, errors.MessageRejectError) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for PositionReport.
func (m PositionReport) CountryOfIssue() (field.CountryOfIssue, errors.MessageRejectError) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for PositionReport.
func (m PositionReport) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for PositionReport.
func (m PositionReport) LocaleOfIssue() (field.LocaleOfIssue, errors.MessageRejectError) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for PositionReport.
func (m PositionReport) RedemptionDate() (field.RedemptionDate, errors.MessageRejectError) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for PositionReport.
func (m PositionReport) StrikePrice() (field.StrikePrice, errors.MessageRejectError) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for PositionReport.
func (m PositionReport) StrikeCurrency() (field.StrikeCurrency, errors.MessageRejectError) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for PositionReport.
func (m PositionReport) OptAttribute() (field.OptAttribute, errors.MessageRejectError) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for PositionReport.
func (m PositionReport) ContractMultiplier() (field.ContractMultiplier, errors.MessageRejectError) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for PositionReport.
func (m PositionReport) CouponRate() (field.CouponRate, errors.MessageRejectError) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for PositionReport.
func (m PositionReport) SecurityExchange() (field.SecurityExchange, errors.MessageRejectError) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for PositionReport.
func (m PositionReport) Issuer() (field.Issuer, errors.MessageRejectError) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for PositionReport.
func (m PositionReport) EncodedIssuerLen() (field.EncodedIssuerLen, errors.MessageRejectError) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for PositionReport.
func (m PositionReport) EncodedIssuer() (field.EncodedIssuer, errors.MessageRejectError) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for PositionReport.
func (m PositionReport) SecurityDesc() (field.SecurityDesc, errors.MessageRejectError) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for PositionReport.
func (m PositionReport) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, errors.MessageRejectError) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for PositionReport.
func (m PositionReport) EncodedSecurityDesc() (field.EncodedSecurityDesc, errors.MessageRejectError) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for PositionReport.
func (m PositionReport) Pool() (field.Pool, errors.MessageRejectError) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for PositionReport.
func (m PositionReport) ContractSettlMonth() (field.ContractSettlMonth, errors.MessageRejectError) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for PositionReport.
func (m PositionReport) CPProgram() (field.CPProgram, errors.MessageRejectError) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for PositionReport.
func (m PositionReport) CPRegType() (field.CPRegType, errors.MessageRejectError) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for PositionReport.
func (m PositionReport) NoEvents() (field.NoEvents, errors.MessageRejectError) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for PositionReport.
func (m PositionReport) DatedDate() (field.DatedDate, errors.MessageRejectError) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for PositionReport.
func (m PositionReport) InterestAccrualDate() (field.InterestAccrualDate, errors.MessageRejectError) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for PositionReport.
func (m PositionReport) Currency() (field.Currency, errors.MessageRejectError) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//SettlPrice is a required field for PositionReport.
func (m PositionReport) SettlPrice() (field.SettlPrice, errors.MessageRejectError) {
	var f field.SettlPrice
	err := m.Body.Get(&f)
	return f, err
}

//SettlPriceType is a required field for PositionReport.
func (m PositionReport) SettlPriceType() (field.SettlPriceType, errors.MessageRejectError) {
	var f field.SettlPriceType
	err := m.Body.Get(&f)
	return f, err
}

//PriorSettlPrice is a required field for PositionReport.
func (m PositionReport) PriorSettlPrice() (field.PriorSettlPrice, errors.MessageRejectError) {
	var f field.PriorSettlPrice
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for PositionReport.
func (m PositionReport) NoLegs() (field.NoLegs, errors.MessageRejectError) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for PositionReport.
func (m PositionReport) NoUnderlyings() (field.NoUnderlyings, errors.MessageRejectError) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//NoPositions is a non-required field for PositionReport.
func (m PositionReport) NoPositions() (field.NoPositions, errors.MessageRejectError) {
	var f field.NoPositions
	err := m.Body.Get(&f)
	return f, err
}

//NoPosAmt is a non-required field for PositionReport.
func (m PositionReport) NoPosAmt() (field.NoPosAmt, errors.MessageRejectError) {
	var f field.NoPosAmt
	err := m.Body.Get(&f)
	return f, err
}

//RegistStatus is a non-required field for PositionReport.
func (m PositionReport) RegistStatus() (field.RegistStatus, errors.MessageRejectError) {
	var f field.RegistStatus
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryDate is a non-required field for PositionReport.
func (m PositionReport) DeliveryDate() (field.DeliveryDate, errors.MessageRejectError) {
	var f field.DeliveryDate
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for PositionReport.
func (m PositionReport) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for PositionReport.
func (m PositionReport) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for PositionReport.
func (m PositionReport) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
