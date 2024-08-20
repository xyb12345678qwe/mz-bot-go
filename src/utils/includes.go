package utils
// Includes 检查切片中是否包含特定的元素
func Includes(slice []interface{}, item interface{}) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}
//转换为interface{}切片
///args 数组
func ToInterfaceSlice(args ...interface{}) []interface{} {
	result := make([]interface{}, 0)
	for _, arg := range args {
		switch slice := arg.(type) {
		case []string:
			for _, v := range slice {
				result = append(result, v)
			}
		case []int:
			for _, v := range slice {
				result = append(result, v)
			}
		}
	}
	return result
}