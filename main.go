package main
import "fmt"

type slot struct {
	ID			int
	carNumber 	string
	isEmpty		bool
}


func getID(slot *slot) int {
	return slot.ID
}

func getCarNumber(slot *slot) string {
	return slot.carNumber
}

func getStatus(slot *slot) bool {
	return slot.isEmpty
}

func setID(slot *slot, ID int) {
	slot.ID = ID
}

func setCarNumber(slot *slot, carNumber string) {
	slot.carNumber = carNumber
}

func setStatus(slot *slot, status bool) {
 	slot.isEmpty = status
}



func create(size int) []slot {
	parkingService := make([]slot, size )
	for i := 0; i < size; i++ {
		setID(&parkingService[i], i)
		setCarNumber(&parkingService[i], "")
		setStatus(&parkingService[i], true)
	}
	fmt.Printf("Create parking lot with %v slots \n", size)
	return parkingService
}

func findAllocatedNumber(parkingService []slot) int {
	for i := 0; i < len(parkingService); i++ {
		if parkingService[i].isEmpty {
			return i
		}
	}
	return -1
}

func park(carNumber string, parkingService []slot ) {
	allocatedNumber := findAllocatedNumber(parkingService)
	if allocatedNumber != -1 {
		setStatus(&parkingService[allocatedNumber], false)
		setCarNumber(&parkingService[allocatedNumber], carNumber)
		fmt.Printf("Allocated slot number: %d\n", allocatedNumber + 1)
		return
	}
	fmt.Printf("Sorry, parking lot is full \n")
}

func leave(carNumber string, parkingService []slot, hours float32) {
	var parkingFee float32
	parkingFee = 0
	for i := 0; i < len(parkingService); i++ {
		if getCarNumber(&parkingService[i]) == carNumber {
			
			setStatus(&parkingService[i], true)
			setCarNumber(&parkingService[i], "")
			if(hours < 2) {
				parkingFee += hours * 5
			} else {
				parkingFee = 10 + (hours - 2) * 10
			}
			fmt.Printf("Registration Number %v from slot %v have left with Charitable %v$\n", carNumber, i + 1, parkingFee )
			return
		}
	}
	fmt.Printf("Registration Number %v not found \n", carNumber)
}

func status(parkingService []slot) {
	fmt.Printf("Slot      Number \n")
	for i := 0; i < len(parkingService); i++ {
		fmt.Println(parkingService[i].ID + 1, "       ", parkingService[i].carNumber)
	}
}

func main() {
	parkingService := create(6)
	status(parkingService)
	park("KA-01-HH-1234", parkingService)
	park("KA-01-HH-9999", parkingService)
	park("KA-01-BB-0001", parkingService)
	park("KA-01-HH-7777", parkingService)
	park("KA-01-HH-2701", parkingService)
	park("KA-01-HH-3141", parkingService)
	// park("KA-01-HH-3142", parkingService)
	status(parkingService)
	leave("KA-01-HH-3141", parkingService, 6)
	status(parkingService)
}