### Optimal Meeting Scheduling

Given a set of meetings with start and end times, and a set of rooms with different capacities, find the optimal schedule that maximizes the number of meetings held. A meeting can only be held in a room if its capacity is greater than or equal to the number of attendees.

So the constraints are:
- Meetings cannot overlap in the same room.
- A meeting can only be scheduled in a room with sufficient capacity.
- In case of a conflict - prioritize the meeting with the higher number of attendees

Input/Output structs:

```go
type Meeting struct {
	start     int // start time, not an actual timestamp, just an index
	end       int // end time, not an actual timestamp, just an index
	attendees int
}

type ScheduledMeeting struct {
	meeting   Meeting
	roomIndex int
}
```

Implement the following function. The results must be sorted as they appear in `meetings` input.

```go
func Meetings(meetings []Meeting, rooms []int) []ScheduledMeeting
```

### Run tests with benchmarks

```
go test -bench .
```
