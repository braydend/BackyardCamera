# Garden Photo

This app will take photos of the garden from a RaspberryPi and send them to s3.

## Installation

### Prerequisites

1. Set up RaspberryPi with fresh Raspbian on the network
2. Install fswebcam `sudo apt-get install fswebcam`

### Cronjob

1. Copy the `logPhoto` file to `/home/pi`
2. Add AWS credentials to `.env`
3. Install the script in the cron table 
   1. `sudo crontab -e`
   2. `0 6,12,18 * * * cd /home/pi && ./logPhoto > /home/pi/cron.log 2>&1`

### Server

1. Copy the `server` and `server.service` files to `/home/pi`
2. Install the service in systemd
   1. `sudo cp /home/pi/server.service /lib/systemd/system`
3. Start service
   1. `sudo systemctl start server.service`

To take a photo, go to `{pi ip address}:8090/takePhoto`