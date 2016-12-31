package admin

import (
	"fox/model"
	"fox/util/db"
)
type Attachment struct {
	*model.Attachment
}

func NewAttachmentSercice() *Attachment {
	return new(Attachment)
}
func (c *Attachment)GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
	mode := model.NewAttachment()
	data, err := mode.GetAll(q, fields, orderBy, page, 20)
	if err != nil {
		return nil, err
	}
	return data, nil
}