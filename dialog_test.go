package slack

import (
	"encoding/json"
	"testing"

	"net/http"

	"github.com/stretchr/testify/assert"
)

// Dialogs
var simpleDialog = `{
	"callback_id":"ryde-46e2b0",
	"title":"Request a Ride",
	"submit_label":"Request",
	"notify_on_cancel":true
}`

var simpleTextElement = `{
	"label": "testing label",
	"name": "testing name",
	"type": "text",
	"placeholder": "testing placeholder",
	"optional": true,
	"value": "testing value",
	"max_length": 1000,
	"min_length": 10,
	"hint": "testing hint",
	"subtype": "email"
}`

var simpleSelectElement = `{
	"label": "testing label",
	"name": "testing name",
	"type": "select",
	"placeholder": "testing placeholder",
	"optional": true,
	"value": "testing value",
	"data_source": "users",
	"selected_options": "",
	"options": [{"label": "option 1", "value": "1"}],
	"option_groups": []
}`

func unmarshalDialog() (*Dialog, error) {
	dialog := &Dialog{}
	// Unmarshall the simple dialog json
	if err := json.Unmarshal([]byte(simpleDialog), &dialog); err != nil {
		return nil, err
	}

	// Unmarshall and append the text element
	textElement := &DialogTextElement{}
	if err := json.Unmarshal([]byte(simpleTextElement), &textElement); err != nil {
		return nil, err
	}

	// Unmarshall and append the select element
	selectElement := &DialogSelectElement{}
	if err := json.Unmarshal([]byte(simpleSelectElement), &selectElement); err != nil {
		return nil, err
	}

	dialog.Elements = []DialogElement{
		textElement,
		selectElement,
	}

	return dialog, nil
}

func TestSimpleDialog(t *testing.T) {
	dialog, err := unmarshalDialog()
	assert.Nil(t, err)
	assertSimpleDialog(t, dialog)
}

func TestCreateSimpleDialog(t *testing.T) {
	dialog := &Dialog{}
	dialog.CallbackId = "ryde-46e2b0"
	dialog.Title = "Request a Ride"
	dialog.SubmitLabel = "Request"
	dialog.NotifyOnCancel = true

	textElement := &DialogTextElement{}
	textElement.Label = "testing label"
	textElement.Name = "testing name"
	textElement.Type = "text"
	textElement.Placeholder = "testing placeholder"
	textElement.Optional = true
	textElement.Value = "testing value"
	textElement.MaxLength = 1000
	textElement.MinLength = 10
	textElement.Hint = "testing hint"
	textElement.Subtype = "email"

	selectElement := &DialogSelectElement{}
	selectElement.Label = "testing label"
	selectElement.Name = "testing name"
	selectElement.Type = "select"
	selectElement.Placeholder = "testing placeholder"
	selectElement.Optional = true
	selectElement.Value = "testing value"
	selectElement.DataSource = "users"
	selectElement.SelectedOptions = ""
	selectElement.Options = []DialogElementOption{
		{Label: "option 1", Value: "1"},
	}
	selectElement.OptionGroups = []DialogElementOption{}

	dialog.Elements = []DialogElement{
		textElement,
		selectElement,
	}

	assertSimpleDialog(t, dialog)
}

func assertSimpleDialog(t *testing.T, dialog *Dialog) {
	assert.NotNil(t, dialog)

	// Test the main dialog fields
	assert.Equal(t, "ryde-46e2b0", dialog.CallbackId)
	assert.Equal(t, "Request a Ride", dialog.Title)
	assert.Equal(t, "Request", dialog.SubmitLabel)
	assert.Equal(t, true, dialog.NotifyOnCancel)

	// Test the text element is correctly parsed
	textElement := dialog.Elements[0].(*DialogTextElement)
	assert.Equal(t, "testing label", textElement.Label)
	assert.Equal(t, "testing name", textElement.Name)
	assert.Equal(t, "text", textElement.Type)
	assert.Equal(t, "testing placeholder", textElement.Placeholder)
	assert.Equal(t, true, textElement.Optional)
	assert.Equal(t, "testing value", textElement.Value)
	assert.Equal(t, 1000, textElement.MaxLength)
	assert.Equal(t, 10, textElement.MinLength)
	assert.Equal(t, "testing hint", textElement.Hint)
	assert.Equal(t, "email", textElement.Subtype)

	// Test the text element is correctly parsed
	selectElement := dialog.Elements[1].(*DialogSelectElement)
	assert.Equal(t, "testing label", selectElement.Label)
	assert.Equal(t, "testing name", selectElement.Name)
	assert.Equal(t, "select", selectElement.Type)
	assert.Equal(t, "testing placeholder", selectElement.Placeholder)
	assert.Equal(t, true, selectElement.Optional)
	assert.Equal(t, "testing value", selectElement.Value)
	assert.Equal(t, "users", selectElement.DataSource)
	assert.Equal(t, "", selectElement.SelectedOptions)
	assert.Equal(t, "option 1", selectElement.Options[0].Label)
	assert.Equal(t, "1", selectElement.Options[0].Value)
	assert.Equal(t, 0, len(selectElement.OptionGroups))
}

// Callbacks
var simpleCallback = `{
    "type": "dialog_submission",
    "submission": {
        "name": "Sigourney Dreamweaver",
        "email": "sigdre@example.com",
        "phone": "+1 800-555-1212",
        "meal": "burrito",
        "comment": "No sour cream please",
        "team_channel": "C0LFFBKPB",
        "who_should_sing": "U0MJRG1AL"
    },
    "callback_id": "employee_offsite_1138b",
    "team": {
        "id": "T1ABCD2E12",
        "domain": "coverbands"
    },
    "user": {
        "id": "W12A3BCDEF",
        "name": "dreamweaver"
    },
    "channel": {
        "id": "C1AB2C3DE",
        "name": "coverthon-1999"
    },
    "action_ts": "936893340.702759",
    "token": "M1AqUUw3FqayAbqNtsGMch72",
    "response_url": "https://hooks.slack.com/app/T012AB0A1/123456789/JpmK0yzoZDeRiqfeduTBYXWQ"
}`

func unmarshalCallback(j string) (*DialogCallback, error) {
	callback := &DialogCallback{}
	if err := json.Unmarshal([]byte(j), &callback); err != nil {
		return nil, err
	}
	return callback, nil
}

func TestSimpleCallback(t *testing.T) {
	callback, err := unmarshalCallback(simpleCallback)
	assert.Nil(t, err)
	assertSimpleCallback(t, callback)
}

func assertSimpleCallback(t *testing.T, callback *DialogCallback) {
	assert.NotNil(t, callback)
	assert.Equal(t, "dialog_submission", callback.Type)
	assert.Equal(t, "employee_offsite_1138b", callback.CallbackID)
	assert.Equal(t, "T1ABCD2E12", callback.Team.ID)
	assert.Equal(t, "coverbands", callback.Team.Domain)
	assert.Equal(t, "C1AB2C3DE", callback.Channel.ID)
	assert.Equal(t, "coverthon-1999", callback.Channel.Name)
	assert.Equal(t, "W12A3BCDEF", callback.User.ID)
	assert.Equal(t, "dreamweaver", callback.User.Name)
	assert.Equal(t, "936893340.702759", callback.ActionTs)
	assert.Equal(t, "M1AqUUw3FqayAbqNtsGMch72", callback.Token)
	assert.Equal(t, "https://hooks.slack.com/app/T012AB0A1/123456789/JpmK0yzoZDeRiqfeduTBYXWQ", callback.ResponseURL)
	assert.Equal(t, "Sigourney Dreamweaver", callback.Submission["name"])
	assert.Equal(t, "sigdre@example.com", callback.Submission["email"])
	assert.Equal(t, "+1 800-555-1212", callback.Submission["phone"])
	assert.Equal(t, "burrito", callback.Submission["meal"])
	assert.Equal(t, "No sour cream please", callback.Submission["comment"])
	assert.Equal(t, "C0LFFBKPB", callback.Submission["team_channel"])
	assert.Equal(t, "U0MJRG1AL", callback.Submission["who_should_sing"])
}

// Suggestion Callbacks
var simpleSuggestionCallback = `{
  "type": "dialog_suggestion",
  "token": "W3VDvuzi2nRLsiaDOsmJranO",
  "action_ts": "1528203589.238335",
  "team": {
    "id": "T24BK35ML",
    "domain": "hooli-hq"
  },
  "user": {
    "id": "U900MV5U7",
    "name": "gbelson"
  },
  "channel": {
    "id": "C012AB3CD",
    "name": "triage-platform"
  },
  "name": "external_data",
  "value": "test",
  "callback_id": "bugs"
}`

func unmarshalSuggestionCallback(j string) (*DialogSuggestionCallback, error) {
	callback := &DialogSuggestionCallback{}
	if err := json.Unmarshal([]byte(j), &callback); err != nil {
		return nil, err
	}
	return callback, nil
}

func TestSimpleSuggestionCallback(t *testing.T) {
	callback, err := unmarshalSuggestionCallback(simpleSuggestionCallback)
	assert.Nil(t, err)
	assertSimpleSuggestionCallback(t, callback)
}

func assertSimpleSuggestionCallback(t *testing.T, callback *DialogSuggestionCallback) {
	assert.NotNil(t, callback)
	assert.Equal(t, "dialog_suggestion", callback.Type)
	assert.Equal(t, "W3VDvuzi2nRLsiaDOsmJranO", callback.Token)
	assert.Equal(t, "1528203589.238335", callback.ActionTs)
	assert.Equal(t, "T24BK35ML", callback.Team.ID)
	assert.Equal(t, "hooli-hq", callback.Team.Domain)
	assert.Equal(t, "U900MV5U7", callback.User.ID)
	assert.Equal(t, "gbelson", callback.User.Name)
	assert.Equal(t, "C012AB3CD", callback.Channel.ID)
	assert.Equal(t, "triage-platform", callback.Channel.Name)
	assert.Equal(t, "external_data", callback.ElementName)
	assert.Equal(t, "test", callback.Value)
	assert.Equal(t, "bugs", callback.CallbackID)
}

func openDialogHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(struct {
		SlackResponse
	}{
		SlackResponse: SlackResponse{Ok: true},
	})
	rw.Write(response)
}

func TestOpenDialog(t *testing.T) {
	http.HandleFunc("/dialog.open", openDialogHandler)
	once.Do(startServer)
	SLACK_API = "http://" + serverAddr + "/"
	api := New("testing-token")
	dialog, err := unmarshalDialog()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	err = api.OpenDialog("TXXXXXXXX", *dialog)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	err = api.OpenDialog("", *dialog)
	if err == nil {
		t.Errorf("Did not error with empty trigger", err)
		return
	}
}
