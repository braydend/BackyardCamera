#!/usr/bin/env bash

DATE=$(date +%s)
FILENAME="$DATE".jpg
BUCKET=backyard-photos

# Set up AWS credentials
export AWS_ACCESS_KEY_ID=
export AWS_SECRET_ACCESS_KEY=

fswebcam --delay 2 "$FILENAME"
# shellcheck disable=SC2094
/home/pi/s3Uploader-ARM -b "$BUCKET" -k "$FILENAME" -d 5m < "$FILENAME"
rm "$FILENAME"
