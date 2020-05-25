package types

type PassportFile struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FileDate     int    `json:"file_date"`
}

type EncryptedPassportElement struct {
	Type        PassportElementType `json:"type"`
	Data        string              `json:"data"`
	PhoneNumber string              `json:"phone_number"`
	Email       string              `json:"email"`
	Files       []*PassportFile     `json:"files"`
	FrontSide   *PassportFile       `json:"front_side"`
	ReverseSide *PassportFile       `json:"reverse_side"`
	Selfie      *PassportFile       `json:"selfie"`
	Translation []*PassportFile     `json:"translation"`
	Hash        string              `json:"hash"`
}

type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

type PassportData struct {
	Data        []*EncryptedPassportElement `json:"data"`
	Credentials *EncryptedCredentials       `json:"credentials"`
}
