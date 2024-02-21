# Go GUI

## wails

- https://wails.io/docs/gettingstarted/firstproject


- requirements
    - golang
    - nodejs
    - wails

- setup:

```ruby

task setup

# or
go install github.com/wailsapp/wails/v2/cmd/wails@latest

```

- new project:

```ruby


wails init -n myproject -t vue-ts


```

- run dev mode:

```ruby
cd project-root/

task gui:run:vue

```