package parameters

import "github.com/deadblue/telegroid/types"

type GetUserProfilePhotosParams struct {
	implApiParameters
}

func (p *GetUserProfilePhotosParams) UserId(userId int) {
	p.setInt("user_id", userId)
}
func (p *GetUserProfilePhotosParams) Range(offset int, limit int) {
	if offset > 0 {
		p.setInt("offset", offset)
	}
	if limit >= 1 && limit <= 100 {
		p.setInt("limit", limit)
	}
}

type FileParams struct {
	implApiParameters
}

func (p *FileParams) FileId(fileId string) {
	p.set("file_id", fileId)
}

type SetMyCommandsParams struct {
	implApiParameters
}

func (p *SetMyCommandsParams) Commands(command ...types.BotCommand) {
	p.setJson("commands", command)
}
