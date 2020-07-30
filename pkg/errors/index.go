package errors

// ErrDefault -
const (
	ErrDefault            = "errors.default"
	ErrNotFound           = "errors.notFound"
	ErrNotFoundUser       = "errors.notFound.user"
	ErrNotFoundBusiness   = "errors.notFound.business"
	ErrNotFoundProject    = "errors.notFound.project"
	ErrNotFoundProperty   = "errors.notFound.property"
	ErrNotFoundPerson     = "errors.notFound.person"
	ErrNotFoundInvoice    = "errors.notFound.invoice"
	ErrNotFoundBlock      = "errors.notFound.block"
	ErrUserNotInBusiness  = "errors.user.notInBusiness"
	ErrInputInvalid       = "errors.inputInvalid"
	ErrMin0               = "errors.min.0"
	ErrDomainDuplicate    = "errors.domain.duplicate"
	ErrUserDuplicate      = "errors.user.duplicate"
	ErrToken              = "errors.token"
	ErrSignupEmailInvalid = "errors.signup.email.invalid"
	ErrSigninIncorrect    = "errors.signin.incorrect"
	ErrInsert             = "errors.insert"
	ErrUpdate             = "errors.update"
	ErrDelete             = "errors.delete"
	ErrEncoding           = "errors.encoding"
	ErrOwner              = "errors.owner"
	ErrDuplicate          = "errors.duplicate"

	// property errors
	ErrPropertyAlreadyTaken = "errors.property.alreadyTaken"
	ErrInvoiceAlreadyPaid   = "errors.invoice.alreadyPaid"
	ErrDuplicateReceipt     = "errors.duplicate.receipt"

	// block errors
	ErrBlockTypeNotAllowed = "errors.blocks.notAllowed"
)
