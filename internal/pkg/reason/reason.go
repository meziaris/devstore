package reason

var (
	InternalServerError = "internal server error"
	RequestFormError    = "request format is not valid"
)

// session
var (
	UserAlreadyExist   = "user already exist"
	RegisterFailed     = "failed to register user"
	UserNotFound       = "user not exist"
	FailedLogin        = "failed to login, your email or password is incorrect"
	FailedLogout       = "failed logout"
	Unauthorized       = "unauthorized request"
	FailedRefreshToken = "failed to refresh token, please check your token"
)

var (
	CategoryNotFound        = "category not found"
	CategoryCannotCreate    = "cannot create category"
	CategoryCannotBrowse    = "cannot browse category"
	CategoryCannotUpdate    = "cannot update category"
	CategoryCannotDelete    = "cannot delete category"
	CategoryCannotGetDetail = "cannot get detail"
)

var (
	ProductNotFound        = "product not found"
	ProductCannotCreate    = "cannot create product"
	ProductCannotBrowse    = "cannot browse product"
	ProductCannotUpdate    = "cannot update product"
	ProductCannotDelete    = "cannot delete product"
	ProductCannotGetDetail = "cannot get detail"
)
