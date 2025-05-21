package utils

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type DateOnly time.Time

func (d *DateOnly) UnmarshalJSON(b []byte) error {
	fmt.Println("Parsing date:", string(b))
	s := strings.Trim(string(b), `"`)
	if s == "" {
		*d = DateOnly(time.Time{}) // empty string becomes zero time
		return nil
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*d = DateOnly(t)
	return nil
}

func (d DateOnly) MarshalJSON() ([]byte, error) {
	fmt.Println("MARSHAL JSON TIME ", d)

	t := time.Time(d)
	return json.Marshal(t.Format("2006-01-02"))
}

func (d DateOnly) ToTime() time.Time {
	return time.Time(d)
}

func (d DateOnly) MarshalBSONValue() (bsontype.Type, []byte, error) {
	t := time.Time(d)
	return bson.MarshalValue(t)
}

func (d *DateOnly) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	var tt time.Time
	err := bson.UnmarshalValue(t, data, &tt)
	if err != nil {
		return err
	}
	*d = DateOnly(tt)
	return nil
}
