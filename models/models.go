package models


import(
	"time"
	"github.com/google/uuid"
	"fmt"
)

// var Users[]User = []User{}
// var Rooms[]Room = []Room{}
// var Bookings[]Booking = []Booking{}
// var Schedules[]Schedule = []Schedule{}
// var Slots[]Slot = []Slot{}

//сущности - идея того как будут хранится данные о реальном объекте в виде кода
//есть реальный объект, мы придумываем как его хранить(идея) и записываем в виде кода
//сам код - это сущность

type User struct {
	Id uuid.UUID `json:"id" db:"id"`
	Email string `json:"email" db:"email"`
	Role string `json:"role" db:"role"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type Room struct {
	Id uuid.UUID `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Capacity int `json:"capacity" db:"capacity"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type Booking struct {
	Id uuid.UUID `json:"id" db:"id"`
	SlotId uuid.UUID `json:"slotId" db:"slot_id"`
	UserId uuid.UUID `json:"userId" db:"user_id"`
	Status string `json:"status" db:"status"`
	ConferenceLink string `json:"conferenceLink" db:"conference_link"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type Slot struct {
	Id uuid.UUID `json:"id" db:"id"`
	RoomId uuid.UUID `json:"roomId" db:"room_id"`
	Start time.Time `json:"start" db:"start"`
	End time.Time `json:"end" db:"end"`
}

type Schedule struct {
	Id uuid.UUID `json:"id" db:"id"`
	RoomId uuid.UUID `json:"roomId" db:"room_id"`
	DaysOfWeek []int `json:"daysOfWeek" db:"days_of_week"`
	StartTime string `json:"startTime" db:"start_time"`
	EndTime string `json:"endTime" db:"end_time"`
}

func CreateSlots(schedule Schedule) {
	now := time.Now().UTC()
	for i := 0; i < 31; i++ {
		cur := now.AddDate(0, 0, i)
		weekday := int(cur.Weekday())
		if weekday == 0 {
			weekday = 7
		}
		var check bool = false 
		for _, dayInSchedule := range schedule.DaysOfWeek {
			if weekday == dayInSchedule {
				check = true
				break
			}
			if dayInSchedule > weekday {
				break
			}
		}
		if !check { //такого нет в расписании
			continue
		}
		year, month, day := cur.Date()
		for curMIn := ToMinutes(schedule.StartTime); curMIn < ToMinutes(schedule.EndTime); curMIn += 30 {
			startDate := time.Date(year, month, day, curMIn / 60, curMIn % 60, 0, 0, time.UTC)
			endDate := time.Date(year, month, day, (curMIn + 30) / 60, (curMIn + 30) % 60, 0, 0, time.UTC)
			newSlot := Slot{
				Id : uuid.New(), 
				RoomId: schedule.RoomId,
				Start:	startDate, 
				End: endDate,
			}
			Slots = append(Slots, newSlot)
		}
	}
}

func ToMinutes(t string) int {
	var h, m int 
	n,_ := fmt.Sscanf(t, "%d:%d", &h, &m)
	if n != 2 {
		return -1
	}
	if h < 0 || h >= 24 || m < 0 || m >= 60 {
		return -1
	}
	return h * 60 + m;
}