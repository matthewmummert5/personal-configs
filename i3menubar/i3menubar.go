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
    var TimeStamp string
    var count int = 0

    //Get initial battery charge information
    BatteryInfo := getBatteryInfo()

    //Infinite loop. Do this forever
    for {

        //Check battery info once every 5 seconds.
        if count >= 10 {
            BatteryInfo = getBatteryInfo()
            count = 0
        }

        //Check the timestamp twice per second
        TimeStamp = getTimeStamp()

        //print the battery charge information and timestamp
        fmt.Printf("%s | %s\r", BatteryInfo, TimeStamp)

        //Sleep for a half second.
        time.Sleep(500 * time.Millisecond)

        //Increment counter for checking battery charge status
        count++
    } //end infinite for


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

//This function gets the current battery charge status from
// "/sys/class/power_supply/BAT1/uevent"
func getBatteryInfo() string {

    //Read all the bytes of the linux system power supply information file into the variable, b
    b, err := ioutil.ReadFile("/sys/class/power_supply/BAT1/uevent")

    //Check if reading the file failed
    if err != nil {
        return "ERROR WITH BATTERY INFO"
    }

    //convert the raw bytes of the file to a unicode string, s
    s := string(b)

    //split the string up into an array of strings by new line characters
    fields := strings.Split(s, "\n")


    //The battery percentage is in the 12th line of the file (or field 11, counting from zero), starting with the 23rd character
    //There must be a "better" way to do this, but it's functional for now.
    percent, err := strconv.Atoi(fields[11][22:])
    if err != nil {
        return "ERROR WITH BATTERY INFO"
    }

    //Get charging/discharging status from the second line of the file, starting on character 20
    status := fields[1][20:]

    if status == "Charging" {
        status = "(+++)"
    } else if status == "Discharging" {
        status = "(---)"
    } else {
        status = "(***)"
    }


    return fmt.Sprintf("Battery: %03d%% %s", percent, status)
}
