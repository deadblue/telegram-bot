package types

type Chat struct {
	Id               int64            `json:"id"`
	Type             ChatType         `json:"type"`
	Title            string           `json:"title"`
	Username         string           `json:"username"`
	FirstName        string           `json:"first_name"`
	LastName         string           `json:"last_name"`
	Photo            *ChatPhoto       `json:"photo"`
	Description      string           `json:"description"`
	InviteLink       string           `json:"invite_link"`
	PinnedMessage    *Message         `json:"pinned_message"`
	Permissions      *ChatPermissions `json:"permissions"`
	SlowModeDelay    int              `json:"slow_mode_delay"`
	StickerSetName   string           `json:"sticker_set_name"`
	CanSetStickerSet bool             `json:"can_set_sticker_set"`
}

type ChatPhoto struct {
	SmallFileId       string `json:"small_file_id"`
	SmallFileUniqueId string `json:"small_file_unique_id"`
	BigFileId         string `json:"big_file_id"`
	BigFileUniqueId   string `json:"big_file_unique_id"`
}

type ChatMember struct {
	User                  *User        `json:"user"`
	Status                MemberStatus `json:"status"`
	CustomTitle           string       `json:"custom_title"`
	UntilDate             int          `json:"until_date"`
	CanBeEdited           bool         `json:"can_be_edited"`
	CanPostMessages       bool         `json:"can_post_messages"`
	CanEditMessages       bool         `json:"can_edit_messages"`
	CanDeleteMessages     bool         `json:"can_delete_messages"`
	CanRestrictMembers    bool         `json:"can_restrict_members"`
	CanPromoteMembers     bool         `json:"can_promote_members"`
	CanChangeInfo         bool         `json:"can_change_info"`
	CanInviteUsers        bool         `json:"can_invite_users"`
	CanPinMessages        bool         `json:"can_pin_messages"`
	IsMember              bool         `json:"is_member"`
	CanSendMessages       bool         `json:"can_send_messages"`
	CanSendMediaMessages  bool         `json:"can_send_media_messages"`
	CanSendPolls          bool         `json:"can_send_polls"`
	CanSendOtherMessages  bool         `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool         `json:"can_add_web_page_previews"`
}

type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages"`
	CanSendMediaMessages  bool `json:"can_send_media_messages"`
	CanSendPolls          bool `json:"can_send_polls"`
	CanSendOtherMessages  bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	CanChangeInfo         bool `json:"can_change_info"`
	CanInviteUsers        bool `json:"can_invite_users"`
	CanPinMessages        bool `json:"can_pin_messages"`
}
