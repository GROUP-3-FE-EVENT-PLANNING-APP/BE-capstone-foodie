package favourites

type Core struct {
	User       User
	Restaurant Restaurant
}

type User struct {
	ID   int
	Name string
}

type Restaurant struct {
	ID           int
	Name         string
	Location     string
	MenuImageUrl string
	Category     string
	TableQuota   uint
	BookingFee   uint64
	Latitude     string
	Longitude    string
	Status       string
	FileImageUrl string
	Facility     string
	RestoImages  []RestoImage
}

type RestoImage struct {
	ID            int
	UserID        int
	User          User
	RestaurantID  int
	Restaurant    Core
	RestoImageUrl string
}

type RestoCore struct {
	ID            int
	RestoName     string
	Location      string
	Category      string
	TableQuota    uint
	Rating        float64
	RestoImageUrl string
	RestoImages   []RestoImage
}

type Business interface {
	AddFav(idResto, idFromToken int) (row int, err error)
	DeleteFavourite(idResto, idFromToken int) (row int, err error)
	GetMyFav(limitint, offsetint, idFromToken int) (response []RestoCore, err error)
}

type Data interface {
	AddFavDB(idResto, idFromToken int) (row int, err error)
	DeleteFavDB(idResto, idFromToken int) (row int, err error)
	RatingData(idResto int) (response float64, err error)
	RestoImageData(idResto int) (response string, err error)
	AllRestoData(limitint, offsetint, idFromToken int) (response []RestoCore, err error)
}
