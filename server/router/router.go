package router

import (
	"github.com/sirawann/go-react-todo/server/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/task",middleware.GetAllTasks).Methods("GET","OPTIONS")
	router.HandleFunc("/api/tasks",middleware.CreateTask).Methods("POST","OPTIONs")
	router.HandleFunc("/api/tasks/{id}",middleware.TaskComplete).Methods("PUT","OPTIONs")
	router.HandleFunc("/api/undoTask/{id}",middleware.UndoTask).Methods("PUT","OPTIONs")
	router.HandleFunc("/api/deleteTask/{id}",middleware.DeleteTask).Methods("DELETE","OPTIONs")
	router.HandleFunc("/api/deleteAllTasks/{id}",middleware.DeleteAllTasks).Methods("DELETE","OPTIONs")
	return router
}
