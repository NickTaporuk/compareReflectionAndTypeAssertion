package interfaceToMap

import (
	"testing"
)

type T struct {
	Ext *interface{}
}

func BenchmarkConvertInterfaceToMapStrings(b *testing.B) {
	mpIface := make(map[string]interface{})
	mpIface["test1"] = "test1"
	mpIface["test2"] = "test2"

	var ext interface{} = mpIface

	t := &T{
		Ext: &ext,
	}
	var mp = make(map[string]string)

	for n := 0; n < b.N; n++ {
		ConvertInterfaceToMapStrings(t.Ext, mp)
	}
}

func BenchmarkInterfaceToMap(b *testing.B) {
	mpIface := make(map[string]interface{})
	mpIface["test1"] = "test1"
	mpIface["test2"] = "test2"

	var ext interface{} = mpIface
	var (
		convertToMap = make(map[string]string)
	)
	t := &T{
		Ext: &ext,
	}

	for n := 0; n < b.N; n++ {
		intf := InterfaceToMap(t.Ext)

		for k, ex := range intf {
			convertToMap[k] = ex.(string)
		}
	}
}


func BenchmarkConvertInterfaceToMapStrings2(b *testing.B) {
	mpIface := make(map[string]interface{})
	mpIface["test1"] = "test1"
	mpIface["test2"] = "test2"

	var ext interface{} = mpIface

	t := &T{
		Ext: &ext,
	}

	for n := 0; n < b.N; n++ {
		ConvertInterfaceToMapStrings2(t.Ext)
	}
}


func BenchmarkInterfaceToMap2(b *testing.B) {
	mpIface := make(map[string]interface{})
	mpIface["test1"] = "test1"
	mpIface["test2"] = "test2"

	var ext interface{} = mpIface

	t := &T{
		Ext: &ext,
	}

	for n := 0; n < b.N; n++ {
		InterfaceToMap(t.Ext)
	}
}