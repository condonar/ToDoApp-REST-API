package tasks_postgres_repository

import (
	core_postgres_pool "github.com/condonar/ToDoApp-REST-API/internal/core/repository/postgres/pool"
)

type TasksReposirory struct {
	pool core_postgres_pool.Pool
}

func NewTaskRepository(pool core_postgres_pool.Pool) *TasksReposirory {
	return &TasksReposirory{
		pool: pool,
	}
}
