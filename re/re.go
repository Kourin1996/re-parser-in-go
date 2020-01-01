package re

const (
	CHAR         = "CHAR"
	CONCAT       = "CONCAT"
	CHOICE       = "CHOICE"
	OPTION       = "OPTION"
	ZERO_OR_MANY = "ZERO_OR_MANY"
	ONE_OR_MANY  = "ONE_OR_MANY"
	GROUP        = "GROUP"
)

type RegularExpressionVisitor interface {
	VisitChar(e *REChar)
	VisitConcat(e *REConcat)
	VisitChoice(e *REChoice)
	VisitOption(e *REOption)
	VisitZeroOrMany(e *REZeroOrMany)
	VisitOneOrMany(e *REOneOrMany)
	VisitGroup(e *REGroup)
}

type RegularExpression interface {
	GetType() string
	Accept(visitor RegularExpressionVisitor)
}

type REChar struct {
	Ch byte
}

func (re *REChar) GetType() string {
	return CHAR
}

func (re *REChar) Accept(visitor RegularExpressionVisitor) {
	visitor.VisitChar(re)
}

type REConcat struct {
	First  RegularExpression
	Second RegularExpression
}

func (re *REConcat) GetType() string {
	return CONCAT
}

func (re *REConcat) Accept(visitor RegularExpressionVisitor) {
	visitor.VisitConcat(re)
}

type REChoice struct {
	First  RegularExpression
	Second RegularExpression
}

func (re *REChoice) GetType() string {
	return CHOICE
}

func (re *REChoice) Accept(visitor RegularExpressionVisitor) {
	visitor.VisitChoice(re)
}

type REOption struct {
	Ex RegularExpression
}

func (re *REOption) GetType() string {
	return OPTION
}

func (re *REOption) Accept(visitor RegularExpressionVisitor) {
	visitor.VisitOption(re)
}

type REZeroOrMany struct {
	Ex RegularExpression
}

func (re *REZeroOrMany) GetType() string {
	return ZERO_OR_MANY
}

func (re *REZeroOrMany) Accept(visitor RegularExpressionVisitor) {
	visitor.VisitZeroOrMany(re)
}

type REOneOrMany struct {
	Ex RegularExpression
}

func (re *REOneOrMany) GetType() string {
	return ONE_OR_MANY
}

func (re *REOneOrMany) Accept(visitor RegularExpressionVisitor) {
	visitor.VisitOneOrMany(re)
}

type REGroup struct {
	GroupName string
	Ex        RegularExpression
}

func (re *REGroup) GetType() string {
	return GROUP
}

func (re *REGroup) Accept(visitor RegularExpressionVisitor) {
	visitor.VisitGroup(re)
}
