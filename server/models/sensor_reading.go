package models

import (
	"time"
)

// Struct to pack results from DB
type SensorReading struct {
	Timestamp   time.Time
	Temperature float32
}
