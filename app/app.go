package app

import (
	"go-banking/domain"
	"go-banking/service"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

const (
	db_driver = "DB_DRIVER"
	db_source = "DB_SOURCE"
)

var (
	DBDriver = os.Getenv(db_driver)
	DBSource = os.Getenv(db_source)
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	//wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	dbClient := getDBClient()
	customerRepositoryDB := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDB := domain.NewAccountRepositoryDB(dbClient)
	transactionRepositoryDB := domain.NewTransactionRepositoryDB(dbClient)
	// accountRepositoryDB := domain.NewAccountRepositoryDB(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDB)}
	ah := AccountHandler{service: service.NewAccountService(accountRepositoryDB)}
	th := TransactionHandlers{service: service.NewTransactionService(transactionRepositoryDB,accountRepositoryDB)}

	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", th.MakeTransaction).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func getDBClient() *sqlx.DB {
	client, err := sqlx.Open(DBDriver, DBSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}