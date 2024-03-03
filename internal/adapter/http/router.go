package http

import "github.com/labstack/echo/v4"

type Router struct {
	*echo.Echo
}

func NewRouter(transactionHandler TransactionHandler) (*Router, error) {
	router := echo.New()

	v1 := router.Group("/v1")
	{
		transaction := v1.Group("/transactions")
		{
			transaction.POST("", transactionHandler.CreateTransaction)
		}
	}

	return &Router{
		router,
	}, nil
}

func (router *Router) Serve(address string) error {
	return router.Start(address)
}
