package arguments

import "io"

type basicPassportErrorBuilder struct {
	data     map[string]interface{}
	receiver func(map[string]interface{})
}

func (b *basicPassportErrorBuilder) init(receiver func(map[string]interface{})) {
	b.data = make(map[string]interface{})
	b.receiver = receiver
}

type PassportDataFieldErrorBuilder struct {
	basicPassportErrorBuilder
}

func (b *PassportDataFieldErrorBuilder) PersonalDetails() *PassportDataFieldErrorBuilder {
	b.data["type"] = "personal_details"
	return b
}
func (b *PassportDataFieldErrorBuilder) Passport() *PassportDataFieldErrorBuilder {
	b.data["type"] = "passport"
	return b
}
func (b *PassportDataFieldErrorBuilder) DriverLicense() *PassportDataFieldErrorBuilder {
	b.data["type"] = "driver_license"
	return b
}
func (b *PassportDataFieldErrorBuilder) IdentityCard() *PassportDataFieldErrorBuilder {
	b.data["type"] = "identity_card"
	return b
}
func (b *PassportDataFieldErrorBuilder) InternalPassport() *PassportDataFieldErrorBuilder {
	b.data["type"] = "internal_passport"
	return b
}
func (b *PassportDataFieldErrorBuilder) Address() *PassportDataFieldErrorBuilder {
	b.data["type"] = "address"
	return b
}
func (b *PassportDataFieldErrorBuilder) FieldName(name string) *PassportDataFieldErrorBuilder {
	b.data["field_name"] = name
	return b
}
func (b *PassportDataFieldErrorBuilder) DataHash(hash string) *PassportDataFieldErrorBuilder {
	b.data["data_hash"] = hash
	return b
}
func (b *PassportDataFieldErrorBuilder) Message(message string) *PassportDataFieldErrorBuilder {
	b.data["message"] = message
	return b
}
func (b *PassportDataFieldErrorBuilder) Finish() {
	b.data["source"] = "data"
	b.receiver(b.data)
}

type PassportFrontSideErrorBuilder struct {
	basicPassportErrorBuilder
}

func (b *PassportFrontSideErrorBuilder) Passport() *PassportFrontSideErrorBuilder {
	b.data["type"] = "passport"
	return b
}
func (b *PassportFrontSideErrorBuilder) DriverLicense() *PassportFrontSideErrorBuilder {
	b.data["type"] = "driver_license"
	return b
}
func (b *PassportFrontSideErrorBuilder) IdentityCard() *PassportFrontSideErrorBuilder {
	b.data["type"] = "identity_card"
	return b
}
func (b *PassportFrontSideErrorBuilder) InternalPassport() *PassportFrontSideErrorBuilder {
	b.data["type"] = "internal_passport"
	return b
}
func (b *PassportFrontSideErrorBuilder) FileHash(hash string) *PassportFrontSideErrorBuilder {
	b.data["file_hash"] = hash
	return b
}
func (b *PassportFrontSideErrorBuilder) Message(message string) *PassportFrontSideErrorBuilder {
	b.data["message"] = message
	return b
}
func (b *PassportFrontSideErrorBuilder) Finish() {
	b.data["source"] = "front_side"
	b.receiver(b.data)
}

type PassportReverseSideErrorBuilder struct {
	basicPassportErrorBuilder
}

func (b *PassportReverseSideErrorBuilder) Passport() *PassportReverseSideErrorBuilder {
	b.data["type"] = "passport"
	return b
}
func (b *PassportReverseSideErrorBuilder) DriverLicense() *PassportReverseSideErrorBuilder {
	b.data["type"] = "driver_license"
	return b
}
func (b *PassportReverseSideErrorBuilder) IdentityCard() *PassportReverseSideErrorBuilder {
	b.data["type"] = "identity_card"
	return b
}
func (b *PassportReverseSideErrorBuilder) InternalPassport() *PassportReverseSideErrorBuilder {
	b.data["type"] = "internal_passport"
	return b
}
func (b *PassportReverseSideErrorBuilder) FileHash(hash string) *PassportReverseSideErrorBuilder {
	b.data["file_hash"] = hash
	return b
}
func (b *PassportReverseSideErrorBuilder) Message(message string) *PassportReverseSideErrorBuilder {
	b.data["message"] = message
	return b
}
func (b *PassportReverseSideErrorBuilder) Finish() {
	b.data["source"] = "reverse_side"
	b.receiver(b.data)
}

type PassportSelfieErrorBuilder struct {
	basicPassportErrorBuilder
}

func (b *PassportSelfieErrorBuilder) Passport() *PassportSelfieErrorBuilder {
	b.data["type"] = "passport"
	return b
}
func (b *PassportSelfieErrorBuilder) DriverLicense() *PassportSelfieErrorBuilder {
	b.data["type"] = "driver_license"
	return b
}
func (b *PassportSelfieErrorBuilder) IdentityCard() *PassportSelfieErrorBuilder {
	b.data["type"] = "identity_card"
	return b
}
func (b *PassportSelfieErrorBuilder) InternalPassport() *PassportSelfieErrorBuilder {
	b.data["type"] = "internal_passport"
	return b
}
func (b *PassportSelfieErrorBuilder) FileHash(hash string) *PassportSelfieErrorBuilder {
	b.data["file_hash"] = hash
	return b
}
func (b *PassportSelfieErrorBuilder) Message(message string) *PassportSelfieErrorBuilder {
	b.data["message"] = message
	return b
}
func (b *PassportSelfieErrorBuilder) Finish() {
	b.data["source"] = "selfie"
	b.receiver(b.data)
}

type PassportFileErrorBuilder struct {
	basicPassportErrorBuilder
}

func (b *PassportFileErrorBuilder) UtilityBill() *PassportFileErrorBuilder {
	b.data["type"] = "utility_bill"
	return b
}
func (b *PassportFileErrorBuilder) BankStatement() *PassportFileErrorBuilder {
	b.data["type"] = "bank_statement"
	return b
}
func (b *PassportFileErrorBuilder) RentalAgreement() *PassportFileErrorBuilder {
	b.data["type"] = "rental_agreement"
	return b
}
func (b *PassportFileErrorBuilder) PassportRegistration() *PassportFileErrorBuilder {
	b.data["type"] = "passport_registration"
	return b
}
func (b *PassportFileErrorBuilder) TemporaryRegistration() *PassportFileErrorBuilder {
	b.data["type"] = "temporary_registration"
	return b
}
func (b *PassportFileErrorBuilder) FileHash(hash ...string) *PassportFileErrorBuilder {
	if len(hash) == 1 {
		b.data["source"] = "file"
		b.data["file_hash"] = hash[0]
	} else {
		b.data["source"] = "files"
		b.data["file_hashes"] = hash
	}
	return b
}
func (b *PassportFileErrorBuilder) Message(message string) *PassportFileErrorBuilder {
	b.data["message"] = message
	return b
}
func (b *PassportFileErrorBuilder) Finish() {
	b.receiver(b.data)
}

type PassportTranslationFileErrorBuilder struct {
	basicPassportErrorBuilder
}

func (b *PassportTranslationFileErrorBuilder) Passport() *PassportTranslationFileErrorBuilder {
	b.data["type"] = "passport"
	return b
}
func (b *PassportTranslationFileErrorBuilder) DriverLicense() *PassportTranslationFileErrorBuilder {
	b.data["type"] = "driver_license"
	return b
}
func (b *PassportTranslationFileErrorBuilder) IdentityCard() *PassportTranslationFileErrorBuilder {
	b.data["type"] = "identity_card"
	return b
}
func (b *PassportTranslationFileErrorBuilder) InternalPassport() *PassportTranslationFileErrorBuilder {
	b.data["type"] = "internal_passport"
	return b
}
func (b *PassportTranslationFileErrorBuilder) UtilityBill() *PassportTranslationFileErrorBuilder {
	b.data["type"] = "utility_bill"
	return b
}
func (b *PassportTranslationFileErrorBuilder) BankStatement() *PassportTranslationFileErrorBuilder {
	b.data["type"] = "bank_statement"
	return b
}
func (b *PassportTranslationFileErrorBuilder) RentalAgreement() *PassportTranslationFileErrorBuilder {
	b.data["type"] = "rental_agreement"
	return b
}
func (b *PassportTranslationFileErrorBuilder) PassportRegistration() *PassportTranslationFileErrorBuilder {
	b.data["type"] = "passport_registration"
	return b
}
func (b *PassportTranslationFileErrorBuilder) TemporaryRegistration() *PassportTranslationFileErrorBuilder {
	b.data["type"] = "temporary_registration"
	return b
}
func (b *PassportTranslationFileErrorBuilder) FileHash(hash ...string) *PassportTranslationFileErrorBuilder {
	if len(hash) == 1 {
		b.data["source"] = "translation_file"
		b.data["file_hash"] = hash[0]
	} else {
		b.data["source"] = "translation_files"
		b.data["file_hashes"] = hash
	}
	return b
}
func (b *PassportTranslationFileErrorBuilder) Message(message string) *PassportTranslationFileErrorBuilder {
	b.data["message"] = message
	return b
}
func (b *PassportTranslationFileErrorBuilder) Finish() {
	b.receiver(b.data)
}

type PassportUnspecifiedErrorBuilder struct {
	basicPassportErrorBuilder
}

func (b *PassportUnspecifiedErrorBuilder) PersonalDetails() *PassportUnspecifiedErrorBuilder {
	b.data["type"] = "personal_details"
	return b
}
func (b *PassportUnspecifiedErrorBuilder) Passport() *PassportUnspecifiedErrorBuilder {
	b.data["type"] = "passport"
	return b
}
func (b *PassportUnspecifiedErrorBuilder) DriverLicense() *PassportUnspecifiedErrorBuilder {
	b.data["type"] = "driver_license"
	return b
}
func (b *PassportUnspecifiedErrorBuilder) IdentityCard() *PassportUnspecifiedErrorBuilder {
	b.data["type"] = "identity_card"
	return b
}
func (b *PassportUnspecifiedErrorBuilder) InternalPassport() *PassportUnspecifiedErrorBuilder {
	b.data["type"] = "internal_passport"
	return b
}
func (b *PassportUnspecifiedErrorBuilder) Address() *PassportUnspecifiedErrorBuilder {
	b.data["type"] = "address"
	return b
}
func (b *PassportUnspecifiedErrorBuilder) UtilityBill() *PassportUnspecifiedErrorBuilder {
	b.data["type"] = "utility_bill"
	return b
}
func (b *PassportUnspecifiedErrorBuilder) BankStatement() *PassportUnspecifiedErrorBuilder {
	b.data["type"] = "bank_statement"
	return b
}
func (b *PassportUnspecifiedErrorBuilder) RentalAgreement() *PassportUnspecifiedErrorBuilder {
	b.data["type"] = "rental_agreement"
	return b
}
func (b *PassportUnspecifiedErrorBuilder) PassportRegistration() *PassportUnspecifiedErrorBuilder {
	b.data["type"] = "passport_registration"
	return b
}
func (b *PassportUnspecifiedErrorBuilder) TemporaryRegistration() *PassportUnspecifiedErrorBuilder {
	b.data["type"] = "temporary_registration"
	return b
}
func (b *PassportUnspecifiedErrorBuilder) PhoneNumber() *PassportUnspecifiedErrorBuilder {
	b.data["type"] = "phone_number"
	return b
}
func (b *PassportUnspecifiedErrorBuilder) Email() *PassportUnspecifiedErrorBuilder {
	b.data["type"] = "email"
	return b
}
func (b *PassportUnspecifiedErrorBuilder) ElementHash(hash string) *PassportUnspecifiedErrorBuilder {
	b.data["element_hash"] = hash
	return b
}
func (b *PassportUnspecifiedErrorBuilder) Message(message string) *PassportUnspecifiedErrorBuilder {
	b.data["message"] = message
	return b
}
func (b *PassportUnspecifiedErrorBuilder) Finish() {
	b.data["source"] = "unspecified"
	b.receiver(b.data)
}

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
}
func (a *SetPassportDataErrorsArgs) UserId(userId int) {
	a.getForm().WithInt("user_id", userId)
}
func (a *SetPassportDataErrorsArgs) DataFieldError() *PassportDataFieldErrorBuilder {
	builder := &PassportDataFieldErrorBuilder{}
	builder.init(a.receiveError)
	return builder
}
func (a *SetPassportDataErrorsArgs) FrontSideError() *PassportFrontSideErrorBuilder {
	builder := &PassportFrontSideErrorBuilder{}
	builder.init(a.receiveError)
	return builder
}
func (a *SetPassportDataErrorsArgs) ReverseSideError() *PassportReverseSideErrorBuilder {
	builder := &PassportReverseSideErrorBuilder{}
	builder.init(a.receiveError)
	return builder
}
func (a *SetPassportDataErrorsArgs) SelfieError() *PassportSelfieErrorBuilder {
	builder := &PassportSelfieErrorBuilder{}
	builder.init(a.receiveError)
	return builder
}
func (a *SetPassportDataErrorsArgs) FileError() *PassportFileErrorBuilder {
	builder := &PassportFileErrorBuilder{}
	builder.init(a.receiveError)
	return builder
}
func (a *SetPassportDataErrorsArgs) TranslationFileError() *PassportTranslationFileErrorBuilder {
	builder := &PassportTranslationFileErrorBuilder{}
	builder.init(a.receiveError)
	return builder
}
func (a *SetPassportDataErrorsArgs) UnspecifiedError() *PassportUnspecifiedErrorBuilder {
	builder := &PassportUnspecifiedErrorBuilder{}
	builder.init(a.receiveError)
	return builder
}
func (a *SetPassportDataErrorsArgs) Archive() (string, io.Reader) {
	a.getForm().WithJson("errors", a.errBuf)
	return a._BasicArgs.Archive()
}
