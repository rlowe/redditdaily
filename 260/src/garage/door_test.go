package garage_test

import (
  "fmt"
  "garage"
  "testing"
)

type clickTests struct {
  Input string
  ExpectedOutput string
}

var clickTestExamples = []clickTests{
  { "CLOSED", "OPENING" },
  { "CLOSING", "STOPPED_WHILE_CLOSING" },
  { "OPEN", "CLOSING" },
  { "OPENING", "STOPPED_WHILE_OPENING" },
  { "STOPPED_WHILE_CLOSING", "OPENING" },
  { "STOPPED_WHILE_OPENING", "CLOSING" },
}

var cycleCompleteTestExamples = []clickTests{
  { "CLOSING", "CLOSED" },
  { "OPENING", "OPEN" },
}

func TestClick(t *testing.T) {
  var d = garage.Door{}
  for _, test := range clickTestExamples {
    d.Click("button_clicked")
    if test.ExpectedOutput != d.CurrentState {
      fmt.Println("FAIL")
    }
  }
}
