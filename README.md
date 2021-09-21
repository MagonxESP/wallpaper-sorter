# wallpaper sorter

Sort wallpapers if mobile or desktop

# Compile

```sh
$ go build -o bin/wallpaper-sorter
```

# Usage

### Sort in current directory
```sh
$ wallpaper-sorter
```

### Sort in specific directory
```sh
$ wallpaper-sorter -dir=/path/to/directory
```

### Sort including wallpapers in subdirectories
```sh
$ wallpaper-sorter -r
```

#### Ignore specific subdirectories

You can create an ```.wallpapersorterignore``` file and use it for ignore specific subdirectories and files you won't process.

This file should be located in the root directory you will execute this program (or the directory specified with the -dir option).

### Help command
```sh
$ wallpaper-sorter -h
```