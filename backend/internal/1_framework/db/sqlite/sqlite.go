package sqlite

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"backend/internal/2_adapter/gateway"
	domain "backend/internal/4_domain"
	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("framework_driver", "sqlite")
	// DBのアップデート

	updateStoreDB()

	// defer storeDB.Close()

	// storeRows, err := storeDB.Query("SELECT * FROM products ORDER BY jan_code")
	// if err != nil {
	// 	log.Printf("Error querying products: %v", err)
	// }
	// defer storeRows.Close()

	// for storeRows.Next() {
	// }
}

func updateStoreDB() error {

	sqliteFilePath := getSqlitePath() + "/master.sqlite3"

	masterDB, err := open(30, sqliteFilePath)
	if err != nil {
		myErr.Logging(err)
		panic(err)
	}
	masterProductList := &domain.ProductList{}

	masterDB.Find(&masterProductList)

	// masterRows, err := masterDB.Query("SELECT * FROM products ORDER BY jan_code")
	// if err != nil {
	// 	log.Printf("Error querying products: %v", err)
	// 	return err
	// }
	// defer masterRows.Close()

	// masterProductList, masterJANCodeList, err := convertToStruct(masterRows)
	// if err != nil {
	// 	return err
	// }

	// storeDBName, err := getDBName()
	// if err != nil {
	// 	return err
	// }

	// storeDB, err := connectDB(storeDBName)
	// if err != nil {
	// 	return err
	// }
	// defer storeDB.Close()

	// storeRows, err := storeDB.Query("SELECT * FROM products ORDER BY jan_code")
	// if err != nil {
	// 	log.Printf("Error querying products: %v", err)
	// 	return err
	// }

	// defer storeRows.Close()
	// _, storeJANCodeList, err := convertToStruct(storeRows)
	// if err != nil {
	// 	return err
	// }

	// if reflect.DeepEqual(masterJANCodeList, storeJANCodeList) {
	// 	return nil
	// }

	// addProductList := []domain.Product{}

	// for _, masterJANCode := range masterJANCodeList {
	// 	if !slices.Contains(storeJANCodeList, masterJANCode) {
	// 		product := returnProduct(masterJANCode, masterProductList)
	// 		addProductList = append(
	// 			addProductList,
	// 			*product,
	// 		)
	// 	}
	// }

	// err = InsertProducts(storeDB, addProductList)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func getSqlitePath() string {

	currentPath, _ := os.Getwd()

	// localの場合は、db/engine/sqlite を確認
	sqlitePath := currentPath + "/db/engine/sqlite"
	// /home/yuji/Workspace/private_dev/template/backend/db/engine/sqlite/master.sqlite3
	if !exists(sqlitePath + "/master.sqlite3") {
		// dockerの場合は、/go/src/backend/sqlite を確認
		sqlitePath = filepath.Dir(currentPath) + "/backend/sqlite"
	}

	return sqlitePath
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

type (
	// Sqlite ...
	Sqlite struct {
		Conn *gorm.DB
	}

	// Vegetable ...
	Vegetable struct {
		ID    int
		Name  string
		Stock int
	}
)

// NewToSQLite ...
func NewToSQLite() gateway.ToSqlite {

	storeDBName, err := getDBName()
	if err != nil {
		return nil
	}

	sqliteFilePath := getSqlitePath() + "/" + storeDBName + ".sqlite3"

	conn, err := open(30, sqliteFilePath)
	if err != nil {
		myErr.Logging(err)
		panic(err)
	}

	s := new(Sqlite)
	s.Conn = conn

	return s
}

func open(
	count uint,
	dbPath string,
) (
	*gorm.DB,
	error,
) {

	db, err := gorm.Open(
		sqlite.Open(dbPath),
		&gorm.Config{},
	)
	if err != nil {
		if count == 0 {
			myErr.Logging(err)
			return nil, fmt.Errorf("retry count over")
		}
		time.Sleep(time.Second)
		// カウントダウンさせるようにする
		count--
		return open(count, dbPath)
	}

	return db, nil
}

func getDBName() (string, error) {
	currentPath, _ := os.Getwd()

	err := godotenv.Load(filepath.Join(currentPath, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
		return "", nil
	}

	return os.Getenv("SQLITE_FILE"), nil
}
