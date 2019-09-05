package interfaceToMap

import (
	"reflect"
	"testing"
)

func TestConvertInterfaceToMapStrings(t *testing.T) {
	type args struct {
		i            *interface{}
		convertToMap map[string]string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestInterfaceToMap(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceToMap(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InterfaceToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

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