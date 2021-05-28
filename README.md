# share.tpl

Global Templating and Snippet sharing Manager
`share.tpl` is pronounced `ʃéər-tee-pee-él`

## Usage

### List template

```
$ tpl list
github.com/takutakahashi/clipboad/git/new_branch
github.com/takutakahashi/templates/project_template/go
github.com/awesome-tpl/good-template
...
```

### Describe template
```
$ tpl describe {clip name}
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

## Settings

### Configuration File

In HOME directory:
```
.tpl.yaml
.tpl/config.yaml
.local/tpl/config.yaml
```
