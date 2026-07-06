package utils

import "strings"

// GetCode mengambil 3 karakter pertama dari nama (spasi dihapus dulu)
// Contoh: "AMARIS" -> "AMA", "PO HOTEL" -> "POH"
func GetCode(name string) string {

	name = strings.ToUpper(strings.TrimSpace(name))
	name = strings.ReplaceAll(name, " ", "")

	runes := []rune(name)
	if len(runes) > 3 {
		runes = runes[:3]
	}
	return string(runes)
}

// Slugify mengubah teks jadi slug (lowercase, spasi jadi strip)
// Contoh: "Hotel Amaris Bandung" -> "hotel-amaris-bandung"
func Slugify(text string) string {

	text = strings.ToLower(strings.TrimSpace(text))
	text = strings.ReplaceAll(text, " ", "-")
	return text
}
