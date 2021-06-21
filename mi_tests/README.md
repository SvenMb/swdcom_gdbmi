## Start

start your debugger openocd if needed. (st-util for stlink v2)

go run main.go

## What?

this should just print out the register r11 from your attached mecrisp-stellaris board, after 1 second running time. 

Also it calculates the start of rx and tx buffer and the current state of the rxtx bytes.

lastly it prints out the current content of the TX-buffer

