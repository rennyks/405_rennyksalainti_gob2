package main

import (
	"GORM/database"
	"GORM/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()
	//createUser("ren@gmail.com")
	//getUserById(1)
	//updateUserbyId(1, "jhonjohn@gmail.com")
	//createProduct(1,"uniqlo", "kemeja")
	//getUSerWtProducts()
	deleteProductbyId(1)


}

func createUser(email string) {
	db := database.GetDB()

	User := models.User{
	Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("error creating user data:", err)
		return
	}

	fmt.Println("new user data", User)
}


func getUserById(id uint) {
	db := database.GetDB()

	user := models.User{}

	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("user data not found")
			return
		}
		print("error finding user:", err)
	}

	fmt.Printf("user data: %+v \n", user)
}


func updateUserbyId(id uint, email string) {
	db := database.GetDB()

	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error

	if err != nil {
		fmt.Println("error updating user data:", err)
		return
	}

	fmt.Printf("update user's email: %+v \n", user.Email)
}


func createProduct(userId uint, brand string, name string) {
	db := database.GetDB()

	Product := models.Product{
		UserID: userId,
		Brand: brand,
		Name: name,
	}

	err := db.Create(&Product).Error

	if err != nil {
		fmt.Println("error creating product data:", err.Error())
		return
	}

	fmt.Println("new product data:", Product)
}


func getUSerWtProducts() {
	db := database.GetDB()

	users := models.User{}
	err := db.Preload("Products").Find(&users).Error


	if err != nil {
		fmt.Println("error getting user datas with products:", err.Error())
		return
	}

	fmt.Println("user datas with products")
	fmt.Printf("%+v", users)
}

func deleteProductbyId(id uint) {
	db := database.GetDB()

	product := models.Product{}
	err := db.Where("id = ?", id).Delete(&product).Error


	if err != nil {
		fmt.Println("error deleting products:", err.Error())
		return
	}

	fmt.Printf("product with id %d has been successfully delete", id)
}