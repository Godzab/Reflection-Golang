# Reflection-Golang

__REFLECTION__

**SUMMARY**
- This is a form of meta programming
- This can be viewed as the ability of a program to introspect its own structures
    - Its a form of metaprogramming
    - The ability of a program to manipulate, read and analyse other programs or itself.
    - Ability to change its behaviour at **runtime**

All programs are a combination of **data** and **instructions** and reflection gives us the ability to inspect the data that the program uses.
Each variable in Go is an _interface_ value which contains the value and the type of the value.

- We can also use type aliases where we reassign types
```go
type bigNumber int
```

- The golang language cannot know those types in advance so we have a general purpose objects that contain the reflection values and types
- these are reflection types, they are an abstracted layer on top of every possible type in the language.
- The first law of reflection is
    - You can go from interface value to reflection object
    - You can go from reflection object to interface value.
    - We can check if the value can be set by using `reflectVal.Canset()`
    - Not all objects can be settable so you always need to check.


```go
var x float64 = 3.14
reflectObjVal := reflect.ValueOf(x)
reflectObjType := reflect.TypeOf(x)
```
_**NB:** If we pass by value, the reflection value will not be able to be set_
```go
var x float64 = 3.14
reflectObjVal := reflect.ValueOf(&x)
reflectObjType := reflect.TypeOf(&x)

//Get pointer value
refObjVal := reflectObjValPtr.Elem()
//Check if we can set a value 
refObjVal.CanSet()
```

We can dereferrence a pointer to its value by using the folowing
```go
refObjVal := reflectObjValPtr.Elem()
//Same as
refObjVal := *reflectObjValPtr
```

We can now use the set method to reassign a value to a reflection value.
```go
refObjVal.Set(reflect.ValueOf(5.124))
```
Set assigns x to the value v. It panics if CanSet returns false. As in Go,
x's value must be assignable to v's type. We can check the **kind** of a reflection item.

```go
reflectObjType.Kind()
```
Custom tags can be defined and used acccordingly from the following manner
```go
type City struct{
Name string `json:"name"`
Postcode string `json:"postcode"`
Mayor string `json:"mayor"`
}

tag := structFieldRefObjType.Tag.Get("json") // Gets the json tag - Very handy
```