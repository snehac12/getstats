package lib

import(L "../lib"
"testing"
"encoding/json"
//"fmt"
//"reflect"
)

type Score struct {
    Action  string `json:"action"`
    Time float64 `json:"time"`
}

type Average struct {
    Action  string `json:"action"`
    Average float64 `json:"avg"`
}

//Validate AddAction Function returns error
func TestAddActionInput(t *testing.T){
   res := L.AddAction("test")
   if res == nil {t.Errorf("Expected  error and got nil")}
}
//Validate Action other than run and jump returns error
func TestUnsupportedAction(t *testing.T){
  bytes, err := json.Marshal(Score{
        Action:  "skip",
        Time: 200,
    })
  if err != nil{panic(err)}
  res := L.AddAction(string(bytes))
  if res == nil {t.Errorf("Expected Unsupported Action Error and  got nil")}
}

func GetRunAvg(s string)float64{
  var run_avg float64
  var avgs []Average
  json.Unmarshal([]byte(s), &avgs)
  for _, avg := range avgs {
    if(avg.Action == "run"){
      run_avg  = avg.Average
    }
    }
    return run_avg
  }

  func GetJumpAvg(s string)float64{
    var jump_avg float64
    var avgs []Average
    json.Unmarshal([]byte(s), &avgs)
    for _, avg := range avgs {
      if(avg.Action == "jump"){
        jump_avg  = avg.Average
      }
      }
      return jump_avg
    }

//Validate correct value of a run average is calculated
func TestRunAverage(t *testing.T){
  var run_avg float64
  bytes, err := json.Marshal(Score{
        Action:  "run",
        Time: 50,
    })
  if err != nil{panic(err)}
   L.AddAction(string(bytes))
  bytes1, err := json.Marshal(Score{
        Action:  "run",
        Time: 150,
    })
  if err != nil{panic(err)}
  L.AddAction(string(bytes1))
  res := L.GetStats()
  run_avg = GetRunAvg(res)
  if run_avg != 100 {t.Errorf("Expected Average 100 got %f ",run_avg)}
}
//Validate correct value of a run average is calculated
func TestJumpAverage(t *testing.T){
  var jump_avg float64
  bytes, err := json.Marshal(Score{
        Action:  "jump",
        Time: 50,
    })
  if err != nil{panic(err)}
   L.AddAction(string(bytes))
  bytes1, err := json.Marshal(Score{
        Action:  "jump",
        Time: 50,
    })
  if err != nil{panic(err)}
  L.AddAction(string(bytes1))
  res := L.GetStats()
  jump_avg = GetJumpAvg(res)
  if jump_avg != 50 {t.Errorf("Expected Average 50 got %f ",jump_avg)}
}
