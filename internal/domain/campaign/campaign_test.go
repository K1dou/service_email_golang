package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var(
	name = "Test Campaign"
	content = "This is a test campaign content"
	contacts = []Contact{{Email: "email@example.com"},{Email: "user@gmail.com"}}
)


func Test_NewCampaign_CreatedCampaing(t *testing.T){
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal(name, campaign.Name, "Campaign name should be 'Test Campaign'")
	assert.Equal(content, campaign.Content, "Campaign content should be 'This is a test campaign content'")
	assert.Equal(len(contacts), len(campaign.Contacts), "Campaign should have 2 contacts")
}

func Test_NewCampaing_IDIsNotNill(t *testing.T){
	assert := assert.New(t)

	campaign, _:= NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID, "Campaign ID should not be nil")
}

func Test_NewCampaing_CreatedOnIsNotNill(t *testing.T){
	assert := assert.New(t)	

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.CreatedOn, "Campaign CreatedOn should not be nil")
}

func Test_NewCampaing_MustValidateName(t *testing.T){
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("name is required", err.Error(), "Campaign name should be required")
}

func Test_NewCampaing_MustValidateContent(t *testing.T){
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content is required", err.Error(), "Content should be required")
}

func Test_NewCampaing_MustValidateContacts(t *testing.T){
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []Contact{})

	assert.Equal("contacts is required", err.Error(), "Contacts should be required")
}