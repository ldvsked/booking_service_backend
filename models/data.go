package models

import (
	"github.com/google/uuid"
	"time"
)

func init() {
    testUser := User{
        Id:        uuid.New(),
        Email:     "student@university.edu",
        Role:      "user",
        CreatedAt: time.Now(),
    }
    Users = append(Users, testUser)

    roomRed := Room{
        Id:          uuid.New(),
        Name:        "Red",
        Description: "Funny cute room for rest",
        Capacity:    2,
        CreatedAt:   time.Now(),
    }
    roomBlue := Room{
        Id:          uuid.New(),
        Name:        "Blue",
        Description: "Like you",
        Capacity:    50,
        CreatedAt:   time.Now(),
    }
    Rooms = append(Rooms, roomRed, roomBlue)

    scheduleRed := Schedule{
        Id:         uuid.New(),
        RoomId:     roomRed.Id,
        DaysOfWeek: []int{1, 3, 5},
        StartTime:  "09:00",
        EndTime:    "11:00",
    }
    scheduleBlue := Schedule{
        Id:         uuid.New(),
        RoomId:     roomBlue.Id,
        DaysOfWeek: []int{2, 4},
        StartTime:  "14:00",
        EndTime:    "16:00",
    }
    Schedules = append(Schedules, scheduleRed, scheduleBlue)

    CreateSlots(scheduleRed)
	CreateSlots(scheduleBlue)
	
    if len(Slots) >= 5 {
        for i := 0; i < 5; i++ {
            newBooking := Booking{
                Id:             uuid.New(),
                SlotId:         Slots[i].Id, // Привязка к слоту
                UserId:         testUser.Id, // Привязка к юзеру
                Status:         "active",
                ConferenceLink: "https://meet.jit.si/room-" + uuid.New().String()[:8],
                CreatedAt:      time.Now(),
            }
            Bookings = append(Bookings, newBooking)
        }
    }
}