#!/bin/bash
./gasweb &> runtime.log
FILENAME="runtime_$(date +%F-%T).log"
mv runtime.log $FILENAME
curl --form "uploadfile=@$FILENAME" ktrllogtest.mybluemix.net
