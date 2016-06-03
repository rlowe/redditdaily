package main

import (
  "errors"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func generatePlacement(placement int) string {
  // Convert to string so we can explode() it
  placementString := strconv.Itoa(placement)
  p := strings.Split(placementString, "")
  last := len(p)-1
  // Special Cases (11, 12, 13)
  if placement > 10 {
    if p[len(p)-2] == "1" {
      if p[last] == "1" {
        return "th"
      } else if p[last] == "2" {
        return "th"
      } else if p[last] == "3" {
        return "th"
      }
    }
  }

  if p[last] == "0" {
    return "th"
  } else if p[last] == "1" {
    return "st"
  } else if p[last] == "2" {
    return "nd"
  } else if p[last] == "3" {
    return "rd"
  }

  return "th"
}

func generateTheRest(myDogsPlace int, lowerBound int, upperBound int) ([]string, error) {
  var theRest []string

  if lowerBound >= upperBound {
    return theRest, errors.New("Invalid Boundaries")
  }
  if !(lowerBound <= myDogsPlace) || !(myDogsPlace <= upperBound) {
    return theRest, errors.New("Invalid Dog Placement")
  }

  for i := lowerBound; i <= upperBound; i++ {
    if i != myDogsPlace {
      theRest = append(theRest, fmt.Sprintf("%d%s", i, generatePlacement(i)))
    }
  }

  return theRest, nil
}

func main() {
  places, err := generateTheRest(1,0,100)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  for _, place := range places {
    fmt.Println(place)
  }
}
