package loader

import (
	"reflect"
	"strings"

	"github.com/gotk3/gotk3/gtk"
)

type Ptr interface {
	Val() interface{}
}

//Map map widgets from glade file to struct's Wd*
func Map(ptr Ptr, builder *gtk.Builder) {
	ptrType := reflect.TypeOf(ptr.Val())
	ptrRef := reflect.ValueOf(ptr)

	// range field
	for i := 0; i < ptrType.NumField(); i++ {
		fieldName := ptrType.Field(i).Name

		//name start with "Wd"
		if strings.Index(fieldName, "Wd") == 0 {
			obj, err := builder.GetObject(fieldName)

			//get object from builder is no error
			if err == nil {
				fieldVal := reflect.ValueOf(obj).Convert(ptrType.Field(i).Type) // Convert type
				ptrRef.Elem().FieldByName(fieldName).Set(fieldVal)              //set value
			}
		}
	}
}

//Connect Connect signal from glade file to struct's Sig*
func Connect(ptr Ptr, builder *gtk.Builder) {
	ptrType := reflect.TypeOf(ptr)
	ptrRef := reflect.ValueOf(ptr)

	sigs := make(map[string]interface{})
	// range Method
	for i := 0; i < ptrType.NumMethod(); i++ {
		methodName := ptrType.Method(i).Name

		//name start with "Sig"
		if strings.Index(methodName, "Sig") == 0 {
			sigs[methodName] = ptrRef.MethodByName(methodName).Interface()
		}
	}
	builder.ConnectSignals(sigs)
}
