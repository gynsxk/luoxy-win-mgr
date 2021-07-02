# luoxy-winmgr
Windows Manager Tools


## introduce

I have some windows, deploy some different service on different user.
for me, upgrade and manage service is tedious, so i consider to develop a tool for me.

## functions

this tool is like ansible, but can do some thing only need by me.

 * gRPC 
 * config center
 * runas another user to start a service and communicating to parent.
 * monitor
 ...
 
 ## task format
 
 ``` json
 [
    {
        "action": "download",
        "url": "http://192.168.0.111/somefile",
        "filename": "somefile",
        "dir": "test"
    },
    {
        "action": "unzip",
        "file": "",
        "dir": ""
    },
    {
        "action": "runas",
        "user": "test",
        "password": "test",
        "service": "serviceName",
        "program": "",
        "argv": []
        
    }
]
 ```

