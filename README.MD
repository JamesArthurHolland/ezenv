# Ezenv Golang

## TLDR

Ezenv uses golang generics to remove boilerplate surrounding retrieving ENV vars when initalising applications.

You specify a custom type, eg `type CorsAllowedOrigins []string` and it will map the semi-colon delimited ENV var 
into that variable, throwing a fatal if the env var isn't set.

## Use case + example

Oftentimes, when using frameworks such as [Uber Fx](https://github.com/uber-go/fx), you'll provision your services 
at the top level, and then use dependency injection to wire the services together.

You'll probably have written code like this:

```
type CorsAllowedOrigins []string

func GetAllowedOrigins() CorsAllowedOrigins {
    allowedOriginsString := os.Getenv("CORS_ALLOWED_ORIGINS")
    if allowedOriginsString == "" {
        log.Fatalln("CORS_ALLOWED_ORIGINS env var not set.")
    }
    return strings.Split(allowedOriginsString, ";")
}

func NewHttpService(corsAllowedOrigins CorsAllowedOrigins, ...) {
   ...
}

...
	app := fx.New(
		fx.Provide(GetAllowedOrigins),
		fx.Provide(NewHttpService),
	)
	app.Run()

```

The first function GetAllowedOrigins is boilerplate. Ezenv removes the need for this boilerplate.

The above code becomes:

```
type CorsAllowedOrigins []string

func NewHttpService(corsAllowedOrigins CorsAllowedOrigins, ...) {
   ...
}

...

app := fx.New(
    fx.Provide(ezenv.Provider[CorsAllowedOrigins]()),
	fx.Provide(NewHttpService),
)
app.Run()

```

## How it works

The function ezenv.Provider[CorsAllowedOrigins]() return a provider, which is a function that returns a function which
has the specified type as a return value.

When it is passed CorsAllowedOrigins, it will look for an env var named CORS_ALLOWED_ORIGINS. This is "magic", and so 
it's important to remember to have naming consistency.