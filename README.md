# wallpaper sorter

Sort wallpapers if mobile or desktop

# Compile

```sh
$ go build -o bin/wallpaper-sorter
```

# Usage

Sort in current directory
```sh
$ wallpaper-sorter
```

Sort in specific directory
```sh
$ wallpaper-sorter -dir=/path/to/directory
```

Sort including wallpapers in subdirectories
```sh
$ wallpaper-sorter -r
```

Help command
```sh
$ wallpaper-sorter -h
```