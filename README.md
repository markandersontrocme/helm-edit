# Helm edit Plugin

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/MarkAndersonTrocme/helm-edit)](https://goreportcard.com/report/github.com/MarkAndersonTrocme/helm-edit)
[![CircleCI](https://circleci.com/gh/MarkAndersonTrocme/helm-edit/tree/main.svg?style=svg)](https://circleci.com/gh/MarkAndersonTrocme/helm-edit/tree/main)
[![Release](https://img.shields.io/github/release/MarkAndersonTrocme/helm-edit.svg?style=flat-square)](https://github.com/MarkAndersonTrocme/helm-edit/releases/latest)

`helm-edit` is an implemenation of `kubectl edit` for Helm values. It allows you to edit the values of a Helm release using a text editor.

## Install

Based on the version in the `plugin.yaml`, release binary will be downloaded from Github:

```console
Downloading and installing helm-edit v0.1.0 ...
https://github.com/MarkAndersonTrocme/helm-edit/releases/download/v0.1.0/helm-edit_0.1.0_darwin_amd64.tar.gz
Installed plugin: edit
```

### Editor
`helm-edit` uses the `EDITOR` environment variable for the text editor and defaults to `vim`.

## Usage

### Edit the values of a Helm release

```console
$ helm edit

edit helm values

Usage:
  edit [RELEASE] [flags]

Flags:
  -a, --all                get all values
      --dry-run            simulate upgrade command
  -h, --help               help for edit
  -n, --namespace string   namespace scope of the release
```
