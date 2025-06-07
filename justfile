_default:
  @just --list

build:
  whiskers templates/android.tera
  whiskers templates/desktop.tera
  whiskers templates/ios.tera
  whiskers templates/macos.tera
