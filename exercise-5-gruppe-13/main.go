package main

import "./elevio"
import "fmt"

func main(){

    numFloors := 4

    elevio.Init("localhost:15657", numFloors)

    var d elevio.MotorDirection = elevio.MD_Stop
    elevio.SetMotorDirection(d)
    var dest int
    var last_floor int
    //lamps := make([]elevio.ButtonType, 1)



    drv_buttons := make(chan elevio.ButtonEvent)
    drv_floors  := make(chan int)
    drv_obstr   := make(chan bool)
    drv_stop    := make(chan bool)
    //drv_destination := make(chan int)

    go elevio.PollButtons(drv_buttons/*, drv_destination*/)
    go elevio.PollFloorSensor(drv_floors)
    go elevio.PollObstructionSwitch(drv_obstr)
    go elevio.PollStopButton(drv_stop)

    for {
        select {
        case a := <- drv_buttons:
            fmt.Printf("%+v\n", a)
            elevio.SetButtonLamp(a.Button, a.Floor, true)
            dest = a.Floor
            if dest <  last_floor {
                d = elevio.MD_Down
            } else if dest > last_floor{
                d = elevio.MD_Up
            }
            elevio.SetMotorDirection(d)

/*else if dest == last_floor{
  d = elevio.MD_Stop
}
       /*case dest = <- drv_destination:
          fmt.Printf("%+v\n", dest)*/

       /*(case a := <- drv_floors:
            if a == dest {
              elevio.SetMotorDirection(elevio.MD_Stop)
            }*/

        case a := <- drv_floors:
          last_floor = a
          if dest == a {
            d = elevio.MD_Stop
            elevio.SetButtonLamp(0, a, false)
            elevio.SetButtonLamp(1, a, false)
            elevio.SetButtonLamp(2, a, false)
          }
          elevio.SetMotorDirection(d)

/*  if dest <  a {
      d = elevio.MD_Down
  } else if dest > a{
      d = elevio.MD_Up
  } else */

        case a := <- drv_obstr:
            fmt.Printf("%+v\n", a)
            if a {
                elevio.SetMotorDirection(elevio.MD_Stop)
            } else {
                elevio.SetMotorDirection(d)
            }

        case a := <- drv_stop:
            fmt.Printf("%+v\n", a)
            for f := 0; f < numFloors; f++ {
                for b := elevio.ButtonType(0); b < 3; b++ {
                    elevio.SetButtonLamp(b, f, false)
                }
            }
        }
    }
}
