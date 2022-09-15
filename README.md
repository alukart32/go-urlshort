# Redirect urls
The goal is to create an http.Handler that will look at the path of any incoming web request and determine if it should redirect the user to a new page, much like URL shortener would.

For instance, if we have a redirect setup for /dogs to https://www.somesite.com/a-story-about-dogs we would look for any incoming web requests with the path /dogs and redirect them.

Redirects rules can be stored in
- YAML (default: config.yaml)
- JSON
- db

Original task: https://courses.calhoun.io/lessons/les_goph_04

## Redirect file configuration
The rules for redirecting requests are located in the assets directory.

Depending on the input choice: yaml or json, the file has /assets/config. there will be a corresponding extension - /assets/config.yaml.

The default input method is the yaml file **_/assets/config.yaml_**.

The files contain data according to the following structure:

```
paths: [ path1{path, url}, ... ]
```

## Redirect db configuration
The default implementation is made for PostgreSQL using database/sql packages and github.com/lib/pq.

The data is stored in a table with the following structure:
```
CREATE TABLE IF NOT EXISTS public.redirects
(
    id integer NOT NULL,
    path character varying(50) COLLATE pg_catalog."default" NOT NULL,
    url character varying(120) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT redirect_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.redirects
    OWNER to postgres;
```

## Run options
The file main.go is located in **_/cmd_**.

### Build command:
```
go to the build -o redirect.
```

### Help command:
```
./redirect --help
```

### Launch Command
```
./redirect --input yaml --filepath "../assets/config.yaml"
```

This command will output:
```
2022/09/15 13:10:52 start decode the yaml file...
2022/09/15 13:10:52 the yaml file has been decoded...

config
   paths:
        - path: /go
          url: https://go.dev/
        - path: /go_pkg
          url: https://pkg.go.dev/
        - path: /path3
          url: http://localhost:8080/redirections/path3

init a new http handler...
start http server at http://localhost:8080...
```

```
./redirect --input db
```

This command will output:
```
2022/09/15 12:54:27 init the new db connection pool...
2022/09/15 12:54:27 db connection pool was created...
2022/09/15 12:54:27 fetch all redirects from db...

config
   paths:
        - path:
          url:
        - path: /go
          url: https://go.dev/
        - path: /go_pkg
          url: https://pkg.go.dev/
        - path: /path3
          url: http://localhost:8080/redirections/path3

init a new http handler...
start http server at http://localhost:8080...
```