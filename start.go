package main

import (
L "./lib"
"encoding/json"
"fmt"
)


type Score struct {
    Action  string `json:"action"`
    Time float64 `json:"time"`
}

func main (){

  bytes, err := json.Marshal(Score{
        Action:  "jump",
        Time: 200,
    })
    if err != nil {panic(err)}
    err = L.AddAction(string(bytes))
    if err != nil {fmt.Println(err)}

    bytes1, err1 := json.Marshal(Score{
          Action:  "jump",
          Time: 100,
      })
      if err1 != nil {
          panic(err1)
      }
      err = L.AddAction(string(bytes1))
      if err != nil {fmt.Println(err)}

      bytes2, err2 := json.Marshal(Score{
            Action:  "run",
            Time: 75,
        })
        if err2 != nil {
            panic(err2)
        }
        err = L.AddAction(string(bytes2))
        if err != nil {fmt.Println(err)}


        var result string = L.GetStats()
        fmt.Println(result)


}
