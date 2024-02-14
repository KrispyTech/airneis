package neon

import (
    "database/sql"
    "fmt"

    "github.com/KrispyTech/airneis/lib/shared/vault"
    "github.com/pkg/errors"
    log "github.com/sirupsen/logrus"

    _ "github.com/lib/pq"
)

type Database struct {
    Neon *sql.DB
}

fun GetNeonCreds(vc vault.VaultClient)(username string, secret string, server string, err error) {
  username, err := vc.ReadSecret("neon_username")
  if err != nil {
    return NeonCred{}, errors.Wrapf(err, "GetNeonCreds, unable to get username")
  }
  secret, err := vc.ReadSecret("neon_secret")
  if err != nil {
    return NeonCred{}, errors.Wrapf(err, "GetNeonCreds, unable to get secret")
  }
  server, err := vc.ReadSecret("neon_server")
  if err != nil {
    return NeonCred{}, errors.Wrapf(err, "GetNeonCreds, unable to get server")
  }
  return 
}

func InitDB(vc vault.VaultClient) (Database, error) {
  username, secret, server, err := GetNeonCreds(vc)
  if err != nil {
    return Database{}, errors.Wrapf(err, "InitDB, unable to get credentials")
  }
  connStr := fmt.Sprintf("postgresql://%s:%s@%s/airneis?sslmode=require", username, secret, server)
  db, err := sql.Open("postgres", connStr)
  if err != nil {
       log.Fatal(err)
   }
   defer db.Close()
    return Database{Neon: db}, nil
}

func CheckVersion(db Database) (version string, err error) {
    rows, err := db.Neon.Query("select version()")
    if err != nil {
        log.Fatal(err)
    }

    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&version)
        if err != nil {
            log.Fatal(err)
        }
    }

    return fmt.Sprintf("version=%s\n", version), nil
}