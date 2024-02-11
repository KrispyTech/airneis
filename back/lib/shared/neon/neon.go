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

  connStr := "postgresql://" + username + ":" + secret + "@" + server + "/airneis?sslmode=require"
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