package neon

import (
  "database/sql"
  "fmt"
  log "github.com/sirupsen/logrus"
  "github.com/KrispyTech/airneis/config"
  
  _ "github.com/lib/pq"
)

func CheckVersion() (string){
  config, err := config.InitializeConfig()
	if err != nil {
		log.Fatal("Unable to load config - ", err.Error())
	}
  username, err := config.Handler.VaultClient.ReadSecret("neon_username")
	if err != nil {
		println(err)
  }
  secret, err := config.Handler.VaultClient.ReadSecret("neon_secret")
	if err != nil {
		println(err)
	}
  server, err := config.Handler.VaultClient.ReadSecret("neon_server")
	if err != nil {
		println(err)
	}

  connStr := fmt.Sprintf("postgresql://%s:%s@%s/airneis?sslmode=require", username, secret, server)
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
  
  return fmt.Sprintf("version=%s\n", version)
}