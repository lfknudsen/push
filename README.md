# Push

Executes `git push` on each remote registered in the current Git directory.

Due to a common Forgejo or Codeberg bug, a failed push will be retried once.

## Dependencies
Requires `git` to be installed and on your PATH.

## Compile

```shell
go build
```

## Usage

```shell
push [OPTIONS]

Options:
-tag <tag name>             Push the given tag to each remote repository
-h                          Print this message and exit.
```

## See also

[Repo](https://gitlab.com/lfknudsen/repo): A programme to automatically create repositories\
with GitLab, GitHub, and/or Codeberg and register them as remotes, all with a single command.\
Also supports editing project description and visibility.


[Install](https://gitlab.com/lfknudsen/install): A programme to easily install newly-compiled binaries.\
It takes a file (if none is given, it will try to find an executable file with the\
same name as the current directory) and moves it to the same directory its\
own executable is in.
