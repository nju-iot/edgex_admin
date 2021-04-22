package utils

// InStringSlice ...
func InStringSlice(str string, strList []string) bool {
	for _, v := range strList {
		if v == str {
			return true
		}
	}
	return false
}

// InInt64Slice ...
func InInt64Slice(i64 int64, i64List []int64) bool {
	for _, v := range i64List {
		if v == i64 {
			return true
		}
	}
	return false
}

// DeduplicationI64List ...
func DeduplicationI64List(i64List []int64) (deduplicationI64List []int64) {
	i64Deduplication := make(map[int64]bool)
	for _, i64 := range i64List {
		if !i64Deduplication[i64] {
			deduplicationI64List = append(deduplicationI64List, i64)
			i64Deduplication[i64] = true
		}
	}
	return
}
