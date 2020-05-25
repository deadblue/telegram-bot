package types

type User struct {
	Id                       int    `json:"id"`
	IsBot                    bool   `json:"is_bot"`
	FirstName                string `json:"first_name"`
	LastName                 string `json:"last_name"`
	Username                 string `json:"username"`
	LanguageCode             string `json:"language_code"`
	CanJoinGroups            bool   `json:"can_join_groups"`
	CallReadAllGroupMessages bool   `json:"call_read_all_group_messages"`
	SupportsInlineQueries    bool   `json:"supports_inline_queries"`
}

type UserProfilePhotos struct {
	TotalCount int            `json:"total_count"`
	Photos     [][]*PhotoSize `json:"photos"`
}
