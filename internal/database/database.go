package database

var DB *gorm.DB

func ConnectDatabase(cfg *config.Config) (*gorm.DB, error) {
  database, err := gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{})

  if err != nil {
    log.Fatal("Failed to connect to Database:", err.Error())
  }

  fmt.Println("Connected to Database ✅")

  // auto migrate
  database.AutoMigrate(
    &models.Comment{},
    &models.Post{},
    &models.Reaction{},
    &models.Tags{},
    &models.User{},
  )

  DB = database
  return database, nil
}
