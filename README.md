# Job application tracker api

#### Creating a new module
###### create handler file

type Handler struct {
	Service *Service
	Router *http.ServeMux
}

func NewHandler(service *Service, router *http.ServeMux) *Handler {
	return &Handler{Service: service, Router: router}
}

func (h *Handler) InitRoutes() {
	color.Green("Initializing **module name** Routes")
	h.Router.HandleFunc("/application", h.GetApplication)
}

###### create Service file

type Service struct {
	Repository *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{Repository: repo}
}


###### create Repository file

type Repository struct {
	DB *sql.DB
}

type **ModuleName** struct {}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}
