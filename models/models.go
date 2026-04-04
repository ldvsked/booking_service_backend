package models


import(
	"time"
	"github.com/google/uuid"
)

var Users[]User = []User{}
var Rooms[]Room = []Room{}
var Bookings[]Booking = []Booking{}
var Schedules[]Schedule = []Schedule{}
var Slots[]Slot = []Slot{}

//сущности - идея того как будут хранится данные о реальном объекте в виде кода
//есть реальный объект, мы придумываем как его хранить(идея) и записываем в виде кода
//сам код - это сущность

func init() {

	Users = append(Users, User{Id: uuid.New(), 
		Email: "vvv@gmail.ru", Role:"user",
		CreatedAt: time.Now(),
	})
	Rooms = append(Rooms, Room{Id: uuid.New(), 
		Name : "Red", 
		Description: "Cuty funny room for relax",
		Capacity: 2,
		CreatedAt: time.Now(),
	})
	
}

type User struct {
	Id uuid.UUID `json:"id"`
	Email string `json:"email"`
	Role string `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
}

type Room struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Capacity int `json:"capacity"`
	CreatedAt time.Time `json:"createdAt"`
}

type Booking struct {
	Id uuid.UUID `json:"id"`
	SlotId uuid.UUID `json:"slotId"`
	UserId uuid.UUID `json:"userId"`
	Status string `json:"status"`
	ConferenceLink string `json:"conferenceLink"`
	CreatedAt time.Time `json:"createdAt"`
}

type Slot struct {
	Id uuid.UUID `json:"id"`
	RoomId uuid.UUID `json:"roomId"`
	Start time.Time `json:"start"`
	End time.Time `json:"end"`
}

type Schedule struct {
	Id uuid.UUID `json:"id"`
	RoomId uuid.UUID `json:"roomId"`
	DaysOfWeek []int `json:"daysOfWeek"`
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
}
