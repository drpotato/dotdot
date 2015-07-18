# .dot

Versioning your dotfiles is a pain. A common solution is to put them in a separate directory, create symbolic links to
your home directory and commit the contents to a version control system.

Enter .dot, a tool to manage the symbolic links of your dotfiles. 

## Installing
Pretty straight forward
```
go get github.com/drpotato.dot
```

## Contributing
DotDot uses [gb](https://getgb.io) as a build system and it's vendor plugin to manage dependencies.

## Why Go?
Begining to use it at work and wanted to have a play with it.