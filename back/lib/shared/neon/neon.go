package neon

import (
	"database/sql"
	"fmt"

	"github.com/KrispyTech/airneis/lib/shared/vault"
	"github.com/pkg/errors"

	_ "github.com/lib/pq"
)

type Database struct {
    Neon *sql.DB
}

type NeonCred struct {
  Username string
  Secret string
  Server string
}

func GetNeonCreds(vc vault.VaultClient)(NeonCred, error) {
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

  return NeonCred{
    Username: username, 
    Secret: secret, 
    Server: server,
    }, nil
}

func InitDB(vc vault.VaultClient) (Database, error) {
  neonCred, err := GetNeonCreds(vc)
  if err != nil {
    return Database{}, errors.Wrapf(err, "InitDB, unable to get credentials")
  }
  connStr := fmt.Sprintf("postgresql://%s:%s@%s?sslmode=require", neonCred.Username, neonCred.Secret, neonCred.Server)
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    return Database{}, errors.Wrapf(err, "InitDB, unable to open the DB")
  }

  return Database{Neon: db}, nil
}

func CheckVersion(db Database) (version string, err error) {
  rows, err := db.Neon.Query("select version()")
  if err != nil {
    return "", errors.Wrapf(err, "CheckVersion, unable to query the version")
  }

  defer rows.Close()

  for rows.Next() {
    err = rows.Scan(&version)
    if err != nil {
      return "", errors.Wrapf(err, "CheckVersion, unable to scan the version")
    }
  }

  return fmt.Sprintf("version=%s\n", version), nil
}