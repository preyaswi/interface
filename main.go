package main

import (
	"fmt"
	"sort"
)

// Flight - a struct that
// contains information about flights
type Flight struct {
  Origin string
  Destination string
  Price int
}

// SortByPrice sorts flights from lowest to  highest
// func SortByPrice(flights []Flight) []Flight {
// 	for i:=0;i<len(flights);i++{
// 		for j:=i;j<len(flights)-i-1;j++{
// 			if flights[j].Price>flights[j+1].Price{
// 				flights[j],flights[j+1]=flights[j+1],flights[j]
// 			}
// 		}
// 	}
//   // implement
//   return flights
// }


// SortByPrice sorts flights from highest to lowest
type ByAge []Flight
func (a ByAge) Less(i, j int) bool {
	 return a[i].Price > a[j].Price
	 }
func(a ByAge) Swap(i, j int)      {
	 a[i], a[j] = a[j], a[i]
	 }
func (a ByAge) Len() int { 
		return len(a) 
	}

func SortByPrice(flights []Flight) []Flight {
sort.Sort(ByAge(flights))
  // implement
  return flights
}


func printFlights(flights []Flight) {
  for _, flight := range flights {
    fmt.Printf("Origin: %s, Destination: %s, Price: %d", flight.Origin, flight.Destination, flight.Price)
	fmt.Println()
  }

}
func main() {
  // an empty slice of flights
  var flights []Flight
  flights=append(flights,Flight{Origin: "kerala",Destination: "mumbai",Price: 100})
  flights = append(flights, Flight{Origin: "kannur",Destination: "kochi",Price: 10})
  flights = append(flights, Flight{Origin: "kollam",Destination: "kottayam",Price: 200})
  flights = append(flights, Flight{Origin: "kollam",Destination: "kottayam",Price: 34})
  flights = append(flights, Flight{Origin: "kollam",Destination: "kottayam",Price: 56})
  flights = append(flights, Flight{Origin: "kollam",Destination: "kottayam",Price: 23})

  sortedList := SortByPrice(flights)
  printFlights(sortedList)
}