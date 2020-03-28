# script-output-converter
It convert script command output (typescript) to plain text.

* It can handle deleting chars operation.
* It can handle multibytes chars.
* It deletes NULL chars.
* It deletes contiguous 5 spaces for cleaning vi output.

## Usage

#### Install Library

CentOS

```
# yum install libvterm
```

Ubuntu

```
# apt-get install libvterm0
```

#### Install script-output-converter

```
# wget https://github.com/t-matsuo/script-output-converter/releases/download/0.1/script-output-converter
# chmod 755 script-output-converter
```

#### run

```
# ./script-output-converter typescript
```

## Building

You need Docker environment.

```
# ./build.sh
```

