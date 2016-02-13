package main

type Record struct {
	GUID       string
	LocationId string  // location of the gas station
	Odometer   uint32  // odometer reading when gas purchased
	Liters     float32 // amount of gas purchased
	Cost       uint16  // cost in cents
}
