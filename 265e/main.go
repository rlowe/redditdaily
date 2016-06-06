package main

import (
  "bufio"
  "encoding/csv"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

type President struct {
  Name       string
  DoB        string
  YoB        int
  BirthPlace string
  DoD        string
  YoD        int
  LoD        string
}

func main() {
  var presidents []President
  minYear := 2016
  curYear := 2016
  var yearAlive = make(map[int]int)

  // Load the file
  f, err := os.Open("data.txt")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  r := csv.NewReader(bufio.NewReader(f))
  for {
    record, err := r.Read()

    // Stop at EOF
    if err == io.EOF {
      break
    }

    // Generate YoB
    yoba := strings.Split(strings.TrimSpace(record[1])," ")
    yob, _ := strconv.Atoi(yoba[2])
    if yob < minYear {
      minYear = yob
    }

    yod := 0
    if len(strings.TrimSpace(record[3])) > 1 {
      yoda := strings.Split(strings.TrimSpace(record[3])," ")
      yod, _ = strconv.Atoi(yoda[2])
    }

    presidents = append(presidents, President{record[0],record[1],yob,record[2],record[3],yod,record[4]})
  }

  for year := minYear; year <= curYear; year++ {
    for _, prez := range presidents {
      if prez.YoB <= year {
        if prez.YoD == 0 || prez.YoD >= year {
          yearAlive[year]++
        }
      }
    }
  }

  theMost := 0
  theMostYear := 0
  for k, v := range yearAlive {
    if v > theMost {
      theMost = v
      theMostYear = k
    }
  }

  fmt.Println("The Year", theMostYear, "had the most presidents at", yearAlive[theMostYear])
}
