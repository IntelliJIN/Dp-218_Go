package entities

import "time"

type ScooterUploaded struct {
	Id 				  int    `json,scv:"id"`
	ModelId 		  int    `json,scv:"model_id"`
	OwnerId			  int 	 `json,scv:"owner_id"`
	SerialNumber 	  string `json,scv:"serial_number"`
	PaymentTypeId	  int 	 `json,scv:"name"`
	ModelName         string `json,scv:"model_name"`
	MaxWeight         int    `json,scv:"max_weight"`
	Speed 			  int    `json,scv:"speed"`
}

type Scooter struct {
	Id		 	 int	`json:"id"`
	ModelId  	 int	`json:"model_id"`
	OwnerId  	 int	`json:"owner_id"`
	SerialNumber string	`json:"serial_number"`
}

type ScooterModel struct {
	Id 		 		  int 	 `json:"id"`
	PaymentTypeId	  int    `json:"payment_type_id"`
	ModelName         string `json:"model_name"`
	MaxWeight         int    `json:"max_weight"`
	Speed 			  int    `json:"speed"`
}

type PaymentType struct {
	Id   int 	 `json:"id"`
	Name string  `json:"name"`
}

type SupplierPrices struct{
	Id 				int 	 `json:"id"`
	Price   		float64  `json:"price"`
	PaymentTypeId   int    	 `json:"payment_type_id"`
	UserId			int 	 `json:"user_id"`
}

type User struct {
	Id  		int 		`json:"id"`
	Email 		string 		`json:"login_email"`
	Blocked		bool 		`json:"is_blocked"`
	Name		string  	`json:"user_name"`
	Surname 	string  	`json:"user_surname"`
	CreatedAt 	time.Time   `json:"created_at"`
	RoleId 		int 		`json:"role_id"`
}

type Location struct {
	Id			int  	`json:"id"`
	Latitude	int  	`json:"latitude"`
	Longitude	int	 	`json:"longitude"`
	Label		string	`json:"label"`
}

type ScooterStation struct {
	Id			int  	`json:"id"`
	LocationId	int  	`json:"location_id"`
	Name		string	 `json:"name"`
	IsActive	bool	`json:"is_active"`
}