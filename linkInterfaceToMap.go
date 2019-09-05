package interfaceToMap

// ConvertInterfaceToMapStrings use for convert *interface{} to map[string]string
func ConvertInterfaceToMapStrings(i *interface{}, convertToMap map[string]string) {
	var (
		convertedEx map[string]interface{}
	)
	ex := *(i)
	convertedEx = ex.(map[string]interface{})

	for k, ex := range convertedEx {
		convertToMap[k] = ex.(string)
	}
}


// ConvertInterfaceToMapStrings use for convert *interface{} to map[string]string
func ConvertInterfaceToMapStrings2(i *interface{}) map[string]interface{} {
	var (
		convertedEx map[string]interface{}
	)
	ex := *(i)
	convertedEx = ex.(map[string]interface{})

	return convertedEx
}