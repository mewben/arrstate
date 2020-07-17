package enums

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

	PropertyTypeLot   = "lot"
	PropertyTypeHouse = "house"

	PaymentSchemeCash        = "cash"
	PaymentSchemeInstallment = "installment"

	PaymentPeriodMonthly = "monthly"
	PaymentPeriodYearly  = "yearly"

	InvoiceBlockIntro   = "intro"
	InvoiceBlockItem    = "item"
	InvoiceBlockSummary = "summary"
)
