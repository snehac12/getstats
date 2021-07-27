package lib

import (
  "encoding/json"
"errors")
//struct for result json
type Average struct {
    Action  string `json:"action"`
    Average float64 `json:"avg"`
}


// Initialize the averages and counts
var jump_avg float64 = 0
var jump_count float64 = 0
var run_avg float64 = 0
var run_count float64 = 0


// Function will accept serialized json and maintain the average time of actions
func AddAction(s string) error {
  m := make(map[string]interface{})
  err1 := json.Unmarshal([]byte(s), &m)
  if err1 != nil{return err1}
  if (m["action"] == "jump"){
        time := m["time"].(float64)
        p := &jump_avg
        c := &jump_count
        *c = jump_count + 1                   //persist count
	     *p = (time + jump_avg)/jump_count      //persist jump_avg
  }else if (m["action"] == "run"){
    run_time := m["time"].(float64)
    r := &run_avg
    d := &run_count
    *d = run_count + 1                       //persist count
   *r = (run_time + run_avg)/run_count       // persist run_avg
  }else{return errors.New("Unsupported Action Found")}//handle Unsupported actions

  return nil

}

func GetStats() string{
  jump := Average{Action:"jump",Average:jump_avg}
  run := Average{Action:"run",Average:run_avg}

  var av [] Average                                  //slice of struct
  av = append(av, jump)
  av = append(av, run)
  res, err := json.Marshal(av)                            //json serialized
  if err != nil {panic(err)}
  return(string(res))

}
