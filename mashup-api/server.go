package main

import (
	"github.com/anaetrezve/golang-crash-course-projects/controller"
	router "github.com/anaetrezve/golang-crash-course-projects/http"
	"github.com/anaetrezve/golang-crash-course-projects/service"
)

func main() {
	const port string = "localhost:8000"

	var (
		postDetailsService    service.PostDetailsService       = service.NewPostDetailsService()
		postDetailsController controller.PostDetailsController = controller.NewPostDetailsController(postDetailsService)
		httpRouter            router.Router                    = router.NewMuxRouter()
	)

	httpRouter.Get("/posts/details", postDetailsController.GetPostDetails)

	httpRouter.Serve(port)
}
