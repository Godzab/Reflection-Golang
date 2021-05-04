package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type bigNumber int

type User struct {
	Name string `json:"name"`
	Age int64 `json:"age"`
}

type City struct{
	Name string `json:"name"`
	Postcode string `json:"postcode"`
	Mayor string `json:"Muyur"`
}

func main(){
	var x float64 = 3.14
	u := User{"Bob", 32}

	fmt.Println(x)
	fmt.Println(u)

	// This is for the float.
	reflectObjValPtr := reflect.ValueOf(&x)
	reflectObjTypePtr := reflect.TypeOf(&x)
	refObjVal := reflectObjValPtr.Elem()
	fmt.Printf("x is of type %s with value %f\n", reflectObjTypePtr, refObjVal.Float())
	fmt.Printf("Can set a new value to the reflection object x: %s\n", refObjVal.CanSet())
	refObjVal.Set(reflect.ValueOf(5.124))
	fmt.Printf("Updated the reflection value of x to: %f\n", x)

	reflectObjValPtr = reflect.ValueOf(u)
	reflectObjTypePtr = reflect.TypeOf(u)
	fmt.Printf("x is of type %s with value %s\n", reflectObjTypePtr, reflectObjValPtr)
	//fmt.Printf("Can set a new value to the reflection object x: %s\n", refObjVal.CanSet())
	//refObjVal.Set(reflect.ValueOf(5.124))
	fmt.Printf("Updated the reflection value of x to: %f\n", u)

	res, err := jsonEncode(u)
	if err != nil{
		panic(err)
	}

	fmt.Println(string(res))

	city := City{"Chicago", "10071", "Ben Miller"}
	res, err = jsonEncode(city)
	if err != nil{
		panic(err)
	}

	fmt.Println(string(res))
}

func jsonEncode(v interface{})([]byte, error){
	refObjectValue := reflect.ValueOf(v)
	refObjectType := reflect.TypeOf(v)
	buffer := bytes.Buffer{}
	if refObjectValue.Kind() != reflect.Struct {
		return buffer.Bytes(), fmt.Errorf("Val of kind %s is not supported!\n", refObjectValue.Kind())
	}
	// iterate over struct fields
	pairs := []string{}
	buffer.WriteString("{")
	for i := 0; i < refObjectValue.NumField(); i++{
		structFieldRefObj := refObjectValue.Field(i)
		structFieldRefObjType := refObjectType.Field(i)
		tag := structFieldRefObjType.Tag.Get("json")
		switch structFieldRefObj.Kind(){
		case reflect.String:
			strVal := structFieldRefObj.Interface().(string)
			// build the pair
			pairs = append(pairs, `"`+string(tag)+`": `+ strVal)
		case reflect.Int64:
			intVal := structFieldRefObj.Interface().(int64)
			pairs = append(pairs, `"`+string(tag)+`": `+ strconv.FormatInt(intVal, 10))
		default:
			return buffer.Bytes(),
				fmt.Errorf("Struct field with name %s and kind %s is not supported!\n",
					structFieldRefObjType.Name, structFieldRefObj.Kind())
		}
	}
	buffer.WriteString(strings.Join(pairs, ","))
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}