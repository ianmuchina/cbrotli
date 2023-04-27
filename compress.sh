#!/bin/bash
set +x
set +e
DICT="./cdnjs.cloudflare.com/ajax/libs/react/17.0.2/umd/react.production.min.js"
FILE="./cdnjs.cloudflare.com/ajax/libs/react/18.2.0/umd/react.production.min.js"

./brotli -D $DICT -q 5 -S ".sbr" $FILE 
./brotli -q 5 -S ".br" $FILE  