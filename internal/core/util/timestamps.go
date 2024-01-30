package util

import "time"

type Timestamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Timestamps) InitTimestamps() {
	t.CreatedAt = time.Now()
	t.Update()
}

func (t *Timestamps) Update() {
	t.UpdatedAt = time.Now()
}
