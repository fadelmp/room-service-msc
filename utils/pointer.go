package utils

// PointerString mengubah string value menjadi pointer,
// berguna untuk DTO/field optional
func PointerString(s string) *string {

	return &s
}

// PointerInt mengubah int value menjadi pointer
func PointerInt(i int) *int {

	return &i
}

// PointerInt64 mengubah int64 value menjadi pointer
func PointerInt64(i int64) *int64 {

	return &i
}

// StringOrDefault mengembalikan string default kalau input kosong
func StringOrDefault(s, def string) string {

	if s == "" {
		return def
	}

	return s
}
