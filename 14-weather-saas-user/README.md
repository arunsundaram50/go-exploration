
Input address or city, state
Output weather

SAAS Service
============
Geocoding
https://get-lat-long?city=Tampa&state=FL
 { lat: "", lng: ""}

https://api.weather.gov/points/{latitude},{longitude}
 { temp: "20"}

 Various ways of getting intput into the program
 ===============================================
 | Program | Options |
 |-|-|
 |Command line| os.Args, flags, cobra|
 |WebApp|Fiber, net/http|
 |Any program|Database, S3 bucket, file (csvutil, JSON, txt), environment variable (os.Getenv("PATH"))|