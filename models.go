package main


import "time"

var Accounts[]Account = []Account{}
var Rooms[]Room = []Room{}
var Bookings[]Booking = []Booking{}

//сущности - идея того как будут хранится данные о реальном объекте в виде кода
//есть реальный объект, мы придумываем как его хранить(идея) и записываем в виде кода
//сам код - это сущность

func init() {
	Accounts = append(Accounts, Account{Id: 1, 
		Email: "vvv@gmail.ru", Role:"user",
		CreatedAt: time.Now(),
	})
	Rooms = append(Rooms, Room{Id: 1, 
		Name : "Red",
	})
	
}

type Account struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Role string `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type Room struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Capacity int `json:"capacity"`
}

type Booking struct {
	AccountId int `json:"account_id"`
	RoomId int `json:"room_id"`
	Status string `json:"status"`
}