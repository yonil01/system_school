package materia

import (
	"database/sql"
	"fmt"
	"foro-hotel/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

// sqlServer estructura de conexión a la BD de mssql
type sqlserver struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newMateriaSqlServerRepository(db *sqlx.DB, user *models.User, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *sqlserver) create(m *Materia) error {
	var id int
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const sqlInsert = `INSERT INTO dbo.materia (name, description, status, is_delete, created_at, updated_at) VALUES (@name, @description, @status, @is_delete, @created_at, @updated_at) SELECT ID = convert(bigint, SCOPE_IDENTITY()) `
	stmt, err := s.DB.Prepare(sqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		sql.Named("name", m.Name),
		sql.Named("description", m.Description),
		sql.Named("status", m.Status),
		sql.Named("is_delete", m.IsDelete),
		sql.Named("created_at", m.CreatedAt),
		sql.Named("updated_at", m.UpdatedAt),
	).Scan(&id)
	if err != nil {
		return err
	}
	m.ID = int(id)
	return nil
}

// Update actualiza un registro en la BD
func (s *sqlserver) update(m *Materia) error {
	date := time.Now()
	m.UpdatedAt = date
	const sqlUpdate = `UPDATE dbo.materia SET name = :name, description = :description, status = :status, is_delete = :is_delete, updated_at = :updated_at WHERE id = :id `
	rs, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *sqlserver) delete(id int) error {
	const sqlDelete = `DELETE FROM dbo.materia WHERE id = :id `
	m := Materia{ID: id}
	rs, err := s.DB.NamedExec(sqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// GetByID consulta un registro por su ID
func (s *sqlserver) getByID(id int) (*Materia, error) {
	const sqlGetByID = `SELECT convert(nvarchar(50), id) id , name, description, status, is_delete, created_at, updated_at FROM dbo.materia  WITH (NOLOCK)  WHERE id = @id `
	mdl := Materia{}
	err := s.DB.Get(&mdl, sqlGetByID, sql.Named("id", id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *sqlserver) getAll() ([]*Materia, error) {
	var ms []*Materia
	const sqlGetAll = `SELECT convert(nvarchar(50), id) id , name, description, status, is_delete, created_at, updated_at FROM dbo.materia  WITH (NOLOCK) `

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *sqlserver) getGradeId(gradeId int) ([]*Materia, error) {
	var ms []*Materia

	const sqlGetByID = `select m.id, m.name, m.description, m.status, m.is_delete, m.created_at, m.updated_at from dbo.grado_curso gc
join dbo.materia m on gc.curso_id = m.id
where grado_id = @grado_id `
	err := s.DB.Select(&ms, sqlGetByID, sql.Named("grado_id", gradeId))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}