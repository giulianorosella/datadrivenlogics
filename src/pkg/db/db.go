package persistance

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/giulianorosella/ddlogic/pkg/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/microsoft/go-mssqldb/azuread"
)

func InitDb(server, user, password, database string, port int, isAzure bool) (*sql.DB, error) {
	var db *sql.DB
	connString := fmt.Sprintf("%s:%s@tcp(172.17.0.1:%d)/%s", user, password, port, database)
	var err error

	if isAzure {
		connString = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
			server, user, password, port, database)
		db, err = sql.Open(azuread.DriverName, connString)

	} else {
		db, err = sql.Open("mysql", connString)
	}
	if err != nil {
		log.Print("Error creating connection pool: ", err.Error())
		return nil, err
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Print(err.Error())
		return nil, err

	}
	fmt.Printf("Connected to db!")
	return db, nil

}

func CreateFormula(db *sql.DB, formula models.Formula, c, v int, isAzure bool) (int64, error) {
	ctx := context.Background()
	var err error

	if db == nil {
		err = fmt.Errorf("createFormula: db is null")
		return -1, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		err = fmt.Errorf("createFormula: db is unalive")
		return -1, err
	}

	tableName := createTableName(c, v)

	tsql := `
      INSERT INTO  ` + tableName + `  (expression, is_classic_th, is_intuitionist_th) 
	  VALUES (@Expression, @IsClassicTh, @IsIntuitionistTh);
      select isNull(SCOPE_IDENTITY(), -1);
    `
	if !isAzure {
		tsql = `
      INSERT INTO  ` + tableName + `  (expression, is_classic_th, is_intuitionist_th) 
	  VALUES (?, ?, ?)
    `
	}

	stmt, err := db.Prepare(tsql)
	if err != nil {
		return -1, err
	}

	if isAzure {
		row := stmt.QueryRowContext(
			ctx,
			sql.Named("Expression", formula.Expression),
			sql.Named("IsClassicTh", formula.IsClassicTh),
			sql.Named("IsIntuitionistTh", formula.IsIntuitionistTh),
		)
		var newID int64
		err = row.Scan(&newID)
		if err != nil {
			return -1, err
		}
		return newID, nil

	} else {
		defer stmt.Close()
		result, err := stmt.ExecContext(
			ctx,
			formula.Expression,
			formula.IsClassicTh,
			formula.IsIntuitionistTh,
		)
		if err != nil {
			return -1, err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return -1, err
		}
		return id, nil
	}

}

func CreateTable(db *sql.DB, c, v int) (bool, error) {
	ctx := context.Background()

	if db == nil {
		return false, fmt.Errorf("CreateTable: db is null")
	}

	err := db.PingContext(ctx)
	if err != nil {
		return false, fmt.Errorf("CreateTable: db is not alive: %w", err)
	}

	tableName := createTableName(c, v)

	createTableQuery := fmt.Sprintf(`
        CREATE TABLE %s (
            id INT AUTO_INCREMENT PRIMARY KEY,
            expression VARCHAR(255) NOT NULL,
            is_classic_th TINYINT NOT NULL DEFAULT -1,
            is_intuitionist_th TINYINT NOT NULL DEFAULT -1
        ) ENGINE=InnoDB;`, tableName)

	_, err = db.ExecContext(ctx, createTableQuery)
	if err != nil {
		return false, fmt.Errorf("CreateTable: error while creating table: %w", err)
	}

	return true, nil
}

func CreateInputs(db *sql.DB, c, v int) (int64, error) {
	ctx := context.Background()

	tsql := `
      INSERT INTO Inputs (VarNum, ConNum) 
	  VALUES (@VarNum, @ConNum);
      select isNull(SCOPE_IDENTITY(), -1);
    `

	stmt, err := db.Prepare(tsql)
	if err != nil {
		log.Printf("Error when Creating ")
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
		log.Print("InputsExist: db is null")
		return false, fmt.Errorf("InputsExist: db is null")
	}

	err := db.PingContext(ctx)
	if err != nil {
		log.Print("Error when checking db: ", err)
		return false, err
	}

	tableName := createTableName(c, v)

	existsTsql := `
	SELECT COUNT(*)
	FROM information_schema.tables
	WHERE table_schema = ? 
	  AND table_name = ?
	LIMIT 1
	`
	var count int

	err = db.QueryRowContext(ctx, existsTsql, "ddlogic", tableName).Scan(&count)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Table %s not exists", tableName)
			return false, nil
		}
		log.Print("Something went wrong when looking for table: ", err)
		return false, err
	}

	if count == 0 {
		log.Printf("Table %s not exists", tableName)
		return false, nil
	}
	return true, nil
}

func InputsExistAzure(db *sql.DB, c, v int) (bool, error) {
	ctx := context.Background()

	if db == nil {
		log.Print("InputsExist: db is null")
		return false, fmt.Errorf("InputsExist: db is null")
	}

	err := db.PingContext(ctx)
	if err != nil {
		log.Print("Error when checking db: ", err)
		return false, err
	}

	tableName := createTableName(c, v)

	existsTsql := `
	SELECT COUNT(*)
	FROM information_schema.tables
	WHERE table_schema = @schema
	  AND table_name = @table;
	`
	var count int

	err = db.QueryRow(existsTsql, sql.Named("schema", "dbo"), sql.Named("table", tableName)).Scan(&count)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Table %s not exists", tableName)
			return false, nil
		}
		log.Print("Something went wrong when looking for table: ", err)
		return false, err
	}

	if count == 0 {
		log.Printf("Table %s not exists", tableName)
		return false, nil
	}
	return true, nil
}

func CreateTableAzure(db *sql.DB, c, v int) (bool, error) {
	ctx := context.Background()

	if db == nil {
		return false, fmt.Errorf("CreateTableAzure: db is null")
	}

	err := db.PingContext(ctx)
	if err != nil {
		return false, fmt.Errorf("CreateTableAzure: db is not alive: %w", err)
	}

	tableName := createTableName(c, v)

	createTableQuery := fmt.Sprintf(`
		CREATE TABLE %s (
			id INT IDENTITY(1,1) PRIMARY KEY,
			expression NVARCHAR(255) NOT NULL,
			is_classic_th BIT NOT NULL DEFAULT 0,
			is_intuitionist_th BIT NOT NULL DEFAULT 0
		);`, tableName)

	_, err = db.ExecContext(ctx, createTableQuery)
	if err != nil {
		return false, fmt.Errorf("CreateTableAzure: error while creating table: %w", err)
	}

	return true, nil
}

func createTableName(c, v int) string {
	return fmt.Sprintf("Inputs_c_%d_v_%d", c, v)
}
