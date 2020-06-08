package errors

// ErrDefault -
const (
	ErrDefault            = "errors.default"
	ErrNotFound           = "errors.notFound"
	ErrNotFoundUser       = "errors.notFound.user"
	ErrNotFoundBusiness   = "errors.notFound.business"
	ErrNotFoundProject    = "errors.notFound.project"
	ErrNotFoundLot        = "errors.notFound.lot"
	ErrNotFoundPerson     = "errors.notFound.person"
	ErrInputInvalid       = "errors.inputInvalid"
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

	// lot errors
	ErrLotAlreadyTaken = "errors.lot.alreadyTaken"
)
