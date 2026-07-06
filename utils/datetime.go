package utils

import "time"

const secondsPerHour = 60 * 60

// AddHours menambah/mengurangi timestamp unix (detik) dengan sejumlah jam
// Contoh: AddHours(unixTimestamp, 7) -> tambah 7 jam (offset WIB)
func AddHours(unixTimestamp int64, hours int) int64 {

	return unixTimestamp + int64(hours*secondsPerHour)
}

// IsExpired mengecek apakah suatu timestamp unix sudah lewat dari waktu sekarang
func IsExpired(expiryUnix int64) bool {

	return time.Now().Unix() > expiryUnix
}

// GetTodayNoonUnix menghasilkan timestamp unix untuk jam 12 siang hari ini,
// berguna untuk rolling expiry QR code Serva
func GetTodayNoonUnix(loc *time.Location) int64 {

	now := time.Now().In(loc)
	noon := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, loc)

	return noon.Unix()
}

// FormatDate memformat unix timestamp menjadi string sesuai layout yang diberikan
// Contoh: FormatDate(unix, "2006-01-02 15:04:05")
func FormatDate(unixTimestamp int64, layout string) string {

	return time.Unix(unixTimestamp, 0).Format(layout)
}
