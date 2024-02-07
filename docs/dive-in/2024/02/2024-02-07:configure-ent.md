# WIP NOT FINISHED

# Doubts that needs to be dispelled

- Cant ent interact with atlas schemas

# Configure ent and atlas dive-in

Author : Barnabé PILLIAUDIN
Date : 2024-02-07

## SUMMARY:
We want to declare a schema of our database and this schema to be applied to the database.
But we don't want it to be like a dark magic. We want to have some control over it.
As we are using golang for our backend we will be using ent and atlas to achieve that.

## SOLUTION:

Ent combined with atlas offers a clean to do versioned migrations.

Ent provides it own schema system, and it looks like this:
```go

package schema

import (
    'entgo.io/ent'
    "entgo.io/ent/schema/field"
)

// Fields of the User.
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.Int("age").
            Positive(),
        field.String("name").
            Default("unknown"),
    }
}
```

It not that bad but atlas has a better system

```sql
table "blog_posts" {
  schema = schema.example
  column "id" {
    null = false
    type = int
  }
  column "title" {
    null = true
    type = varchar(100)
  }
  column "body" {
    null = true
    type = text
  }
  column "author_id" {
    null = true
    type = int
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "author_fk" {
    columns     = [column.author_id]
    ref_columns = [table.users.column.id]
  }
}
```

which is the equivalent of this sql code 

```SQL
CREATE TABLE `blog_posts` (
  `id` int NOT NULL,
  `title` varchar(100) NULL,
  `body` text NULL,
  `author_id` int NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `author_fk` FOREIGN KEY (`author_id`) REFERENCES `example`.`users` (`id`)
);
```

Then we create a migration that will allow us to update the state of ou DB

#

## Useful links
- Atlas documentation: https://atlasgo.io/
- Ent documentations: https://entgo.io/

