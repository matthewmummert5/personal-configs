/*
*    Copyright (C) Matthew Mummert, 2016
*
*    This program is licensed under the GNU Public License V3
*    https://www.gnu.org/licenses/gpl-3.0.en.html
*/

package main

import (
    "fmt"
    "time"
    "io/ioutil"
    "strings"
    "strconv"
)

func main() {

    for true {
        fmt.Printf("\r%s | %s", getBatteryInfo(), getTimeStamp())
        time.Sleep(500 * time.Millisecond)
    }


}

//This function gets a timestamp for the current time
func getTimeStamp() string {
    //get current time
    currentTime := time.Now()

    //Current Time
    hour    := currentTime.Hour()
    minute  := currentTime.Minute()
    second  := currentTime.Second()

    //Current Date
    year    := currentTime.Year()
    month   := currentTime.Month()
    day     := currentTime.Day()

    return fmt.Sprintf("%02d-%02d-%04d | %02d:%02d:%02d ", month, day, year, hour, minute, second)
}

//This function gets the current battery charge % from
// "/sys/class/power_supply/BAT1/uevent"
func getBatteryInfo() string {

    b, err := ioutil.ReadFile("/sys/class/power_supply/BAT1/uevent")
    if err != nil {
        return "ERROR WITH BATTERY INFO"
    }

    //convert the raw bytes of the file to a unicode string
    s := string(b)

    //split the string up into seperate fields by new line characters
    fields := strings.Split(s, "\n")


    //The battery percentage is in the 12th line of the file (or field 11, counting from zero), starting with the 23rd character
    //There must be a "better" way to do this, but it's functional for now.
    percent, err := strconv.Atoi(fields[11][22:])
    if err != nil {
        return "ERROR WITH BATTERY INFO"
    }

    //Get charging/discharging status
    status := fields[1][20:]

    if status == "Charging" {
        status = "(+++)"
    } else  if status == "Discharging" {
        status = "(---)"
    } else {
        status = "(***)"
    }


    return fmt.Sprintf("Battery: %03d%% %s", percent, status)
}
