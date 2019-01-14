package util

type Course struct {
	ID            string   `json:"id"`
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Category      string   `json:"category"`
	CatSlug       string   `json:"cat_slug"`
	RemoteUrl     string   `json:"remote_url"`
	Slug          string   `json:"slug"`
	PostIMG       string   `json:"post_img"`
	FrontIMG      string   `json:"front_img"`
	Rating        string   `json:"rating"`
	Enrolled      string   `json:"enrolled"`
	Language      string   `json:"language"`
	CurrentPrice  string   `json:"current_price"`
	OriginalPrice string   `json:"original_price"`
	DiscountRate  string   `json:"discount_rate"`
	CouponCode    string   `json:"coupon_code"`
	ValidCoupon   bool     `json:"valid_coupon"`
	Timestamp     int64    `json:"timestamp"`
	PrettyURL     string   `json:"pretty_url"`
	WillLearn     []string `json:"will_learn"`
}

type ApiData struct {
	RedeemCoupon CouponComponents `json:"redeem_coupon"`
	Purchase Purchase `json:"purchase"`
}

type CouponComponents struct {
	Error     string `json:"error"`
	IsApplied bool   `json:"is_applied"`
	Code      string `json:"code"`
}

type Purchase struct {
	Discount Discount `json:"discount"`
}

type Discount struct {
	Price Price `json:"price"`
	ListPrice ListPrice `json:"list_price"`
	DiscountPercent int `json:"discount_percent"`
}

type Price struct {
	Amount float32 `json:"amount"`
}

type ListPrice struct {
	Amount float32 `json:"amount"`
}

type CourseComponents struct {
	Class       string            `json:"_class"`
	ID          int               `json:"id"`
	Title       string            `json:"title"`
	Headline    string            `json:"headline"`
	SubCategory CourseSubCategory `json:"primary_subcategory"`
	Category    CourseCategory    `json:"primary_category"`
	URL         string            `json:"url"`
	PostIMG     string            `json:"image_480x270"`
	FrontIMG    string            `json:"image_304x171"`
	Rating      float64           `json:"avg_rating"`
	Subscribers int               `json:"num_subscribers"`
	Language    CourseLocale      `json:"locale"`
	Price       string            `json:"price"`
}

type CourseCategory struct {
	Class string `json:"_class"`
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type CourseLocale struct {
	Class        string `json:"_class"`
	Locale       string `json:"locale"`
	EnglishTitle string `json:"english_title"`
	Title        string `json:"title"`
}

type CourseSubCategory struct {
	Class string `json:"_class"`
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type DBData struct {
	Links []UdemyLink `json:"links"`
}

type UdemyLink struct {
	Link    string `json:"link"`
	Message string `json:"message"`
	Time    int64  `json:"time"`
}


