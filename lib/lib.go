package lib

import (
  "encoding/json"
"errors")

type Average struct {
    Action  string `json:"action"`
    Average float64 `json:"avg"`
}

var jump_avg float64 = 0
var jump_count float64 = 0
var run_avg float64 = 0
var run_count float64 = 0

func AddAction(s string) error {
  m := make(map[string]interface{})
  err1 := json.Unmarshal([]byte(s), &m)
  if err1 != nil{panic(err1)}
  if (m["action"] == "jump"){
        time := m["time"].(float64)
        p := &jump_avg
        c := &jump_count
        *c = jump_count + 1
	     *p = (time + jump_avg)/jump_count
  }else if (m["action"] == "run"){
    run_time := m["time"].(float64)
    r := &run_avg
    d := &run_count
    *d = run_count + 1
   *r = (run_time + run_avg)/run_count
  }else{return errors.New("Unsupported Action Found")}

  return nil

}

func GetStats() string{
  jump := Average{Action:"jump",Average:jump_avg}
  run := Average{Action:"run",Average:run_avg}

  var av [] Average
  av = append(av, jump)
  av = append(av, run)
  j, _ := json.Marshal(av)
  return(string(j))

}
