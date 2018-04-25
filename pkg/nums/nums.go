package nums

// PtrToInt64OrDefault
func PtrToInt64OrDefault(ptr *int64, def int64) int64 {
	if ptr == nil {
		return def
	}
	return *ptr
}

// PtrToInt64OrDefault
func PtrToInt32OrDefault(ptr *int32, def int32) int32 {
	if ptr == nil {
		return def
	}
	return *ptr
}
