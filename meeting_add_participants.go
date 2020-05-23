package zoom

import "fmt"

const MeetingRigtrantsPath = "/meetings/%d/registrants"

type MeetingAddRegistrantsOptions struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	MeetingID int    `json:"meeting_id"`
}

type MeetingParticipant struct {
	ID            int    `json:"id,omitempty"`
	StartTime     string `json:"start_time"`
	Topic         string `json:"topic"`
	RegistrantURL string `json:"registrant_url"`
	JoinURL       string `json:"join_url"`
}

// CreateMeeting calls POST /users/{userId}/meetings
func AddRegistrantMeeting(opts MeetingAddRegistrantsOptions) (MeetingParticipant, error) {
	return defaultClient.AddRegistrantMeeting(opts)
}

// Add Rigtrants calls POST /meetings/%d/registrants
//https://marketplace.zoom.us/docs/api-reference/zoom-api/meetings/meetingregistrantcreate
func (c *Client) AddRegistrantMeeting(opts MeetingAddRegistrantsOptions) (MeetingParticipant, error) {
	var ret = MeetingParticipant{}
	return ret, c.requestV2(requestV2Opts{
		Method:         Post,
		Path:           fmt.Sprintf(MeetingRigtrantsPath, opts.MeetingID),
		DataParameters: &opts,
		Ret:            &ret,
	})
}
