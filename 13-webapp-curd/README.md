|Abstract | Physical|
|-|-|
|package | folder|
||packages having `main()` can be converted to executable|
|module | collection of packages|

## How to create a module?
```
go mod init MODULE_NAME
```

## How to include our module into the current workspace?
```
 go work use .
 ```
 
 RestfulAPI
 ==========
 BROWSER, Another program ---HTTP--> SERVER (coded) 
 Scenario #1 add numbers
   Input: ctx.Fiber (query)
   Output: error (string)
 Scenario #2 resize
   Input: ctx.Fiber (image file)
   Output: error (image file)
 Scenario #3 object recognition
   Input: ctx.Fiber (image file)
   Output: error (string)

SAAS
====
Our Server (string, audio, video, image) ---HTTP/gRPC--> SAAS SERVER
