package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `json:"email"`
}

type Campaign struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedOn time.Time `json:"created_on"`
	Content   string    `json:"content"`
	Contacts  []Contact `json:"contacts"`
}

func NewCampaign(name, content string, contacts []Contact) (*Campaign, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}else if content == "" {
		return nil, errors.New("content is required")
	}else if len(contacts) == 0 {
		return nil, errors.New("contacts is required")}


	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}, nil
}