package types

type ChatType string

func (t ChatType) IsPrivate() bool {
	return t == "private"
}

func (t ChatType) IsGroup() bool {
	return t == "group"
}

func (t ChatType) IsSupergroup() bool {
	return t == "supergroup"
}

func (t ChatType) IsChannel() bool {
	return t == "channel"
}

type EntityType string

func (t EntityType) IsMention() bool {
	return t == "mention"
}

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

func (t EntityType) IsUnderline() bool {
	return t == "underline"
}

func (t EntityType) IsStrikeThrough() bool {
	return t == "strikethrough"
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

type PollType string

func (t PollType) IsRegular() bool {
	return t == "regular"
}

func (t PollType) IsQuiz() bool {
	return t == "quiz"
}

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

type PassportElementType string

func (t PassportElementType) IsPersonalDetails() bool {
	return t == "personal_details"
}

func (t PassportElementType) IsPassport() bool {
	return t == "passport"
}

func (t PassportElementType) IsDriverLicense() bool {
	return t == "driver_license"
}

func (t PassportElementType) IsIdentityCard() bool {
	return t == "identity_card"
}

func (t PassportElementType) IsInternalPassport() bool {
	return t == "internal_passport"
}

func (t PassportElementType) IsAddress() bool {
	return t == "address"
}

func (t PassportElementType) IsUtilityBill() bool {
	return t == "utility_bill"
}

func (t PassportElementType) IsBankStatement() bool {
	return t == "bank_statement"
}

func (t PassportElementType) IsRentalAgreement() bool {
	return t == "rental_agreement"
}

func (t PassportElementType) IsPassportRegistration() bool {
	return t == "passport_registration"
}

func (t PassportElementType) IsTemporaryRegistration() bool {
	return t == "temporary_registration"
}

func (t PassportElementType) IsPhoneNumber() bool {
	return t == "phone_number"
}

func (t PassportElementType) IsEmail() bool {
	return t == "email"
}
