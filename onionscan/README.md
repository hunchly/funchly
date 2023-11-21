
# OnionScan on Google Cloud Shell

NOTE! This repo is a patched version of https://github.com/s-rah/onionscan and is designed to be run from a Google Cloud Shell.

#### 1. To get a free cloud shell account: https://console.cloud.google.com/getting-started?pli=1

#### 2. Open your Cloud shell you click here:

![Screen Shot 2022-02-08 at 3 35 31 PM](https://user-images.githubusercontent.com/12143192/153079726-0332a2f7-8dae-4786-81c5-d70e71fc4dd1.png)

#### 3. From your Cloud Shell:

`wget https://raw.githubusercontent.com/hunchly/funchly/main/onionscan/cloudshell_install_onionscan.sh`

`chmod +x cloudshell_install_onionscan.sh`

`./cloudshell_install_onionscan.sh`

#### 4. You now need a way to connect into Tor, and we can use Docker (already setup in Google Cloud Shell) for this:

`docker run -it -p 127.0.0.1:9050:9050 –-name torproxy -d dperson/torproxy`

`docker inspect torproxy`

You will need to re-run the Tor Proxy / docker commands above each time you drop into a cloud shell.

#### 5. Note the IP address listed in the IPAddress field and then run onionscan:

`onionscan -torProxyAddress <IP ADDRESS OF DOCKER CONTAINER>:9050 -verbose <ONION ADDRESS>`

#### 6. Use the web preview in Google Cloud Shell to view the Onion Scan Correlation Lab:

![Screen Shot 2022-02-14 at 4 09 09 PM](https://user-images.githubusercontent.com/12143192/153955066-eb52a0ac-bf80-4fb2-be52-e123c5182106.png)

Your browser should open the Onion Scan Correlation lab in a new tab:

![Screen Shot 2022-02-14 at 4 13 06 PM](https://user-images.githubusercontent.com/12143192/153955318-e68827c3-e8d9-424d-8ea2-438cb4e0937a.png)


#### 7. Some common searches:

`snapshot` - finds all types of all data extracted.

`crawl` - finds all successful crawls.

`ssh` - find all SSH banners retrieved.

`email-address` - search for extracted email addresses.

`mod_status` - this will find hidden services potentially leaking the "real" IP address.

Or you can search for the Title of a Hidden Service from the Hunchly Dark Web report.

# What is OnionScan?

Head to the original repo: https://github.com/s-rah/onionscan

