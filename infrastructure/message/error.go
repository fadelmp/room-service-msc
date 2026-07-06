package message

var (
	Failed          = "FAILED"
	ErrInitializeDB = "ERROR_INITIALIZE_DATABASE"
	ErrBeginTrx     = "ERROR_BEGIN_TRANSACTIONS"

	ErrHotelExists       = "ERROR_HOTEL_ALREADY_EXISTS"
	ErrGenerateHotelCode = "ERROR_GENERATE_HOTEL_CODE"

	ErrFailedGetHotel    = "ERROR_FAILED_TO_GET_HOTEL"
	ErrHotelNotFound     = "ERROR_HOTEL_NOT_FOUND"
	ErrCreateHotelFailed = "ERROR_CREATE_HOTEL"
	ErrUpdateHotelFailed = "ERROR_UPDATE_HOTEL"
	ErrDeleteHotelFailed = "ERROR_DELETE_HOTEL"
)
