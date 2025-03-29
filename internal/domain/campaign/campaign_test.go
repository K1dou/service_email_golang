package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T){
	assert := assert.New(t)
	name := "Test Campaign"
	content := "This is a test campaign content"	
	contacts := []Contact{{Email: "email@example.com"}, {Email: "user@gmail.com"}}

	campaign := NewCampaign("1", name, content, contacts)

	assert.Equal("1", campaign.ID, "Campaign ID should be '1'")
	assert.Equal(name, campaign.Name, "Campaign name should be 'Test Campaign'")
	assert.Equal(content, campaign.Content, "Campaign content should be 'This is a test campaign content'")
	assert.Equal(len(contacts), len(campaign.Contacts), "Campaign should have 2 contacts")
}