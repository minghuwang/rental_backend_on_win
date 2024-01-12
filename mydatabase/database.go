package mydatabase

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Register the MySQL driver
)

func CloseDB() {
	if Db != nil {
		Db.Close()
	}
}

func InitDB() {
	var err error
	Db, err = sql.Open("mysql", "root:Richard123@tcp(localhost:3306)/rental_system")
	// Oracle cloud server
	// Db, err = sql.Open("mysql", "ADMIN:Richard123Fornzpr0!@tcp(adb.ap-singapore-1.oraclecloud.com:1522)/rentalsystem")
	if err != nil {
		fmt.Printf("sql.Open error:%v", err)
		panic(err)
		// Handle error
	}

	createTableStr := `CREATE TABLE IF NOT EXISTS property_info
		(property_id INT AUTO_INCREMENT PRIMARY KEY, 
		address VARCHAR(300) UNIQUE NOT NULL, 
		picture_link VARCHAR(50),
		open_time_1 INT(64),
		open_time_2 INT(64))`
	_, err = Db.Exec(createTableStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table created successfully!")

}

func GetProperties(db *sql.DB, tableName string) ([]PropertyInfoSchema, error) {
	fmt.Printf("Start GetProperties..\n")
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %v", tableName))
	if err != nil || rows == nil {
		return nil, fmt.Errorf("rows is null or db.Query error: %v", err)
	}
	defer rows.Close()
	pIs := []PropertyInfoSchema{}
	for rows.Next() {
		pI := PropertyInfoSchema{}
		// Scan column values into variables
		if err := rows.Scan(
			&pI.PropertyID,
			&pI.Address,
			&pI.PictureLink,
			&pI.OpenTime1,
			&pI.OpenTime2); err != nil {
				return nil, fmt.Errorf("rows.Scan error: %v", err)
			}
			pIs = append(pIs, pI)
			
			// Utilize retrieved data
			fmt.Printf("%v\n", pI)
		}
		
		fmt.Printf("End GetProperties..\n")
		return pIs, nil
}

func InsertProperty(db *sql.DB, pI *PropertyInfoSchema) (*PropertyInfoSchema, error) {

	insertStmt := fmt.Sprintf(`INSERT INTO %v (address, picture_link, open_time_1, open_time_2) VALUES ('%v', '%v', '%v', '%v')`, "property_info", pI.Address, pI.PictureLink, pI.OpenTime1, pI.OpenTime2)

	_, err := db.Exec(insertStmt)
	if err != nil {
		return nil, fmt.Errorf("rows is null or db.Query error: %v", err)
	}

	return pI, nil
}
