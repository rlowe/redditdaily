package main

import (
  "errors"
  "fmt"
)

func main() {
  //nodeCount := 3
  //input := [][]int {
  //  { 1, 2 },
  //  { 1, 3 },
  //}
  nodeCount := 16
  input := [][]int {
    { 1, 2 },
    { 1, 3 },
    { 2, 3 },
    { 1, 4 },
    { 3, 4 },
    { 1, 5 },
    { 2, 5 },
    { 1, 6 },
    { 2, 6 },
    { 3, 6 },
    { 3, 7 },
    { 5, 7 },
    { 6, 7 },
    { 3, 8 },
    { 4, 8 },
    { 6, 8 },
    { 7, 8 },
    { 2, 9 },
    { 5, 9 },
    { 6, 9 },
    { 2, 10 },
    { 9, 10 },
    { 6, 11 },
    { 7, 11 },
    { 8, 11 },
    { 9, 11 },
    { 10, 11 },
    { 1, 12 },
    { 6, 12 },
    { 7, 12 },
    { 8, 12 },
    { 11, 12 },
    { 6, 13 },
    { 7, 13 },
    { 9, 13 },
    { 10, 13 },
    { 11, 13 },
    { 5, 14 },
    { 8, 14 },
    { 12, 14 },
    { 13, 14 },
    { 1, 15 },
    { 2, 15 },
    { 5, 15 },
    { 9, 15 },
    { 10, 15 },
    { 11, 15 },
    { 12, 15 },
    { 13, 15 },
    { 1, 16 },
    { 2, 16 },
    { 5, 16 },
    { 6, 16 },
    { 11, 16 },
    { 12, 16 },
    { 13, 16 },
    { 14, 16 },
    { 15, 16 },
  }

  calculatedDegrees, err := calculateDegrees(nodeCount, input)
  if err != nil {
    fmt.Println(err)
  } else {
    for k, v := range calculatedDegrees {
      fmt.Println("Node", k, "has a degree of", v)
    }
  }
}

func calculateDegrees(nodeCount int, eList [][]int) (map[int]int, error) {
  ret := make(map[int]int)
  for _, nodes := range eList {
    if _, ok := ret[nodes[0]]; !ok {
      ret[nodes[0]] = 1
    } else {
      ret[nodes[0]]++
    }
    if _, ok := ret[nodes[1]]; !ok {
      ret[nodes[1]] = 1
    } else {
      ret[nodes[1]]++
    }
  }

  if len(ret) != nodeCount {
    return ret, errors.New("Node Count incorrect")
  }

  return ret, nil
}