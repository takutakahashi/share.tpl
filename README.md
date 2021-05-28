# shclip
Global Clipboard Manager

## Usage

### List clip

```
$ clip list
github.com/takutakahashi/clipboad/git/new_branch
github.com/takutakahashi/clipboad/project_template/go
github.com/awesome-clip/good-clipboard/good-template
...
```

### Describe clip
```
$ clip describe {clip name}
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
.clip.yaml
.clip/config.yaml
.local/clip/config.yaml
```
