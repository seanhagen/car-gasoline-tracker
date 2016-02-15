package main

type Record struct {
	UUID       string
	LocationId string  // location of the gas station
	Odometer   uint32  // odometer reading when gas purchased
	Liters     float32 // amount of gas purchased
	Cost       uint16  // cost in cents
	_saved     bool
	_changed   bool
}

}
