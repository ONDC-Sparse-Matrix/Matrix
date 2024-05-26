package utils

func GetDatabaseName(pincode int32) string {
	regionCode := pincode / 100000
	switch regionCode {
	case 1, 2:
		return "matrix_map_1"
	case 3, 4:
		return "matrix_map_2"
	case 5, 6:
		return "matrix_map_3"
	case 7, 8, 9:
		return "matrix_map_4"
	default:
		return "0"
	}
}