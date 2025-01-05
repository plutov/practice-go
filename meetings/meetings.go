package meetings

type Meeting struct {
	start     int // start time, not an actual timestamp, just an index
	end       int // end time, not an actual timestamp, just an index
	attendees int
}

type ScheduledMeeting struct {
	meeting   Meeting
	roomIndex int
}

func Meetings(meetings []Meeting, rooms []int) []ScheduledMeeting {
	return []ScheduledMeeting{}
}
