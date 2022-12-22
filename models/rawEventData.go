package models
// Example:
//
// event: score
// data: {"studentId":"Gretchen.Mayert","exam":13118,"score":0.7895103274760336}


// Path: rawEventData.go func(

type RawData struct {
  StudentId string `json:"studentId"`
  ExamId int `json: "exam"`
  Score float64 `json: "score"`
}

type RawEventData struct {
  Event string `json:"event"`
  Data RawData `json:"data"`
}
