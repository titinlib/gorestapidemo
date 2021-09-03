package routers

import cn "restapidemo/controllers"

type ProductRouter struct{}

func (productRouter ProductRouter) initialize() {
	productSubrouter := Router.NewRoute().Subrouter()

	productSubrouter.HandleFunc("/product", cn.GetAllProducts).Methods("GET")
	productSubrouter.HandleFunc("/product/{id}", cn.GetProductById).Methods("GET")
	productSubrouter.HandleFunc("/product", cn.CreateProduct).Methods("POST")
	productSubrouter.HandleFunc("/product/{id}", cn.DeleteProductById).Methods("DELETE")
	productSubrouter.HandleFunc("/product/{id}", cn.UpdateProductId).Methods("PUT")
}
