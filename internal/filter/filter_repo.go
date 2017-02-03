package filter

import (
	//"encoding/json"
	//"fmt"
	//"net/http"
)

type Filter struct {
  minAge string `json:"minage",omitempty`
  maxAge string `json:"maxage",omitempty`
  minDistance string `json:"mindistance",omitempty`
  maxDistance string `json:"maxdistance",omitempty`
  sex string `json:"sex",omitempty`
  userID string `json:"userid",omitempty`
}
