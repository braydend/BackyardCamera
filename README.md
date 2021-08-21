# Garden Photo

This app will take photos of the garden from a RaspberryPi and send them to s3.

## Installation

1. Set up RaspberryPi with fresh Raspbian on the network
2. Install fswebcam `sudo apt-get install fswebcam`
3. Copy the files to `/home/pi`
4. Add AWS credentials to `takePhoto.sh`
5. Install the script in the cron table 
   1. `sudo crontab -e`
   2. `0 6,12,18 * * * /home/pi/takePhoto.sh > /home/pi/cron.log 2>&1`