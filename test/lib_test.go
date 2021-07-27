package lib

import(L "../lib"
"testing"
"encoding/json")

type Score struct {
    Action  string `json:"action"`
    Time float64 `json:"time"`
}

func TestAddActionInput(t *testing.T){
   res := L.AddAction("test")
   if res == nil {t.Errorf("Expected  error and got nil")}
}

func TestUnsupportedAction(t *testing.T){
  bytes, err := json.Marshal(Score{
        Action:  "skip",
        Time: 200,
    })
  if err != nil{panic(err)}
  res := L.AddAction(string(bytes))
  if res == nil {t.Errorf("Expected Unsupported Action Error and  got nil")}
}
