# swdcom_gdbmi

reimplementation of swdcom (https://github.com/Crest/swdcom) with use of gdb machine interface, so you can use your prefered debugger instead of stlink.

golang was choosen to implement this access method later into folie (https://git.jeelabs.org/jcw/folie)

## current status

just testing the connection to debugger with go via gdb mi 

## install

### install gdb mi lib for golang
go get github.com/cyrus-and/gdb

