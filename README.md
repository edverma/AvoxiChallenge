Avoxi Challenge

This project includes the GeoLite country database, implementation for a mongodb database, and an api endpoint to authorize login based on the client's IP address.

The database holds the whitelisted countries' ISO codes. When a client performs the /login GET request, its IP address is captured, and the IP address is passed to a GeoLite API endpoint which returns the country info of the IP address. This country info contains the ISO country code which corresponds to this IP address. If this ISO country code is in the whitelist, then the server responds with a string "authorized." Otherwise, the server responds with the string "login failure: client ip address is unauthorized."

My (very general) plan for keeping the GeoLite database updated is to create a job on the system that this server runs on. This job would run nightly and would delete our existing GeoLite database and download it and decompress it from the permalink GeoLink provides: https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country&license_key=YOUR_LICENSE_KEY&suffix=tar.gz.