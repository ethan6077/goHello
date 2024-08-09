package main

type Listing struct {
	Id              string
	Title           string
	PropertyType    string
	GeneralFeatures GeneralFeatures
}

type GeneralFeatures struct {
	Bedrooms      Bedrooms
	Bathrooms     Bathrooms
	ParkingSpaces ParkingSpaces
}

type Bedrooms struct {
	Value int
}

type Bathrooms struct {
	Value int
}

type ParkingSpaces struct {
	Value int
}
