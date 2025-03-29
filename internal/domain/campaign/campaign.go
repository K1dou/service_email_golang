package campaign

import "time"

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

func NewCampaign(id, name, content string, contacts []Contact) *Campaign {
	
	return &Campaign{
		ID:        id,
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}
}