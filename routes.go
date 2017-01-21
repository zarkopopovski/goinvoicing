package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func RoutesMap(api *ApiConnection) Routes {
	var routes = Routes{
		Route{"Index", "GET", "/", api.Index},
		Route{"SignIn", "POST", "/signin", api.uHandlers.SignIn},
		Route{"Login", "POST", "/login", api.uHandlers.Login},
		Route{"NewProduct", "POST", "/newproduct", api.pHandlers.NewProduct},
		Route{"UpdateProduct", "POST", "/updateproduct", api.pHandlers.UpdateProduct},
		Route{"DeleteProduct", "POST", "/deleteproduct", api.pHandlers.DeleteProduct},
		Route{"ListProducts", "POST", "/listproducts", api.pHandlers.ListProducts},
		Route{"NewCustomer", "POST", "/newcustomer", api.cHandlers.NewCustomer},
		Route{"UpdateCustomer", "POST", "/updatecustomer", api.cHandlers.UpdateCustomer},
		Route{"DeleteCustomer", "POST", "/deletecustomer", api.cHandlers.DeleteCustomer},
		Route{"ListCustomers", "POST", "/listcustomers", api.cHandlers.ListCustomers},
		Route{"NewInvoice", "POST", "/newinvoice", api.NewInvoice},
		Route{"UpdateInvoice", "POST", "/updateinvoice", api.UpdateInvoice},
		Route{"DeleteInvoice", "POST", "/deleteinvoice", api.DeleteInvoice},
		Route{"ListInvoices", "POST", "/listinvoices", api.ListInvoices},
	}

	return routes
}
