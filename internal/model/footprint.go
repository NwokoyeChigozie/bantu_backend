package model

import (	
	"time"
)

type Footprint struct {

	CurrentTime     time.Time 	
    IPAddress		string 		
    DeviceInfo      string		
	BrowserType     string		
	Location       struct {
		Longitude uint `json:"longitude"`
		Latitude uint `json:"latitude"`
		City    string `json:"city"`
		Country string `json:"country"`
	}
}

