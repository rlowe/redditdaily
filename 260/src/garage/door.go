package garage

import (
  "errors"
  "fmt"
  "io/ioutil"
  "os"
  "strings"
)

const stateFile = "/etc/garage/current.state"

var inputs = []string{
  "button_clicked",
  "cycle_complete",
}

var states = []string{
  "CLOSED",
  "CLOSING",
  "OPEN",
  "OPENING",
  "STOPPED_WHILE_CLOSING",
  "STOPPED_WHILE_OPENING",
}

type Door struct {
  CurrentState string
}

func (d *Door) Click(input string) error {
  // Validate the input
  validInput := false
  for _, i := range inputs {
    if input == i {
      validInput = true
    }
  }
  if !validInput {
    return errors.New("Invalid Input")
  }

  d.CurrentState = d.State()

  // Time to make the doughnuts
  if input == "button_clicked" {
    if d.CurrentState == "CLOSED" {
      d.ChangeState("OPENING")
    } else if d.CurrentState == "CLOSING" {
      d.ChangeState("STOPPED_WHILE_CLOSING")
    } else if d.CurrentState == "OPEN" {
      d.ChangeState("CLOSING")
    } else if d.CurrentState == "OPENING" {
      d.ChangeState("STOPPED_WHILE_OPENING")
    } else if d.CurrentState == "STOPPED_WHILE_CLOSING" {
      d.ChangeState("OPENING")
    } else if d.CurrentState == "STOPPED_WHILE_OPENING" {
      d.ChangeState("CLOSING")
    }
  } else if input == "cycle_complete" {
    if d.CurrentState == "CLOSING" {
      d.ChangeState("CLOSED")
    } else if d.CurrentState == "OPENING" {
      d.ChangeState("OPEN")
    }
  }
  return nil
}

func (d Door) State() string {
  s, _ := ioutil.ReadFile(stateFile)
  d.CurrentState = strings.TrimSpace(string(s))
  return d.CurrentState
}

func (d *Door) ChangeState(newState string) error {

  err := ioutil.WriteFile(stateFile, []byte(newState), os.ModeExclusive)
  if err != nil {
    fmt.Println(err)
    return err
  } else {
    d.CurrentState = newState
    return nil
  }
}
