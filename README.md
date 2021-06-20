# snip

Global Templating and Snippet sharing Manager

## Key Features

This software have 2 key features.

### 1. Repository-based snippet sharing

You can share your own snippets with any other peaple, like co-worker.

```
% snip our-project/operation/list-users
mysql -ufoo -pbar -h db.our-project.com -e "SELECT * from users;" our-project
```

### 2. Embed values to your snippets

You can define and embed values to snippets.

```
% snip list
snippets/hello-world

% snip show snippets/hello-world

  Description: hello snippet
  Embedded values:
    user ... user name, default: alice

% snip snippets/hello-world
Hello! alice

% snip --set user=bob snippets/hello-world
Hello! bob
```

## Usage

### Write Configuration file

Configuration file can be located on HOME directory:
```
.snip.yaml
.snip/config.yaml
.local/snip/config.yaml
```

or specified by commandline option

```
snip --config path/to/config.yaml list
```

Simple example is below:

```
setting:
  basedir: "path/to/snippets_bin" # default: $HOME/.snip
repositories:
  - name: snippets
    type: manual
  - name: snippets-from-git
    type: git
    uri: https://github.com/takutakahashi/snippets.git
```

### Configurations for Snippet Repository

Snippet Repository is a set of snippets and metadata config.
It can be served by git repository now.

Snippet Repository must have two type of configuration file.

1. root config on top of the repository
2. snippet config on each snippet directories

#### Root Config

Snippet Repository needs to have root config named `.root.snip.yaml` on top of repo. ex: repository root of `takutakahashi/snippets`.

Example of `.root.snip.yaml` is below.

```
snippets:
  - name: goreleaser
  - name: github-workflow/controller-release
  - name: go-cli
  - name: new-snip
```

`name` must be a relative path from snippet repository root.

#### Snippet Config

Snippet is a set of files and a snippet config.
If you want to set `foobar` snippet on your repo, you need to locate some file and `.snip.yaml` file. like below:

```
% cat foobar/.snip.yaml
description: "foobar template"
values:
- name: name
  describe: name of user
  default: alice
```

## All Features

### List template

```
% snip list
takutaka/goreleaser
takutaka/github-workflow/controller-release
takutaka/go-cli
takutaka/new-snip
...
```

### Describe snippet

```
$ snip describe {snippet name}
Title: awesome snippet
Description: This is VeryVery Good snippet.
Argments:
- name: username
  desc: good arg
  default: alice
- name: email
  desc: very good arg
  default: bob@bob.com
Data: |
Hello @@(username)! your email address is @@(email).
```

### Export snippet to the specified directory

```
% snip --output . takutaka/goreleaser
```

### Export snippet to stdout

When the snippet is a set of `.snip.yaml` and only the name of file `snippet`, it exports to stdout.
```
% ls misc/snippets/single -la
total 16
drwxrwxr-x 2 owner owner 4096 Jun 20 02:26 ./
drwxrwxr-x 4 owner owner 4096 Jun 20 02:26 ../
-rw-rw-r-- 1 owner owner   93 Jun 20 02:26 .snip.yaml
-rw-rw-r-- 1 owner owner   24 Jun 20 02:26 snippet

% snip --config misc/config_test.yaml show snippets/single

  Description: test template
  Embedded values:
    name ... no Description, default: alice

% snip --config misc/config_test.yaml snippets/single
hello alice
hello

```
