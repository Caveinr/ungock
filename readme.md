# ungock

server side to drive micro servo in order to open door,
in response of different auth ways.

# Table structure

auth_uid: to store ```UID``` and ```SAK``` of qualified cards.
```mysql
CREATE TABLE `auth_uid` (
  `UID` varchar(10) NOT NULL,
  `SAK` tinyint(3) unsigned NOT NULL,
  `Note` varchar(16) DEFAULT NULL,
  PRIMARY KEY (`UID`,`SAK`)
) DEFAULT CHARSET=utf8mb4
```
auth_log: to store auth way ```METHOD``` and ```REQUESTER``` 
who  is requesting to open door.
```mysql
CREATE TABLE `auth_log` (
  `ID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `REQUESTER` varchar(16) NOT NULL,
  `METHOD` enum('MFRC522','GO_CQHTTP') NOT NULL,
  `STATUS` enum('PASSED','FAILED') NOT NULL,
  `TIME` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`ID`)
) DEFAULT CHARSET=utf8mb4 
```

# API

## MFRC522

```
POST /MFRCScan
Provide the UID+SAK of a rfid card and unix timestamp,
calling to open door.

Parameters:
No parameters

Header:
X-signature: hmac with sha1 of body.

Request body: application/json
Example value:
{
  "UID":"AABBCCDD",
  "SAK":8,
  "timestamp":1711025917
}

Responses:
Code: Description
200: operation successed.
400: can`t bind request.
401: auth failed.
403: not a qualified card.
```

## CQ-HTTP 
[Deprecated]
```
POST /cqhttp
using qq bot to request to open door.

Parameters:
No parameters

Header:
X-signature: hmac with sha1 of body.

Request body: application/json

Responses:
200: operation successed.
```
# Config file
config.toml:
```toml
PORT = ":port"

[CQHTTP]
# hmac key.
SECRET = ""
# anyone in those qq group will pass auth.
GROUP_ID = []
# key word to let qqbot to request opening door.
KEY_WORD = ""
# if passed, bot will reply:
FAST_REPLAY_OK = ""
# if failed, bot will reply:
FAST_REPLAY_NO = ""

[MFRCScan]
# hmac key
SECRET = ""
# [not implement] 
# max seconds different from the time of the request and the current.
TOLERANCE = 5

[GORM]
DSN = "user:password@tcp(url:port)/database?charset=utf8mb4&parseTime=True&loc=Local"
```
