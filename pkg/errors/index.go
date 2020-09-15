package errors

// just for the i18next-parser to include this file for parsing
func t(err string) string {
	return err
}

// ErrDefault -
// t if for the i18n parser to be parsed in the frontend
var (
	ErrDefault = t("errors.default")

	ErrNotFound         = t("errors.notFound.default")
	ErrNotFoundUser     = t("errors.notFound.user")
	ErrNotFoundBusiness = t("errors.notFound.business")
	ErrNotFoundProject  = t("errors.notFound.project")
	ErrNotFoundProperty = t("errors.notFound.property")
	ErrNotFoundPerson   = t("errors.notFound.person")
	ErrNotFoundInvoice  = t("errors.notFound.invoice")
	ErrNotFoundBlock    = t("errors.notFound.block")
	ErrNotFoundFile     = t("errors.notFound.file")

	ErrDuplicate        = t("errors.duplicate.default")
	ErrDuplicateDomain  = t("errors.duplicate.domain")
	ErrDuplicateUser    = t("errors.duplicate.user")
	ErrDuplicateReceipt = t("errors.duplicate.receipt")

	// Validation errors
	ErrInputInvalid = t("errors.inputInvalid")
	ErrMin0         = t("errors.min.0")
	ErrToken        = t("errors.token")

	// Operation Errors
	ErrInsert   = t("errors.insert")
	ErrUpdate   = t("errors.update")
	ErrDelete   = t("errors.delete")
	ErrEncoding = t("errors.encoding")
	ErrOwner    = t("errors.owner")

	// Entity Errors
	ErrSigninIncorrect      = t("errors.signin.incorrect")
	ErrSignupEmailInvalid   = t("errors.signup.emailInvalid")
	ErrUserNotInBusiness    = t("errors.user.notInBusiness")
	ErrPropertyAlreadyTaken = t("errors.property.alreadyTaken")
	ErrInvoiceAlreadyPaid   = t("errors.invoice.alreadyPaid")
	ErrBlockTypeNotAllowed  = t("errors.blocks.notAllowed")
)
