package enums

// Countries slice
var Countries = make([]string, 0)

// Currency -
type Currency struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

// Currencies slice
var Currencies = make([]Currency, 0)

// Enums
const (
	EntityBusiness = "business"
	EntityProject  = "project"
	EntityPerson   = "person"
	EntityProperty = "property"
	EntityInvoice  = "invoice"

	RoleOwner   = "owner"
	RoleCoOwner = "co-owner"
	RoleAgent   = "agent"
	RoleClient  = "client"

	AccountStatusActive   = "active"
	AccountStatusInactive = "inactive"
	AccountStatusPending  = "pending"
	AccountStatusDeleted  = "deleted"

	StatusActive     = "active"
	StatusPending    = "pending"
	StatusPaid       = "paid"
	StatusCancelled  = "cancelled"
	StatusTerminated = "terminated"
	StatusApproved   = "approved"
	StatusAvailable  = "available"
	StatusOnGoing    = "ongoing"
	StatusAcquired   = "acquired"
	StatusDraft      = "draft"
	StatusOverdue    = "overdue"

	PropertyTypeLot   = "lot"
	PropertyTypeHouse = "house"

	PaymentSchemeCash        = "cash"
	PaymentSchemeInstallment = "installment"

	PaymentPeriodMonthly = "monthly"
	PaymentPeriodYearly  = "yearly"

	InvoiceBlockIntro   = "intro"
	InvoiceBlockItem    = "item"
	InvoiceBlockSummary = "summary"

	DateTimeForm = "2006-01-02T15:04" // ex. 2020-12-04T17:30
)
