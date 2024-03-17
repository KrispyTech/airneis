# Dive-in: Configure GORM

Author : Barnabé PILLIAUDIN  
Co-Assignee: Cédric Gautier  
Date : 2024-02-07

## SUMMARY:

We need an ORM system to be able to input data in our database.

## Context :

We want to declare a schema of our data models that will be applied to our database.
But we don't want it to be like dark magic. We want to have some control over it.
As we are using Golang as our backend language, we'll need to provide an ORM for our language.

## RESEARCH AND TESTS:

What we're looking for :

- A migration system
- A ORM to handle the logic required to interact with databases.
- A query builder that will build SQL statements from defining our models our backend language.

### Atlas

We had in mind to find a tool that was able to do versioned migration, that's where we found Atlas.
You can find this tool available with Homebrew or your linux package manager.

To setup Atlas, we'll need to create an configuration file called `atlas.hcl` in this following example :

We'll use an ORM called GORM and Atlas as our migration system.

```hcl

data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./model",
    "--dialect", "postgres", // | postgres | sqlite | sqlserver
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "Your dev database uri" # Don't forget to add sslmode=disable if ssl is not enable and search_path=public as queries string
  url =  "Your database uri" # In local it can be the same url
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
```

To create a migration run the following command

```shell
 atlas migrate diff  migration_name --env gorm
```

To apply the migration run

```shell
atlas migrate apply --env gorm
```

### Issues with Atlas

We had issues trying to rollback a migration or when trying to create a second migration errors occured.

![img.png](img.png)

For the rollback, the only way to reverse a migration without error was to clear the whole database.
To do so I used this command:

```shell
atlas schema clean -u "dburi"
```

But clearing the database means loosing all the data in it and that's not a satisfying solution in our case.

As for the second migration, I wasn't able to find a solution.
So we will need to use another tool to handle our migrations.

### GORM

Gorm is an ORM and a query builder for Go. It's known for it's simplicity and provides multiple features like an auto-migration system. It basically has all that we're looking for in one tool.

It has query building to ease up the task for schema creation, a practical API that wil help get make CRUD operations on our different models and has it's own migration system.

Let's start of by declaring our schema before creating a migration.

To declare a schema with gorm you can create a struct that represents a table where each attribute is a column.
It's a pretty nice and simple way of declaring a schema, and when operating on our data models we'll be able to use those structs as our types.

```go
package model

import "gorm.io/gorm"

type User struct {
gorm.Model
Name    string `json:"name" gorm:"unique"`
Age     int    `json:"age"`
IsAdmin bool   `json:"isAdmin"`
}

type Car struct {
gorm.Model
Owner       User   `json:"owner" gorm:"-"`
Constructor string `json:"constructor"`
ModelName   string `json:"modelName"`
}
```

# Solution

After some discussions within our team, and with the advice of Senior Developers. It seems that Gorm's automigrate system would be enough for our product. If needed we can still create SQL scripts to correct things that the automigrate would have done wrong.

The advantage of having auto-migrations is that Gorm will create tables, missing foreign keys, constraints, columns and indexes. It will change existing column’s type if its size, precision, nullable changed and It WON’T delete unused columns to protect our data.

## Useful links

- Atlas documentation: https://atlasgo.io/
- Gorm documentation: https://gorm.io/docs/

# Conclusion

We will be using GORM as our ORM, query builder and migration engine.
We won't have versioned migrations as we planed to have but automatic migrations.
It will be easier for us to work with and is enough for our usage.
