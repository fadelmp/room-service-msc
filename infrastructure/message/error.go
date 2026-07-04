package message

var (
	Failed          = "Failed"
	ErrInitializeDB = "Error Initialize Database"
	ErrBeginTrx     = "Error Begin Transaction"

	ErrFailedGetBrand    = "Error Failed Get Brand"
	ErrBrandNotFound     = "Error Brand Not Found"
	ErrBrandExists       = "Error Brand Exists"
	ErrCreateBrandFailed = "Error Create Brand Failed"
	ErrUpdateBrandFailed = "Error Update Brand Failed"
	ErrDeleteBrandFailed = "Error Delete Brand Failed"

	ErrFailedGetHotel    = "Error Failed Get Hotel"
	ErrHotelNotFound     = "Error Hotel Not Found"
	ErrCreateHotelFailed = "Error Create Hotel Failed"
	ErrUpdateHotelFailed = "Error Update Hotel Failed"
	ErrDeleteHotelFailed = "Error Delete Hotel Failed"
)
