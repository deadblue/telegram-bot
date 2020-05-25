package types

type Poll struct {
	Id                    string           `json:"id"`
	Question              string           `json:"question"`
	Options               []*PollOption    `json:"options"`
	TotalVoterCount       int              `json:"total_voter_count"`
	IsClosed              bool             `json:"is_closed"`
	IsAnonymous           bool             `json:"is_anonymous"`
	Type                  PollType         `json:"type"`
	AllowsMultipleAnswers bool             `json:"allows_multiple_answers"`
	CorrectOptionId       int              `json:"correct_option_id"`
	Explanation           string           `json:"explanation"`
	ExplanationEntities   []*MessageEntity `json:"explanation_entities"`
	OpenPeriod            int              `json:"open_period"`
	CloseDate             int              `json:"close_date"`
}

type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

type PollAnswer struct {
	PollId    string `json:"poll_id"`
	User      *User  `json:"user"`
	OptionIds []int  `json:"option_ids"`
}
