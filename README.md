Go Environment Loader
---

This package provides a simple way to load environment variables from .env files into your Go application, and an atomic API to read and write environment variables.

It supports basic value types such as strings, integers, booleans, and floats.

# About

The application uses a global environment variable for the general application environment: `ENV`.
The value can be overwritten by calling `SetEnv`. Common environment shorthands are automatically parsed. The main environment types are:
- development
- production
- testing
- staging

The package comes with some helper functions to check the global environment, e.g. `IsDevelopment()`.
Feel free to use a custom environment type, fitting your application's needs.

## DotEnv

This package relies on [joho/godotenv](https://github.com/joho/godotenv) to load environment variables from .env files, adding some syntactic sugar to make it easier to overload (custom) files into your application environment.

Files overload each other in the following order:
- .env
- .env.local
- .env.<app_env>
- .env.<app_env>.local
- <custom_file>

## Manual injection

You can inject variables using the `Set` function. This supports all aforementioned basic data types.

# Examples

```go
package main

import (
    env "github.com/vpmv/goenv"	
)

func main() {
    env.LoadDotEnv(`/config/`)
    
    if env.IsDevelopment() {
        env.Set(`SEED_DB`, true)
    }
    
    database := env.MustString(`DATABASE_URL`) // will panic if unset
    databasePort := env.GetInt(`DATABASE_PORT`, 3306) // will return default value (3306) if unset
    // ...
}
```
