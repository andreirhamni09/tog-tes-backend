package routes

import (
	"backend/delivery/bank_account_delivery"
	"backend/delivery/customer_delivery"
	"backend/delivery/deposito_delivery"
	"backend/delivery/pocket_information_delivery"
	"backend/repository/bank_account_repository"
	"backend/repository/customer_repository"
	"backend/repository/deposito_repository"
	"backend/repository/pocket_information_repository"
	"backend/usecase/bank_account_usecase"
	"backend/usecase/customer_usecase"
	"backend/usecase/deposito_usecase"
	"backend/usecase/pocket_information_usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(dbConn *gorm.DB) *gin.Engine {

	r := gin.Default()
	r.Use(CorsConfig())

	cust_repo := customer_repository.GetCustomerRepository(dbConn)
	cust_usec := customer_usecase.GetCustomerUsecase(cust_repo)
	cust_delivery := customer_delivery.GetCustomerDelivery(cust_usec)
	customer := r.Group("/customer")
	{
		customer.GET("/", cust_delivery.ShowAll)
		customer.GET("/:id", cust_delivery.ShowById)
		customer.POST("/", cust_delivery.CreateCustomer)
		customer.PUT("/:id", cust_delivery.UpdateCustomer)
		customer.DELETE("/:id", cust_delivery.DeleteCustomer)
	}

	bank_account_repo := bank_account_repository.GetBankAccountRepository(dbConn)
	bank_account_usec := bank_account_usecase.GetBankAccountUsecase(bank_account_repo)
	bank_account_delivery := bank_account_delivery.GetBankAccountDelivery(bank_account_usec)
	bank_accounts := r.Group("/bank_account")
	{
		bank_accounts.GET("/", bank_account_delivery.ShowAll)
		bank_accounts.GET("/:id", bank_account_delivery.ShowById)
		bank_accounts.POST("/", bank_account_delivery.CreateBankAccount)
		bank_accounts.PUT("/:id", bank_account_delivery.UpdateBankAccount)
		bank_accounts.DELETE("/:id", bank_account_delivery.DeleteBankAccount)
	}

	pocket_information_repo := pocket_information_repository.GetPocketInformationRepository(dbConn)
	pocket_information_usec := pocket_information_usecase.GetPocketInformationUsecase(pocket_information_repo)
	pocket_information_delivery := pocket_information_delivery.GetPocketInformationDelivery(pocket_information_usec)
	pocket_informations := r.Group("/pocket_information")
	{
		pocket_informations.GET("/", pocket_information_delivery.ShowAll)
		pocket_informations.GET("/:bank_account_id", pocket_information_delivery.ShowById)
		pocket_informations.POST("/", pocket_information_delivery.Create)
		pocket_informations.DELETE("/:id", pocket_information_delivery.Delete)
	}

	deposito_repo := deposito_repository.GetDepositoRepository(dbConn)
	deposito_usec := deposito_usecase.GetDepositoUsecase(deposito_repo)
	deposito_deliv := deposito_delivery.GetDepositoDelivery(deposito_usec)

	deposito := r.Group("/deposito")
	{
		deposito.GET("/", deposito_deliv.ShowAll)
		deposito.GET("/:id", deposito_deliv.ShowById)
		deposito.POST("/", deposito_deliv.CreateDeposito)
	}
	return r
}

func CorsConfig() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Access-Control-Allow-Origin"}
	return cors.New(config)
}
