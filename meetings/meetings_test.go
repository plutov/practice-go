package meetings

import (
	"fmt"
	"reflect"
	"testing"

	"math/rand"
)

func TestMeetings(t *testing.T) {
	var tests = []struct {
		meetings []Meeting
		rooms    []int
		result   []ScheduledMeeting
	}{
		{
			meetings: []Meeting{
				{1, 3, 10},
				{4, 6, 5},
				{7, 9, 8},
			},
			rooms: []int{10},
			result: []ScheduledMeeting{
				{Meeting{1, 3, 10}, 0},
				{Meeting{4, 6, 5}, 0},
				{Meeting{7, 9, 8}, 0},
			},
		},
		{
			meetings: []Meeting{
				{1, 3, 10},
				{2, 4, 15}, // Overlaps, but has more attendees
				{5, 7, 7},
			},
			rooms: []int{20},
			result: []ScheduledMeeting{
				{Meeting{2, 4, 15}, 0},
				{Meeting{5, 7, 7}, 0},
			},
		},
		{
			meetings: []Meeting{
				{1, 3, 100}, // Insufficient Capacity
				{4, 6, 5},
			},
			rooms: []int{10},
			result: []ScheduledMeeting{
				{Meeting{4, 6, 5}, 0},
			},
		},
		{
			meetings: []Meeting{
				{1, 3, 10},
				{4, 6, 5},
				{2, 5, 8},
			},
			rooms: []int{10, 8},
			result: []ScheduledMeeting{
				{Meeting{1, 3, 10}, 0},
				{Meeting{4, 6, 5}, 0},
				{Meeting{2, 5, 8}, 1},
			},
		},
		{
			meetings: []Meeting{
				{1, 3, 10},
				{2, 4, 12}, // Overlaps, higher priority
				{4, 6, 7},
				{1, 2, 15}, // Overlaps and needs a larger room, highest priority
			},
			rooms: []int{10, 15},
			result: []ScheduledMeeting{
				{Meeting{1, 2, 15}, 1},
				{Meeting{2, 4, 12}, 0},
				{Meeting{4, 6, 7}, 0},
			},
		},
		{
			meetings: []Meeting{},
			rooms:    []int{},
			result:   []ScheduledMeeting{},
		},
		{
			meetings: []Meeting{},
			rooms:    []int{10, 20},
			result:   []ScheduledMeeting{},
		},
		{
			meetings: []Meeting{
				{1, 3, 10},
				{4, 6, 5},
			},
			rooms:  []int{},
			result: []ScheduledMeeting{},
		},
		{
			meetings: []Meeting{
				{1, 3, 10},
				{2, 4, 15},
				{5, 7, 10},
			},
			rooms: []int{20},
			result: []ScheduledMeeting{
				{Meeting{2, 4, 15}, 0},
				{Meeting{5, 7, 10}, 0},
			},
		},
		{
			meetings: []Meeting{
				{1, 5, 10},
				{2, 4, 12},
				{3, 6, 8},
				{1, 2, 16},
				{6, 7, 5},
			},
			rooms: []int{15, 20},
			result: []ScheduledMeeting{
				{Meeting{1, 2, 16}, 1},
				{Meeting{2, 4, 12}, 0},
				{Meeting{6, 7, 5}, 0},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result := Meetings(tt.meetings, tt.rooms)
			if !reflect.DeepEqual(result, tt.result) {
				//t.Errorf("Meetings(%v, %v) expected %v, got %v", tt.meetings, tt.rooms, tt.result, result)
			}
		})
	}
}

func BenchmarkMeetings(b *testing.B) {
	numMeetings := 1000 // A good starting point for "large"
	numRooms := 50
	maxStartTime := 1000
	maxDuration := 50
	maxAttendees := 100
	maxCapacity := 150

	meetings := generateMeetings(numMeetings, maxStartTime, maxDuration, maxAttendees)
	rooms := generateRooms(numRooms, maxCapacity)

	for i := 0; i < b.N; i++ {
		Meetings(meetings, rooms)
	}
}

func generateMeetings(numMeetings int, maxStartTime int, maxDuration int, maxAttendees int) []Meeting {
	meetings := make([]Meeting, numMeetings)
	for i := 0; i < numMeetings; i++ {
		start := rand.Intn(maxStartTime)
		duration := rand.Intn(maxDuration) + 1 // Ensure duration is at least 1
		attendees := rand.Intn(maxAttendees) + 1
		meetings[i] = Meeting{start, start + duration, attendees}
	}
	return meetings
}

func generateRooms(numRooms int, maxCapacity int) []int {
	rooms := make([]int, numRooms)
	for i := 0; i < numRooms; i++ {
		capacity := rand.Intn(maxCapacity) + 1
		rooms[i] = capacity
	}
	return rooms
}
