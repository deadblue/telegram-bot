package arguments

type SetPassportDataErrorsArgs struct {
	_BasicArgs
	errBuf []map[string]interface{}
}

func (a *SetPassportDataErrorsArgs) receiveError(err map[string]interface{}) {
	if a.errBuf == nil {
		a.errBuf = []map[string]interface{}{err}
	} else {
		a.errBuf = append(a.errBuf, err)
	}
	if a.beforeArchive == nil {
		a.beforeArchive = func() {
			a.getForm().WithJson("errors", a.errBuf)
		}
	}
}
func (a *SetPassportDataErrorsArgs) UserId(userId int) {
	a.getForm().WithInt("user_id", userId)
}
func (a *SetPassportDataErrorsArgs) AddDataFieldError() PassportDataFieldErrorBuilder {
	return new(implPassportDataFieldErrorBuilder).Init(a.receiveError)
}
func (a *SetPassportDataErrorsArgs) AddFrontSideError() PassportFrontSideErrorBuilder {
	return new(implPassportFrontSideErrorBuilder).Init(a.receiveError)
}
func (a *SetPassportDataErrorsArgs) AddReverseSideError() PassportReverseSideErrorBuilder {
	return new(implPassportReverseSideErrorBuilder).Init(a.receiveError)
}
func (a *SetPassportDataErrorsArgs) AddSelfieError() PassportSelfieErrorBuilder {
	return new(implPassportSelfieErrorBuilder).Init(a.receiveError)
}
func (a *SetPassportDataErrorsArgs) AddFileError() PassportFileErrorBuilder {
	return new(implPassportFileErrorBuilder).Init(a.receiveError)
}
func (a *SetPassportDataErrorsArgs) AddTranslationFileError() PassportTranslationFileErrorBuilder {
	return new(implPassportTranslationFileErrorBuilder).Init(a.receiveError)
}
func (a *SetPassportDataErrorsArgs) AddUnspecifiedError() PassportUnspecifiedErrorBuilder {
	return new(implPassportUnspecifiedErrorBuilder).Init(a.receiveError)
}

type PassportDataFieldErrorBuilder interface {
	ArgumentBuilder
	PersonalDetails() PassportDataFieldErrorBuilder
	Passport() PassportDataFieldErrorBuilder
	DriverLicense() PassportDataFieldErrorBuilder
	IdentityCard() PassportDataFieldErrorBuilder
	InternalPassport() PassportDataFieldErrorBuilder
	Address() PassportDataFieldErrorBuilder
	FieldName(name string) PassportDataFieldErrorBuilder
	DataHash(hash string) PassportDataFieldErrorBuilder
	Message(message string) PassportDataFieldErrorBuilder
}

type PassportFrontSideErrorBuilder interface {
	ArgumentBuilder
	Passport() PassportFrontSideErrorBuilder
	DriverLicense() PassportFrontSideErrorBuilder
	IdentityCard() PassportFrontSideErrorBuilder
	InternalPassport() PassportFrontSideErrorBuilder
	FileHash(hash string) PassportFrontSideErrorBuilder
	Message(message string) PassportFrontSideErrorBuilder
}

type PassportReverseSideErrorBuilder interface {
	ArgumentBuilder
	Passport() PassportReverseSideErrorBuilder
	DriverLicense() PassportReverseSideErrorBuilder
	IdentityCard() PassportReverseSideErrorBuilder
	InternalPassport() PassportReverseSideErrorBuilder
	FileHash(hash string) PassportReverseSideErrorBuilder
	Message(message string) PassportReverseSideErrorBuilder
}

type PassportSelfieErrorBuilder interface {
	ArgumentBuilder
	Passport() PassportSelfieErrorBuilder
	DriverLicense() PassportSelfieErrorBuilder
	IdentityCard() PassportSelfieErrorBuilder
	InternalPassport() PassportSelfieErrorBuilder
	FileHash(hash string) PassportSelfieErrorBuilder
	Message(message string) PassportSelfieErrorBuilder
}

type PassportFileErrorBuilder interface {
	ArgumentBuilder
	UtilityBill() PassportFileErrorBuilder
	BankStatement() PassportFileErrorBuilder
	RentalAgreement() PassportFileErrorBuilder
	PassportRegistration() PassportFileErrorBuilder
	TemporaryRegistration() PassportFileErrorBuilder
	FileHash(hash ...string) PassportFileErrorBuilder
	Message(message string) PassportFileErrorBuilder
}

type PassportTranslationFileErrorBuilder interface {
	ArgumentBuilder
	Passport() PassportTranslationFileErrorBuilder
	DriverLicense() PassportTranslationFileErrorBuilder
	IdentityCard() PassportTranslationFileErrorBuilder
	InternalPassport() PassportTranslationFileErrorBuilder
	UtilityBill() PassportTranslationFileErrorBuilder
	BankStatement() PassportTranslationFileErrorBuilder
	RentalAgreement() PassportTranslationFileErrorBuilder
	PassportRegistration() PassportTranslationFileErrorBuilder
	TemporaryRegistration() PassportTranslationFileErrorBuilder
	FileHash(hash ...string) PassportTranslationFileErrorBuilder
	Message(message string) PassportTranslationFileErrorBuilder
}

type PassportUnspecifiedErrorBuilder interface {
	ArgumentBuilder
	PersonalDetails() PassportUnspecifiedErrorBuilder
	Passport() PassportUnspecifiedErrorBuilder
	DriverLicense() PassportUnspecifiedErrorBuilder
	IdentityCard() PassportUnspecifiedErrorBuilder
	InternalPassport() PassportUnspecifiedErrorBuilder
	Address() PassportUnspecifiedErrorBuilder
	UtilityBill() PassportUnspecifiedErrorBuilder
	BankStatement() PassportUnspecifiedErrorBuilder
	RentalAgreement() PassportUnspecifiedErrorBuilder
	PassportRegistration() PassportUnspecifiedErrorBuilder
	TemporaryRegistration() PassportUnspecifiedErrorBuilder
	PhoneNumber() PassportUnspecifiedErrorBuilder
	Email() PassportUnspecifiedErrorBuilder
	ElementHash(hash string) PassportUnspecifiedErrorBuilder
	Message(message string) PassportUnspecifiedErrorBuilder
}

// The basic struct for all passport error builder
type basicPassportErrorBuilder struct {
	data     map[string]interface{}
	receiver func(map[string]interface{})
}

func (b *basicPassportErrorBuilder) set(name string, value interface{}) {
	if b.data == nil {
		b.data = make(map[string]interface{})
	}
	b.data[name] = value
}
func (b *basicPassportErrorBuilder) setSource(value string) {
	b.set("source", value)
}
func (b *basicPassportErrorBuilder) setType(value string) {
	b.set("type", value)
}
func (b *basicPassportErrorBuilder) setMessage(value string) {
	b.set("message", value)
}
func (b *basicPassportErrorBuilder) Finish() {
	b.receiver(b.data)
}

// The implementation for "PassportDataFieldErrorBuilder"
type implPassportDataFieldErrorBuilder struct {
	basicPassportErrorBuilder
}

func (b *implPassportDataFieldErrorBuilder) Init(receiver func(map[string]interface{})) PassportDataFieldErrorBuilder {
	b.receiver = receiver
	b.setSource("data")
	return b
}
func (b *implPassportDataFieldErrorBuilder) PersonalDetails() PassportDataFieldErrorBuilder {
	b.setType(passportPersonalDetails)
	return b
}
func (b *implPassportDataFieldErrorBuilder) Passport() PassportDataFieldErrorBuilder {
	b.setType(passportPassport)
	return b
}
func (b *implPassportDataFieldErrorBuilder) DriverLicense() PassportDataFieldErrorBuilder {
	b.setType(passportDriverLicense)
	return b
}
func (b *implPassportDataFieldErrorBuilder) IdentityCard() PassportDataFieldErrorBuilder {
	b.setType(passportDriverLicense)
	return b
}
func (b *implPassportDataFieldErrorBuilder) InternalPassport() PassportDataFieldErrorBuilder {
	b.setType(passportInternalPassport)
	return b
}
func (b *implPassportDataFieldErrorBuilder) Address() PassportDataFieldErrorBuilder {
	b.setType(passportAddress)
	return b
}
func (b *implPassportDataFieldErrorBuilder) FieldName(name string) PassportDataFieldErrorBuilder {
	b.set("field_name", name)
	return b
}
func (b *implPassportDataFieldErrorBuilder) DataHash(hash string) PassportDataFieldErrorBuilder {
	b.set("data_hash", hash)
	return b
}
func (b *implPassportDataFieldErrorBuilder) Message(message string) PassportDataFieldErrorBuilder {
	b.setMessage(message)
	return b
}

// The implementation for "PassportFrontSideErrorBuilder"
type implPassportFrontSideErrorBuilder struct {
	basicPassportErrorBuilder
}

func (b *implPassportFrontSideErrorBuilder) Init(receiver func(map[string]interface{})) PassportFrontSideErrorBuilder {
	b.receiver = receiver
	b.setSource("front_side")
	return b
}
func (b *implPassportFrontSideErrorBuilder) Passport() PassportFrontSideErrorBuilder {
	b.setType(passportPassport)
	return b
}
func (b *implPassportFrontSideErrorBuilder) DriverLicense() PassportFrontSideErrorBuilder {
	b.setType(passportDriverLicense)
	return b
}
func (b *implPassportFrontSideErrorBuilder) IdentityCard() PassportFrontSideErrorBuilder {
	b.setType(passportInternalPassport)
	return b
}
func (b *implPassportFrontSideErrorBuilder) InternalPassport() PassportFrontSideErrorBuilder {
	b.setType(passportInternalPassport)
	return b
}
func (b *implPassportFrontSideErrorBuilder) FileHash(hash string) PassportFrontSideErrorBuilder {
	b.set("file_hash", hash)
	return b
}
func (b *implPassportFrontSideErrorBuilder) Message(message string) PassportFrontSideErrorBuilder {
	b.setMessage(message)
	return b
}

// The implementation for "PassportReverseSideErrorBuilder"
type implPassportReverseSideErrorBuilder struct {
	basicPassportErrorBuilder
}

func (b *implPassportReverseSideErrorBuilder) Init(receiver func(map[string]interface{})) PassportReverseSideErrorBuilder {
	b.receiver = receiver
	b.setSource("reverse_side")
	return b
}
func (b *implPassportReverseSideErrorBuilder) Passport() PassportReverseSideErrorBuilder {
	b.setType(passportPassport)
	return b
}
func (b *implPassportReverseSideErrorBuilder) DriverLicense() PassportReverseSideErrorBuilder {
	b.setType(passportDriverLicense)
	return b
}
func (b *implPassportReverseSideErrorBuilder) IdentityCard() PassportReverseSideErrorBuilder {
	b.setType(passportIdentityCard)
	return b
}
func (b *implPassportReverseSideErrorBuilder) InternalPassport() PassportReverseSideErrorBuilder {
	b.setType(passportInternalPassport)
	return b
}
func (b *implPassportReverseSideErrorBuilder) FileHash(hash string) PassportReverseSideErrorBuilder {
	b.set("file_hash", hash)
	return b
}
func (b *implPassportReverseSideErrorBuilder) Message(message string) PassportReverseSideErrorBuilder {
	b.setMessage(message)
	return b
}

// The implementation for "PassportSelfieErrorBuilder"
type implPassportSelfieErrorBuilder struct {
	basicPassportErrorBuilder
}

func (b *implPassportSelfieErrorBuilder) Init(receiver func(map[string]interface{})) PassportSelfieErrorBuilder {
	b.receiver = receiver
	b.setSource("selfie")
	return b
}
func (b *implPassportSelfieErrorBuilder) Passport() PassportSelfieErrorBuilder {
	b.setType(passportPassport)
	return b
}
func (b *implPassportSelfieErrorBuilder) DriverLicense() PassportSelfieErrorBuilder {
	b.setType(passportDriverLicense)
	return b
}
func (b *implPassportSelfieErrorBuilder) IdentityCard() PassportSelfieErrorBuilder {
	b.setType(passportIdentityCard)
	return b
}
func (b *implPassportSelfieErrorBuilder) InternalPassport() PassportSelfieErrorBuilder {
	b.setType(passportInternalPassport)
	return b
}
func (b *implPassportSelfieErrorBuilder) FileHash(hash string) PassportSelfieErrorBuilder {
	b.data["file_hash"] = hash
	return b
}
func (b *implPassportSelfieErrorBuilder) Message(message string) PassportSelfieErrorBuilder {
	b.setMessage(message)
	return b
}

// The implementation for "PassportFileErrorBuilder"
type implPassportFileErrorBuilder struct {
	basicPassportErrorBuilder
}

func (b *implPassportFileErrorBuilder) Init(receiver func(map[string]interface{})) PassportFileErrorBuilder {
	b.receiver = receiver
	return b
}
func (b *implPassportFileErrorBuilder) UtilityBill() PassportFileErrorBuilder {
	b.setType(passportUtilityBill)
	return b
}
func (b *implPassportFileErrorBuilder) BankStatement() PassportFileErrorBuilder {
	b.setType(passportBankStatement)
	return b
}
func (b *implPassportFileErrorBuilder) RentalAgreement() PassportFileErrorBuilder {
	b.setType(passportRentalAgreement)
	return b
}
func (b *implPassportFileErrorBuilder) PassportRegistration() PassportFileErrorBuilder {
	b.setType(passportPassportRegistration)
	return b
}
func (b *implPassportFileErrorBuilder) TemporaryRegistration() PassportFileErrorBuilder {
	b.setType(passportTemporaryRegistration)
	return b
}
func (b *implPassportFileErrorBuilder) FileHash(hash ...string) PassportFileErrorBuilder {
	if len(hash) == 1 {
		b.setSource("file")
		b.set("file_hash", hash[0])
	} else {
		b.setSource("files")
		b.set("file_hashes", hash)
	}
	return b
}
func (b *implPassportFileErrorBuilder) Message(message string) PassportFileErrorBuilder {
	b.setMessage(message)
	return b
}

// The implementation for "PassportTranslationFileErrorBuilder"
type implPassportTranslationFileErrorBuilder struct {
	basicPassportErrorBuilder
}

func (b *implPassportTranslationFileErrorBuilder) Init(receiver func(map[string]interface{})) PassportTranslationFileErrorBuilder {
	b.receiver = receiver
	return b
}
func (b *implPassportTranslationFileErrorBuilder) Passport() PassportTranslationFileErrorBuilder {
	b.setType(passportPassport)
	return b
}
func (b *implPassportTranslationFileErrorBuilder) DriverLicense() PassportTranslationFileErrorBuilder {
	b.setType(passportDriverLicense)
	return b
}
func (b *implPassportTranslationFileErrorBuilder) IdentityCard() PassportTranslationFileErrorBuilder {
	b.setType(passportIdentityCard)
	return b
}
func (b *implPassportTranslationFileErrorBuilder) InternalPassport() PassportTranslationFileErrorBuilder {
	b.setType(passportInternalPassport)
	return b
}
func (b *implPassportTranslationFileErrorBuilder) UtilityBill() PassportTranslationFileErrorBuilder {
	b.setType(passportUtilityBill)
	return b
}
func (b *implPassportTranslationFileErrorBuilder) BankStatement() PassportTranslationFileErrorBuilder {
	b.setType(passportBankStatement)
	return b
}
func (b *implPassportTranslationFileErrorBuilder) RentalAgreement() PassportTranslationFileErrorBuilder {
	b.setType(passportRentalAgreement)
	return b
}
func (b *implPassportTranslationFileErrorBuilder) PassportRegistration() PassportTranslationFileErrorBuilder {
	b.setType(passportPassportRegistration)
	return b
}
func (b *implPassportTranslationFileErrorBuilder) TemporaryRegistration() PassportTranslationFileErrorBuilder {
	b.setType(passportTemporaryRegistration)
	return b
}
func (b *implPassportTranslationFileErrorBuilder) FileHash(hash ...string) PassportTranslationFileErrorBuilder {
	if len(hash) == 1 {
		b.setSource("translation_file")
		b.set("file_hash", hash[0])
	} else {
		b.setSource("translation_files")
		b.set("file_hashes", hash)
	}
	return b
}
func (b *implPassportTranslationFileErrorBuilder) Message(message string) PassportTranslationFileErrorBuilder {
	b.setMessage(message)
	return b
}

// The implementation for "PassportUnspecifiedErrorBuilder"
type implPassportUnspecifiedErrorBuilder struct {
	basicPassportErrorBuilder
}

func (b *implPassportUnspecifiedErrorBuilder) Init(receiver func(map[string]interface{})) PassportUnspecifiedErrorBuilder {
	b.receiver = receiver
	b.setSource("unspecified")
	return b
}
func (b *implPassportUnspecifiedErrorBuilder) PersonalDetails() PassportUnspecifiedErrorBuilder {
	b.setType(passportPersonalDetails)
	return b
}
func (b *implPassportUnspecifiedErrorBuilder) Passport() PassportUnspecifiedErrorBuilder {
	b.setType(passportPassport)
	return b
}
func (b *implPassportUnspecifiedErrorBuilder) DriverLicense() PassportUnspecifiedErrorBuilder {
	b.setType(passportDriverLicense)
	return b
}
func (b *implPassportUnspecifiedErrorBuilder) IdentityCard() PassportUnspecifiedErrorBuilder {
	b.setType(passportIdentityCard)
	return b
}
func (b *implPassportUnspecifiedErrorBuilder) InternalPassport() PassportUnspecifiedErrorBuilder {
	b.setType(passportInternalPassport)
	return b
}
func (b *implPassportUnspecifiedErrorBuilder) Address() PassportUnspecifiedErrorBuilder {
	b.setType(passportAddress)
	return b
}
func (b *implPassportUnspecifiedErrorBuilder) UtilityBill() PassportUnspecifiedErrorBuilder {
	b.setType(passportUtilityBill)
	return b
}
func (b *implPassportUnspecifiedErrorBuilder) BankStatement() PassportUnspecifiedErrorBuilder {
	b.setType(passportBankStatement)
	return b
}
func (b *implPassportUnspecifiedErrorBuilder) RentalAgreement() PassportUnspecifiedErrorBuilder {
	b.setType(passportRentalAgreement)
	return b
}
func (b *implPassportUnspecifiedErrorBuilder) PassportRegistration() PassportUnspecifiedErrorBuilder {
	b.setType(passportPassportRegistration)
	return b
}
func (b *implPassportUnspecifiedErrorBuilder) TemporaryRegistration() PassportUnspecifiedErrorBuilder {
	b.setType(passportTemporaryRegistration)
	return b
}
func (b *implPassportUnspecifiedErrorBuilder) PhoneNumber() PassportUnspecifiedErrorBuilder {
	b.setType(passportPhoneNumber)
	return b
}
func (b *implPassportUnspecifiedErrorBuilder) Email() PassportUnspecifiedErrorBuilder {
	b.setType(passportEmail)
	return b
}
func (b *implPassportUnspecifiedErrorBuilder) ElementHash(hash string) PassportUnspecifiedErrorBuilder {
	b.set("element_hash", hash)
	return b
}
func (b *implPassportUnspecifiedErrorBuilder) Message(message string) PassportUnspecifiedErrorBuilder {
	b.setMessage(message)
	return b
}
