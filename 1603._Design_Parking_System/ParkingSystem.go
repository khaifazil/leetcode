package main

import "fmt"

func main() {
	carpark := Constructor(1, 1, 1)

	fmt.Println(carpark.AddCar(1))
	fmt.Println(carpark.AddCar(1))
}

//type ParkingSystem struct {
//	big    int
//	medium int
//	small  int
//}
//
//func Constructor(big int, medium int, small int) ParkingSystem {
//	return ParkingSystem{
//		big:    big,
//		medium: medium,
//		small:  small,
//	}
//}
//
//func (this *ParkingSystem) AddCar(carType int) bool {
//	switch carType {
//	case 1:
//		if this.big == 0 {
//			return false
//		}
//		this.big--
//	case 2:
//		if this.medium == 0 {
//			return false
//		}
//		this.medium--
//	case 3:
//		if this.small == 0 {
//			return false
//		}
//		this.small--
//	}
//	return true
//}

type ParkingSystem struct {
	CarType map[int]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		CarType: map[int]int{1: big, 2: medium, 3: small},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.CarType[carType] > 0 {
		this.CarType[carType] -= 1
		return true
	}

	return false
}
