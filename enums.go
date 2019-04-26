package telegroid

// The enum type for "Chat.Type".
// Developer can simply call "IsXXX" method to judge its value.
type ChatType string

func (t ChatType) IsPrivate() bool {
	return t == "private"
}
func (t ChatType) IsGroup() bool {
	return t == "group"
}
func (t ChatType) IsSuperGroup() bool {
	return t == "supergroup"
}
func (t ChatType) IsChannel() bool {
	return t == "channel"
}

// The enum type for "MessageEntity.Type".
// Developer can simply call "IsXXX" method to judge its value.
type EntityType string

func (t EntityType) IsHashtag() bool {
	return t == "hashtag"
}
func (t EntityType) IsCashtag() bool {
	return t == "cashtag"
}
func (t EntityType) IsBotCommand() bool {
	return t == "bot_command"
}
func (t EntityType) IsUrl() bool {
	return t == "url"
}
func (t EntityType) IsEmail() bool {
	return t == "email"
}
func (t EntityType) IsPhoneNumber() bool {
	return t == "phone_number"
}
func (t EntityType) IsBold() bool {
	return t == "bold"
}
func (t EntityType) IsItalic() bool {
	return t == "italic"
}
func (t EntityType) IsCode() bool {
	return t == "code"
}
func (t EntityType) IsPre() bool {
	return t == "pre"
}
func (t EntityType) IsTextLink() bool {
	return t == "text_link"
}
func (t EntityType) IsTextMention() bool {
	return t == "text_mention"
}

// The enum type for "ChangeMember.Status".
// Developer can simply call "IsXXX" method to judge its value.
type MemberStatus string

func (s MemberStatus) IsCreator() bool {
	return s == "creator"
}
func (s MemberStatus) IsAdministrator() bool {
	return s == "administrator"
}
func (s MemberStatus) IsMember() bool {
	return s == "member"
}
func (s MemberStatus) IsRestricted() bool {
	return s == "restricted"
}
func (s MemberStatus) IsLeft() bool {
	return s == "left"
}
func (s MemberStatus) IsKicked() bool {
	return s == "kicked"
}

// The enum type for "PassportData.Type".
// Developer can simply call "IsXXX" method to judge its value.
type PassportType string

func (t PassportType) IsPersonalDetails() bool {
	return t == "personal_details"
}

func (t PassportType) IsPassport() bool {
	return t == "passport"
}

func (t PassportType) IsDriverLicense() bool {
	return t == "driver_license"
}

func (t PassportType) IsIdentityCard() bool {
	return t == "identity_card"
}

func (t PassportType) IsInternalPassport() bool {
	return t == "internal_passport"
}

func (t PassportType) IsAddress() bool {
	return t == "address"
}

func (t PassportType) IsUtilityBill() bool {
	return t == "utility_bill"
}

func (t PassportType) IsBankStatement() bool {
	return t == "bank_statement"
}

func (t PassportType) IsRentalAgreement() bool {
	return t == "rental_agreement"
}

func (t PassportType) IsPassportRegistration() bool {
	return t == "passport_registration"
}

func (t PassportType) IsTemporaryRegistration() bool {
	return t == "temporary_registration"
}

func (t PassportType) IsPhoneNumber() bool {
	return t == "phone_number"
}

func (t PassportType) IsEmail() bool {
	return t == "email"
}
