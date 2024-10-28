#!/usr/bin/bash

FTP_HOST=$1
FTP_USER=$2
FTP_PASS=$3
LCD="/tmp/rc-upload"
RCD="/"

lftp -f "
set sftp:auto-confirm yes
open $FTP_HOST
user $FTP_USER $FTP_PASS
mirror --continue --reverse --verbose $LCD $RCD
bye
"
