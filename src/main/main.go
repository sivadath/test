package main

import (
	"fmt"
	"net/http"
	"parking/pkg/parkinglot"
	"parking/pkg/slot"
)

type user struct {
	Vehicle     string
	RequestType string
	ParkingType parkinglot.Parkinglot
}

func (u *user) Unmarshal(r *http.Request) {

}

func (u *user) findSlot(w http.ResponseWriter, r *http.Request) {
	u.Unmarshal(r)
	slots := u.ParkingType.GetASlot(slot.Vehicle(u.Vehicle))
	w.Write([]byte(fmt.Sprintf("available slots:%v", slots)))
}

func main() {
	u := user{}

	http.HandleFunc("/findSlot", u.findSlot)
	http.ListenAndServe("http:127.0.0.1:8080", nil)
}
