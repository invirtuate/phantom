# phantom

Authorization middleware for [Fiber](https://github.com/gofiber/fiber).

## Usage

More details about this middleware can be found at [How to use authorization to secure REST APIs](http://invirtuate.com/blog/golang/how-to-use-authorization-to-secure-rest-apis).

This middleware by itself does not include any authorization callbacks, you must code them yourself.

Examples are in `internal/` and `cmd/example/main.go`.

### Install package

```console
go get -u codeberg.org/invirtuate/phantom
```

### Example using phantom middleware

```go
package main

import (
	"codeberg.org/invirtuate/phantom"
	"codeberg.org/invirtuate/phantom/internal/ipban"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Use(phantom.New(&phantom.Config{
		Callback: func(c fiber.Ctx) (int, error) {
			// check ip banned
			if err := ipban.Authorized("1.1.1.1"); err != nil {
				return fiber.StatusForbidden, err
			}

			return fiber.StatusOK, nil
		},
	}))

	app.Listen(":3000")
}
```

## License

This middleware is licensed under GPLv3.
