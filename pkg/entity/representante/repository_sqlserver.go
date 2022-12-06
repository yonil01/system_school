package representante

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

func newRepresentanteSqlServerRepository(db *sqlx.DB, user *models.User, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *sqlserver) create(m *Representante) error {
	var id int
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const sqlInsert = `INSERT INTO entity.representante (matricula_user, type_representante, notification, dni, direction, names, lastnames, cell_phone, email, status, is_delete, created_at, updated_at) VALUES (@matricula_user, @type_representante, @notification, @dni, @direction, @names, @lastnames, @cell_phone, @email, @status, @is_delete, @created_at, @updated_at) SELECT ID = convert(bigint, SCOPE_IDENTITY()) `
	stmt, err := s.DB.Prepare(sqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		sql.Named("matricula_user", m.MatriculaUser),
		sql.Named("type_representante", m.TypeRepresentante),
		sql.Named("notification", m.Notification),
		sql.Named("dni", m.Dni),
		sql.Named("direction", m.Direction),
		sql.Named("names", m.Names),
		sql.Named("lastnames", m.Lastnames),
		sql.Named("cell_phone", m.CellPhone),
		sql.Named("email", m.Email),
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
func (s *sqlserver) update(m *Representante) error {
	date := time.Now()
	m.UpdatedAt = date
	const sqlUpdate = `UPDATE entity.representante SET matricula_user = :matricula_user, type_representante = :type_representante, notification = :notification, dni = :dni, direction = :direction, names = :names, lastnames = :lastnames, cell_phone = :cell_phone, email = :email, status = :status, is_delete = :is_delete, updated_at = :updated_at WHERE id = :id `
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
	const sqlDelete = `DELETE FROM entity.representante WHERE id = :id `
	m := Representante{ID: id}
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
func (s *sqlserver) getByID(id int) (*Representante, error) {
	const sqlGetByID = `SELECT convert(nvarchar(50), id) id , matricula_user, type_representante, notification, dni, direction, names, lastnames, cell_phone, email, status, is_delete, created_at, updated_at FROM entity.representante  WITH (NOLOCK)  WHERE id = @id `
	mdl := Representante{}
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
func (s *sqlserver) getAll() ([]*Representante, error) {
	var ms []*Representante
	const sqlGetAll = `SELECT convert(nvarchar(50), id) id , matricula_user, type_representante, notification, dni, direction, names, lastnames, cell_phone, email, status, is_delete, created_at, updated_at FROM entity.representante  WITH (NOLOCK) `

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *sqlserver) getByMatriculaUser(matriculaUser int64) (*Representante, error) {
	const sqlGetByID = `SELECT top 1 convert(nvarchar(50), id) id , matricula_user, type_representante, notification, dni, direction, names, lastnames, cell_phone, email, status, is_delete, created_at, updated_at FROM entity.representante  WITH (NOLOCK)  WHERE matricula_user = @matricula_user ORDER BY created_at DESC `
	mdl := Representante{}
	err := s.DB.Get(&mdl, sqlGetByID, sql.Named("matricula_user", matriculaUser))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}
