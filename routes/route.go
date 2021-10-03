package routes

import (
	"fngc/controller"

	_ "fngc/docs"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter(appPort, hostAddress string) *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Use(middleware.Logger)

	router.Get("/", controller.FNGCInit)
	router.Mount("/auth", authRouter())
	router.Mount("/user", userRouter())
	router.Mount("/student", studentRouter())
	router.Mount("/admin", adminRouter())

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(hostAddress+appPort+"/swagger/doc.json"),
	))

	return router
}

func authRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/signup", controller.Registration)
	router.Post("/login", controller.UserLogin)
	router.Post("/verifyotp", controller.VerifyUser)
	router.Post("/tutor/signup", controller.TutorRegistration)

	return router
}

func userRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/studyabroad/", controller.StudyAbroad)
	router.Post("/contactus/", controller.ContactUs)

	return router
}

func studentRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/examination/", controller.ExaminationApply)
	// router.Post()

	return router
}

func adminRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/examination/all", controller.GetAllExaminationProfile)
	router.Get("/examination/{profile}", controller.GetExaminationProfile)
	router.Get("/tutor/all", controller.GetAllTutor)
	router.Get("/student/all", controller.GetAllStudents)
	router.Get("/student/abroad/all", controller.GetAllAbroadStudies)

	return router
}
