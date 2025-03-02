# Ezenv Golang

## TLDR

Ezenv uses golang generics to remove boilerplate surrounding retrieving ENV vars when initalising applications.

You specify a custom type, eg `type CorsAllowedOrigins []string` and it will map the semi-colon delimited env var
CORS_ALLOWED_ORIGINS into that variable, throwing a fatal if the env var isn't set.

It can be used in the context of Dependency Injection, or without DependencyInjection.

## Use case + example

Oftentimes, when using frameworks such as [Uber Fx](https://github.com/uber-go/fx), you'll provision your services 
at the top level, and then use dependency injection to wire the services together.

You'll probably have written code like this:

```go
// service.go

type CorsAllowedOrigins []string
type Port string

...

func NewHttpService(corsAllowedOrigins CorsAllowedOrigins, port Port) {
   ...
}

-----
// main.go

func GetPort() Port {
    port := os.Getenv("PORT")
    if port == "" {
        log.Fatalln("PORT env var not set.")
    }
    return port
}

func GetAllowedOrigins() CorsAllowedOrigins {
    allowedOriginsString := os.Getenv("CORS_ALLOWED_ORIGINS")
    if allowedOriginsString == "" {
        log.Fatalln("CORS_ALLOWED_ORIGINS env var not set.")
    }
    return strings.Split(allowedOriginsString, ";")
}

app := fx.New(
    fx.Provide(GetPort),
    fx.Provide(GetAllowedOrigins),
    fx.Provide(NewHttpService),
)
app.Run()

```

The first function GetAllowedOrigins is boilerplate. Ezenv removes the need for this boilerplate.

The above code becomes:

```go
// service.go

type CorsAllowedOrigins []string
type Port string

...

func NewHttpService(corsAllowedOrigins CorsAllowedOrigins, port Port) {
   ...
}

-----
// main.go

app := fx.New(
    fx.Provide(ezenv.Provider[Port]),
    fx.Provide(ezenv.SliceProvider[CorsAllowedOrigins]),
    fx.Provide(NewHttpService),
)
app.Run()

```

## How it works

The function `ezenv.SliceProvider[CorsAllowedOrigins]() ` returns a provider, which is a function that returns 2 output variables:
- a function which has the specified type as a return value
- secondly, an error



When it is passed CorsAllowedOrigins, it will look for an env var named CORS_ALLOWED_ORIGINS. This is "magic", and so 
it's important to remember to have naming consistency.

### `provider, err := ezenv.Provider[T]`

Return a custom type that is a builtin data type, errors if not set.

### `provider, err := ezenv.SliceProvider[T]`

Returns a custom type that is a slice of builtin data types, errors if not set.

#### Supported types

Currently has support for:

##### Builtins

int, int32, int64   
string
bool

##### Slices

[]int, []int32, []int64   
[]string
[]bool

## When not using dependency injection

```go
func TestParseStringArrayEnvVar(t *testing.T) {   
    os.Setenv("STRING_LIST", "Alice;Bob;Charlie")
	
    parts, err := ezenv.SliceProvider[StringList]()
    if err != nil {
        t.Error("Should not error")
    }

    if parts[0] != "Alice" || parts[1] != "Bob" || parts[2] != "Charlie" {   
        t.Error("parts slice elements should be {1, 2, 3}")   
    }
}
```