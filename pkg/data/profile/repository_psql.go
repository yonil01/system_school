package profile

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"foro-hotel/internal/logger"
	"foro-hotel/internal/models"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
)

type psql struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

const (
	sqlserverExecuteSP = `execute`
)

func (s *psql) getUserbyId(id string) (*models.User, int, error) {
	document_id, _ := strconv.Atoi(id)
	mdl := models.User{}
	user_id := 1
	const sqlGetInfoKeywords = ` EXECUTE dbo.Get_User @document_id= %d, @user_id= %d `
	sqlExecute := fmt.Sprintf(sqlGetInfoKeywords, document_id, user_id)
	err := s.DB.Get(&mdl, sqlExecute, sql.Named("user_id", user_id))
	if err != nil {
		return nil, 1, errors.New("El suuario no existe!")
	}
	return &mdl, 82, nil
}

func (s *psql) getComentariosProcedure(id string) ([]*models.Comentario, int, error) {
	document_id, _ := strconv.Atoi(id)
	user_id := 1
	res := []*models.Comentario{}
	const sqlGetInfoKeywords = ` EXECUTE dbo.Get_Comentarios @document_id= %d, @user_id= %d `
	sqlExecute := fmt.Sprintf(sqlGetInfoKeywords, document_id, user_id)
	err := s.DB.Select(&res, sqlExecute, sql.Named("user_id", user_id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, 3, nil
		}
		logger.Error.Printf(s.TxID, " - couldn't execute getDatosBasicos: %v", err)
		return nil, 3, nil
	}

	return res, 3, nil
}

func (s *psql) getPdf(text_hash string) (*models.Reservacion, int, error) {
	mdl := models.Reservacion{}
	query := `SELECT id, cliente_id, negocio_id, email, datos, telefono, celular, direccion, url, created_at, text_hash FROM dbo.reservaciones WHERE text_hash = @text_hash`
	err := s.DB.Get(&mdl, query, sql.Named("text_hash", text_hash))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, 3, nil
		}
		logger.Error.Printf(s.TxID, " - couldn't execute getDatosBasicos: %v", err)
		return nil, 3, nil
	}

	return &mdl, 3, nil
}

func (s *psql) getReservacionByDocument(id string) ([]*models.Reservacion, int, error) {
	var mdl []*models.Reservacion
	query := `SELECT id, cliente_id, email, datos, telefono, celular, direccion, created_at, text_hash, name_negocio FROM dbo.reservaciones WHERE cliente_id = @id`
	err := s.DB.Select(&mdl, query, sql.Named("id", id))
	if err != nil {
		return nil, 29, errors.New("Error en obtener el negocio!")
	}

	return mdl, 1, nil
}

func (s *psql) getReservaciones(id string) ([]*models.Reservacion, int, error) {
	var mdl []*models.Reservacion
	query := `SELECT id, cliente_id, negocio_id, email, datos, telefono, celular, direccion, created_at, text_hash, name_negocio FROM dbo.reservaciones`
	err := s.DB.Select(&mdl, query)
	if err != nil {
		return nil, 29, errors.New("Error en obtener el negocio!")
	}

	return mdl, 1, nil
}

func (s *psql) getCountries(data string) ([]*models.Country, int, error) {
	var mdl []*models.Country
	query := `SELECT PaisCodigo, PaisNombre FROM dbo.pais`
	err := s.DB.Select(&mdl, query)
	if err != nil {
		return nil, 29, errors.New("Error en obtener el negocio!")
	}

	return mdl, 1, nil
}

func (s *psql) getCities(data string) ([]*models.City, int, error) {
	var mdl []*models.City
	query := `SELECT CiudadNombre, PaisCodigo FROM dbo.ciudad`
	err := s.DB.Select(&mdl, query)
	if err != nil {
		return nil, 29, errors.New("Error en obtener el negocio!")
	}

	return mdl, 1, nil
}

func (s *psql) getUser(role int) ([]*models.User, error) {
	var mdl []*models.User
	query := `SELECT id, dni, matricula, username, names, lastnames, sexo, status, date_admission, date_birth, email, is_delete,password, created_at, updated_at, role FROM dbo.users WHERE role = @role`
	err := s.DB.Select(&mdl, query, sql.Named("role", role))
	if err != nil {
		return nil, errors.New("Error en obtener el negocio!")
	}

	return mdl, nil
}

func (s *psql) updateUser(mdl models.User) (*models.User, error) {
	const sqlUpdate = `UPDATE dbo.users SET dni = :dni, username = :username, sexo = :sexo, email = :email, date_birth = :date_birth, date_admission = :date_admission WHERE matricula = :matricula `

	_, err := s.DB.NamedExec(sqlUpdate, &mdl)
	if err != nil {
		logger.Error.Println(s.TxID, "-cound't update User: %V", err)
		return &mdl, errors.New("-cound't update User: %V")
	}

	return &mdl, nil
}

func (s *psql) insertUser(mdl models.User) (*models.User, error) {
	const sqlUpdate = `insert into dbo.users(dni,matricula,username,names,lastnames,sexo,status,date_admission, date_birth,email,is_delete,password,created_at, updated_at, role) values(:dni,(SELECT MAX(matricula) FROM users),:username,:names,:lastnames,:sexo,:status,:date_admission, :date_birth,:email,:is_delete,:password,:created_at, :updated_at, :role)
`
	_, err := s.DB.NamedExec(sqlUpdate, &mdl)
	if err != nil {
		logger.Error.Println(s.TxID, "-cound't insert User: %V", err)
		return &mdl, errors.New("Error en insertar el negocio!")
	}

	return &mdl, nil
}

func (s *psql) deleteUser(mdl models.User) (*models.User, error) {
	const sqlUpdate = `DELETE FROM dbo.users WHERE matricula = :matricula `
	_, err := s.DB.NamedExec(sqlUpdate, &mdl)
	if err != nil {
		logger.Error.Println(s.TxID, "-cound't insert User: %V", err)
		return &mdl, errors.New("Error en insertar el negocio!")
	}

	return &mdl, nil
}

func (s *psql) getClassrooms() ([]*models.Classroom, error) {
	var mdl []*models.Classroom
	query := `SELECT id, name, description, nivel, range, status, created_at, updated_at FROM dbo.grados`
	err := s.DB.Select(&mdl, query)
	if err != nil {
		return nil, errors.New("Error en obtener el negocio!")
	}

	return mdl, nil
}

func (s *psql) updateClassroom(mdl models.Classroom) (*models.Classroom, error) {
	const sqlUpdate = `UPDATE dbo.aula SET name = :name, description = :description, nivel = :nivel, grado = :grado, section = :section WHERE matricula = :matricula `

	_, err := s.DB.NamedExec(sqlUpdate, &mdl)
	if err != nil {
		logger.Error.Println(s.TxID, "-cound't insert User: %V", err)
		return &mdl, errors.New("Error en insertar el negocio!")
	}

	return &mdl, nil
}

func (s *psql) insertClassroom(mdl models.Classroom) (*models.Classroom, error) {
	const sqlUpdate = `insert into dbo.aula(name,description,nivel,status,grado,section,created_at, updated_at) values(:dni,(SELECT MAX(matricula) FROM User),:username,:names,:lastnames,:sexo,:status,:date_admission, :date_birth,:email,:is_delete,:password,:created_at, :updated_at)
`
	_, err := s.DB.NamedExec(sqlUpdate, &mdl)
	if err != nil {
		logger.Error.Println(s.TxID, "-cound't insert User: %V", err)
		return &mdl, errors.New("Error en insertar el negocio!")
	}

	return &mdl, nil
}

func (s *psql) deleteClassroom(mdl models.Classroom) (*models.Classroom, error) {
	const sqlUpdate = `DELETE FROM dbo.aula WHERE id = :id `
	_, err := s.DB.NamedExec(sqlUpdate, &mdl)
	if err != nil {
		logger.Error.Println(s.TxID, "-cound't insert User: %V", err)
		return &mdl, errors.New("Error en insertar el negocio!")
	}

	return &mdl, nil
}

func (s *psql) getSubjects() ([]*models.Subject, error) {
	var mdl []*models.Subject
	query := `SELECT id, name, description, status, is_delete, created_at, updated_at FROM dbo.materia`
	err := s.DB.Select(&mdl, query)
	if err != nil {
		return nil, errors.New("Error en obtener el negocio!")
	}

	return mdl, nil
}

func (s *psql) updateSubject(mdl models.Subject) (*models.Subject, error) {
	const sqlUpdate = `UPDATE dbo.materia SET name = :name, description = :description, updated_at = :updated_at WHERE matricula = :matricula `

	_, err := s.DB.NamedExec(sqlUpdate, &mdl)
	if err != nil {
		logger.Error.Println(s.TxID, "-cound't insert User: %V", err)
		return &mdl, errors.New("Error en insertar el negocio!")
	}

	return &mdl, nil
}

func (s *psql) insertSubject(mdl models.Subject) (*models.Subject, error) {
	const sqlUpdate = `insert into dbo.materia(name,description,status,is_delete,created_at, updated_at) values(:name,:description,:status,:is_delete,:created_at, :updated_at)
`
	_, err := s.DB.NamedExec(sqlUpdate, &mdl)
	if err != nil {
		logger.Error.Println(s.TxID, "-cound't insert User: %V", err)
		return &mdl, errors.New("Error en insertar el negocio!")
	}

	return &mdl, nil
}

func (s *psql) deleteSubject(mdl models.Subject) (*models.Subject, error) {
	const sqlUpdate = `DELETE FROM dbo.materia WHERE id = :id `
	_, err := s.DB.NamedExec(sqlUpdate, &mdl)
	if err != nil {
		logger.Error.Println(s.TxID, "-cound't insert User: %V", err)
		return &mdl, errors.New("Error en insertar el negocio!")
	}

	return &mdl, nil
}

func (s *psql) getPayments() ([]*models.Payment, error) {
	var mdl []*models.Payment
	query := `SELECT id, matricula, name, description, motivo, status, date_payment, user_matricula, amount, role, created_at, updated_at FROM dbo.payment_debt`
	err := s.DB.Select(&mdl, query)
	if err != nil {
		return nil, errors.New("Error en obtener el negocio!")
	}

	return mdl, nil
}

func (s *psql) updatePayment(mdl models.Payment) (*models.Payment, error) {
	const sqlUpdate = `UPDATE dbo.payment_debt SET id = :id, name = :name, description = :description,motivo = :motivo, status = :status, date_payment = :date_payment, user_matricula = :user_matricula, amount = :amount, role = :role, updated_at = :updated_at WHERE matricula = :matricula `

	_, err := s.DB.NamedExec(sqlUpdate, &mdl)
	if err != nil {
		logger.Error.Println(s.TxID, "-cound't insert User: %V", err)
		return &mdl, errors.New("Error en insertar el negocio!")
	}

	return &mdl, nil
}

func (s *psql) insertPayment(mdl models.Payment) (*models.Payment, error) {
	const sqlUpdate = `insert into dbo.payment_debt(name, description,motivo, status, date_payment, user_matricula, amount, role, updated_at) values(:name, :description,: motivo, : status, : date_payment, : user_matricula, : amount, :role, :updated_at)
`
	_, err := s.DB.NamedExec(sqlUpdate, &mdl)
	if err != nil {
		logger.Error.Println(s.TxID, "-cound't insert User: %V", err)
		return &mdl, errors.New("Error en insertar el negocio!")
	}

	return &mdl, nil
}

func (s *psql) deletePayment(mdl models.Payment) (*models.Payment, error) {
	const sqlUpdate = `DELETE FROM dbo.payment_debt WHERE matricula = :matricula `
	_, err := s.DB.NamedExec(sqlUpdate, &mdl)
	if err != nil {
		logger.Error.Println(s.TxID, "-cound't insert User: %V", err)
		return &mdl, errors.New("Error en insertar el negocio!")
	}

	return &mdl, nil
}

func (s psql) ExecuteSP(m *models.Report) ([]map[string]interface{}, error) {
	rs := make([]map[string]interface{}, 0)
	r := bytes.Buffer{}
	r.WriteString(fmt.Sprintf(`%s %s `, sqlserverExecuteSP, m.Procedure))
	vs := make([]interface{}, 0)
	for i, v := range m.Parameters {
		valueClear := strings.Replace(v, `'`, "", -1)
		_, err := r.WriteString(fmt.Sprintf(`@%s='%s',`, i, valueClear))
		if err != nil {
			logger.Error.Printf("agregando parametros a la ejecucion del SP en sqlserverExecuteSP: %v", err)
			return rs, err
		}
		vs = append(vs, sql.Named(fmt.Sprintf("%s", i), v))

	}

	r.Truncate(r.Len() - 1)
	fmt.Println(r.String())
	stmt, err := s.DB.Prepare(r.String())
	if err != nil {
		logger.Error.Printf("preparando consulta sqlserverExecuteSP: %v", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		logger.Error.Printf("ejecutando sqlserverExecuteSP user: %v", err)
		return rs, err
	}
	defer rows.Close()

	cols, _ := rows.Columns()
	for rows.Next() {
		r := make(map[string]interface{})
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}
		// Scan the result into the column pointers...
		if err = rows.Scan(columnPointers...); err != nil {
			logger.Error.Printf("no se pudo escanear las columnas de la consulta sqlserverGetInfoktg: %v", err)
			return nil, err
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			r[colName] = *val
		}
		rs = append(rs, r)
	}
	return rs, nil
}

func (s *psql) getSections() ([]*models.Sections, error) {
	var mdl []*models.Sections
	query := `SELECT id, name, grado_id, status, created_at, updated_at FROM dbo.section`
	err := s.DB.Select(&mdl, query)
	if err != nil {
		return nil, errors.New("Error en obtener el negocio!")
	}

	return mdl, nil
}

func (s *psql) updateSection(mdl models.Sections) (*models.Sections, error) {
	const sqlUpdate = `UPDATE dbo.aula SET name = :name, description = :description, nivel = :nivel, grado = :grado, section = :section WHERE section = :matricula `

	_, err := s.DB.NamedExec(sqlUpdate, &mdl)
	if err != nil {
		logger.Error.Println(s.TxID, "-cound't insert User: %V", err)
		return &mdl, errors.New("Error en insertar el negocio!")
	}

	return &mdl, nil
}

func (s *psql) insertSection(mdl models.Sections) (*models.Sections, error) {
	const sqlUpdate = `insert into dbo.section(name,description,nivel,status,grado,section,created_at, updated_at) values(:dni,(SELECT MAX(matricula) FROM User),:username,:names,:lastnames,:sexo,:status,:date_admission, :date_birth,:email,:is_delete,:password,:created_at, :updated_at)
`
	_, err := s.DB.NamedExec(sqlUpdate, &mdl)
	if err != nil {
		logger.Error.Println(s.TxID, "-cound't insert User: %V", err)
		return &mdl, errors.New("Error en insertar el negocio!")
	}

	return &mdl, nil
}

func (s *psql) deleteSection(mdl models.Sections) (*models.Sections, error) {
	const sqlUpdate = `DELETE FROM dbo.section WHERE id = :id `
	_, err := s.DB.NamedExec(sqlUpdate, &mdl)
	if err != nil {
		logger.Error.Println(s.TxID, "-cound't insert User: %V", err)
		return &mdl, errors.New("Error en insertar el negocio!")
	}

	return &mdl, nil
}

func newDataPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}
