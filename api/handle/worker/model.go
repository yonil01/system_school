package worker

import (
	"foro-hotel/pkg/wf/worker"
)

type requestWorker struct {
	Id        int   `json:"id" db:"id" valid:"-"`
	Matricula int64 `json:"matricula" db:"matricula" valid:"required"`
	Status    int   `json:"status" db:"status" valid:"required"`
}

type responseWorkers struct {
	Error bool             `json:"error"`
	Data  []*worker.Worker `json:"data"`
	B64   string           `json:"image"`
	Code  int              `json:"code"`
	Type  string           `json:"type"`
	Msg   string           `json:"msg"`
}

type responseWorker struct {
	Error bool           `json:"error"`
	Data  *worker.Worker `json:"data"`
	B64   string         `json:"image"`
	Code  int            `json:"code"`
	Type  string         `json:"type"`
	Msg   string         `json:"msg"`
}
