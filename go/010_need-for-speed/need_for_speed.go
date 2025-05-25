package speed

type Car struct {
	battery			int
	batteryDrain	int
	speed			int
	distance		int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:			speed,
		batteryDrain:	batteryDrain,
		battery:		100,
		distance:		0,
	}
}

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery < car.batteryDrain {
		return car
	}

	car.distance += car.speed
	car.battery -= car.batteryDrain
	
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	remainingDistance := track.distance - car.distance
	
	if remainingDistance <= 0 {
		return true
	}
	
	drivesNeeded := (remainingDistance + car.speed - 1) / car.speed 
	batteryNeeded := drivesNeeded * car.batteryDrain
	
	return car.battery >= batteryNeeded
}

func main() {
	speed := 5
	batteryDrain := 2
	distance := 800

	car := NewCar(speed, batteryDrain)
	track := NewTrack(distance)
	car = Drive(car)


	distance = 100
	CanFinish(car, track)
}
