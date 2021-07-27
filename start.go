package main

import (
L "./lib"
"encoding/json"
"fmt"
)

// Struct for input JSON 
type Score struct {
    Action  string `json:"action"`
    Time float64 `json:"time"`
}

func main (){
 //Send first Json
  bytes, err := json.Marshal(Score{
        Action:  "jump",
        Time: 200,
    })
    if err != nil {panic(err)}
    err = L.AddAction(string(bytes))
    if err != nil {fmt.Println(err)}
    fmt.Println("Sending input :",string(bytes))

    // Send Second Json
    bytes1, err1 := json.Marshal(Score{
          Action:  "jump",
          Time: 100,
      })
      if err1 != nil {
          panic(err1)
      }
      err = L.AddAction(string(bytes1))
      if err != nil {fmt.Println(err)}
      fmt.Println("Sending input :",string(bytes1))

      //Send third Json
      bytes2, err2 := json.Marshal(Score{
            Action:  "run",
            Time: 75,
        })
        if err2 != nil {
            panic(err2)
        }
        err = L.AddAction(string(bytes2))
        if err != nil {fmt.Println(err)}
        fmt.Println("Sending input :",string(bytes2))

        //Get result
        var result string = L.GetStats()
        fmt.Println(result)
}
