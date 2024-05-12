package models

import "time"

type Measurements struct {
  Id        int64       `json:"id"`
  UserId    int64       `json:"user_id"`
  Weight    float64     `json:"weight"`
  Height    float64     `json:"height"`
  BodyFat   float64     `json:"body_fat"`
  CreatedAt time.Time   `json:"created_at"`
  UpdatedAt time.Time   `json:"updated_at"`
}