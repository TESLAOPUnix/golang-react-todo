package router

import (
	"github.com/teslaopunix/golang-react-todo/middlewares"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/task", middlewares.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tasks", middlewares.CreateTasks).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/tasks/{id}", middlewares.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", middlewares.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", middlewares.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTask", middlewares.DeleteAllTask).Methods("DELETE", "OPTIONS")
	return router
}
