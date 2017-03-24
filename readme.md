## Routes

Method | Path | Description | Query string | Request headers | Request body | Response
------ | ---- | ----------- | ------------ | --------------- | ------------ | --------
GET | /games | Get list of gamses | --- | --- | --- | {"data": games list JSON}
GET | /games/:id | Get full game info | --- | --- | --- | {"game": game info JSON}
POST | /games | Create game | --- | x-token | {"title": "Game title", "username": "User Name"} | {"game": game JSON, "access_token": ObjectID}
POST | /games/:id/join | Join to new game | --- | x-token | {"username": "User Name"}  | {"game": game JSON, "access_token": ObjectID}
POST | /games/:id/move | Make move | --- | x-token | {"row": "1", "col": "2"}  | {"field": array with field, "winnerIndex": winnerIndex, "winnerName": winnerName}

## JSON examples

### Response format

``` javascript
{
    "status": ["ok", "error"],
    ["data"]: {some data},
    ["error"]: {error JSON}
}
```

### Games list

``` javascript
[
    {"_id":"58cb7e6370e544685b3431bd","title":"Первая игра","status":0},
    {"_id":"58cbb98cf90fbe084c1f5a23","title":"Title","status":10},
    {"_id":"58d25a2ecb47275a068f6f32","title":"Новая игра","status":20}
]
```

### Game info

``` javascript
{
    "_id":"58cbb98cf90fbe084c1f5a23",
    "title":"Title",
    "status":10,
    "poleSize":3,
    "currentTurn":1,
    "winnerIndex":0,
    "winnerName":"",
    "field":[
        [{"state":1},{"state":1},{"state":2}],
        [{"state":2},{"state":2},{"state":2}],
        [{"state":2},{"state":2},{"state":0}]
    ],
    "createdDate":"2017-03-17T13:25:16.567+03:00",
    "player1id":"58cb7e6370e544685b3431bc",
    "player2id":"58d2648bcb47275a068f6f3a",
    "player1":{
        "_id":"58cb7e6370e544685b3431bc",
        "name":"Максим"
    },
    "player2":{
        "_id":"58d2648bcb47275a068f6f3a",
        "name":"Петя"
    }
}
```