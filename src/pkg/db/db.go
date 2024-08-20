package persistance

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/giulianorosella/ddlogic/pkg/models"
	"github.com/microsoft/go-mssqldb/azuread"
	"golang.org/x/tools/go/cfg"
)

func InitDb() (*sql.DB, error) {
	var db = cfg.Db
	var server = cfg.Server
	var port = cfg.Port
	var user = cfg.User
	var password = cfg.Password
	var database = cfg.Database
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	var err error
	// Create connection pool
	db, err = sql.Open(azuread.DriverName, connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
		return nil, err
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err

	}
	fmt.Printf("Connected to db!")
	return db, nil

}

func CreateFormula(db *sql.DB, formula models.Formula) (int64, error) {
	ctx := context.Background()
	var err error

	if db == nil {
		err = errors.New("createFormula: db is null")
		return -1, err
	}

	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := `
      INSERT INTO Formulas (Expression, IsClassicTh, IsIntuitionistTh) 
	  VALUES (@Expression, @IsClassicTh, @IsIntuitionistTh);
      select isNull(SCOPE_IDENTITY(), -1);
    `

	stmt, err := db.Prepare(tsql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(
		ctx,
		sql.Named("Expression", formula.Expression),
		sql.Named("IsClassicTh", formula.IsClassicTh),
		sql.Named("IsIntuitionistTh", formula.IsIntuitionistTh))
	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		return -1, err
	}

	return newID, nil
}

func CreateInputs(db *sql.DB, c, v int) (int64, error) {
	ctx := context.Background()
	var err error

	if db == nil {
		err = errors.New("createInputs: db is null")
		return -1, err
	}

	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := `
      INSERT INTO Inputs (VarNum, ConNum) 
	  VALUES (@VarNum, @ConNum);
      select isNull(SCOPE_IDENTITY(), -1);
    `

	stmt, err := db.Prepare(tsql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(
		ctx,
		sql.Named("VarNum", v),
		sql.Named("ConNum", c))
	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		return -1, err
	}

	return newID, nil
}

func InputsExist(db *sql.DB, c, v int) (bool, error) {
	ctx := context.Background()

	if db == nil {
		log.Fatal("InputsExist: db is null")
		return false, errors.New("InputsExist: db is null")
	}

	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error when checking db: ", err)
		return false, err
	}

	tsql := `
        SELECT COUNT(*)
        FROM Inputs
        WHERE VarNum = @VarNum AND ConNum = @ConNum
    `

	var count int
	err = db.QueryRowContext(
		ctx,
		tsql,
		sql.Named("VarNum", v),
		sql.Named("ConNum", c),
	).Scan(&count)
	if err != nil {
		log.Fatal("Error when executing query: ", err)
		return false, err
	}

	return count > 0, nil
}
