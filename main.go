package main

import (
	"flag"
	"fmt"
)

var err error

func main() {
	

	// database.Db.AutoMigrate(&model.ProductCourse{})
	// database.Db.AutoMigrate(&model.Grade{})
	// database.Db.AutoMigrate(&model.ProductCourseRelationship{})
	// err := godotenv.Load()
	// if err != nil {
	// 	panic("Failed To Load Env")
	// }
	// lib.JWTSignatureKey = os.Getenv("JWT_SIGNATURE_KEY")
	// lib.JWTIssuer = os.Getenv("JWT_ISSUER")
	// println(os.Getenv("JWT_SIGNATURE_KEY"))
	// server := routes.InitRoutes()
	// server.Run(":8000")
	
	migrate := flag.String("m", "", "Unsupport Command")
	flag.Parse()
	command  := *migrate
	if command != "abc" && command == "migrate" {
		// dbEnv := database.DBUrl(database.BuildDB())
		// database.Db, err = gorm.Open(mysql.Open(dbEnv), &gorm.Config{})
		// if err != nil {
		// 	panic("Failed to connect database")
		// }
		// database.Db.AutoMigrate(&model.Grade{})
		fmt.Printf("oke")
		return
	}
	fmt.Printf("Choose %s", *migrate)
}
