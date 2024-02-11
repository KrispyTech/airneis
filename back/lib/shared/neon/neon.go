package neon

import (
  "database/sql"
  "fmt"
  "log"
  "github.com/KrispyTech/airneis/config"
  
  _ "github.com/lib/pq"
)

func CheckVersion() {
  config, err := config.InitializeConfig()
	if err != nil {
		log.Fatal("Unable to load config - ", err.Error())
	}
  secret, err := config.Handler.VaultClient.ReadSecret("neon")
	if err != nil {
		println(err)
	}

  connStr := "postgresql://krispytech:" + secret +"@ep-misty-shape-a28kyc2k.eu-central-1.aws.neon.tech/airneis?sslmode=require"
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  rows, err := db.Query("select version()")
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  var version string
  for rows.Next() {
    err := rows.Scan(&version)
    if err != nil {
      log.Fatal(err)
    }
  }
  fmt.Printf("version=%s\n", version)
}