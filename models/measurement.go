package models

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type Measurement struct {
	Value float32 `json:"value"`
	Unit  Unit    `json:"unit"`
}

type Unit string

const (
	Gram       Unit = "g"
	Centimeter Unit = "cm"
)

type GramMeasurement struct {
	Measurement
}

func (g *GramMeasurement) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	for k, v := range m {
		switch k {
		case "value":
			g.Value = float32(v.(float64))
		case "unit":
			g.Unit = Unit(v.(string))

		}
	}

	if err := validateStruct(g); err != nil {
		return err
	}

	if g.Unit != Gram {
		return fmt.Errorf("Unit must be %s", Gram)
	}

	return nil
}

func (g GramMeasurement) MarshalJSON() ([]byte, error) {
	if g.Unit != Gram {
		return nil, fmt.Errorf("Unit must be %s", Gram)
	}

	// removes infinite loop
	type Tmp GramMeasurement
	var tmp Tmp = Tmp(g)
	return json.Marshal(tmp)
}

type CentimeterMeasurement struct {
	Measurement
}

func (c *CentimeterMeasurement) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	for k, v := range m {
		switch k {
		case "value":
			c.Value = float32(v.(float64))
		case "unit":
			c.Unit = Unit(v.(string))
		}
	}

	if err := validateStruct(c); err != nil {
		return err
	}

	if c.Unit != Centimeter {
		return fmt.Errorf("Unit must be %s", Centimeter)
	}

	return nil
}

func (c CentimeterMeasurement) MarshalJSON() ([]byte, error) {
	if c.Unit != Centimeter {
		return nil, fmt.Errorf("Unit must be %s", Centimeter)
	}

	// removes infinite loop
	type Tmp CentimeterMeasurement
	var tmp Tmp = Tmp(c)
	return json.Marshal(tmp)
}

// Snippet from https://biscuit.ninja/posts/go-avoid-an-infitine-loop-with-custom-json-unmarshallers/
func validateStruct(item interface{}) error {
	value := reflect.ValueOf(item)

	if value.Kind() == reflect.Ptr && !value.IsNil() {
		value = value.Elem()
	}

	if value.Kind() == reflect.Interface {
		value = reflect.ValueOf(value)
	}

	if value.Kind() != reflect.Struct {
		return fmt.Errorf("value type was %s rather than struct", reflect.TypeOf(value))
	}

	for i := 0; i < value.NumField(); i++ {
		isMandatory, _ := strconv.ParseBool(value.Type().Field(i).Tag.Get("mandatory"))
		isZero := value.Field(i).IsZero()

		if isMandatory && isZero {
			return fmt.Errorf("%s not set when tagged with 'mandatory:\"true\"'", value.Type().Field(i).Name)
		}
	}

	return nil
}
