# ds18b20-viewer 

ds18b20-viewer is a programm that continuously displays the current values of connected ds18b20 sensors.

It's great for  individual sensors by watching temperatre changes.


## Installation

```sh
$ go get -v github.com/th3noname/ds18b20-viewer
```
### Cross-Compile for ARM

```sh
$ GOOS=linux GOARCH=arm GOARM=5 go build
```