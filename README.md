# hasGoriraAPI

API server made as a trial

## Set UP

```
1. make up
2. make migrate
3. make seed
```

## Response Data


```
curl localhost:5050/zoos
=>

[
 {"Name":"北海道ゴリラ大学","PrefName":"北海道","HasGorira":true},
 {"Name":"ぽんぽこ動物園","PrefName":"千葉県","HasGorira":false},
 {"Name":"わにわに博物館","PrefName":"岩手県","HasGorira":false},
 {"Name":"ゴリラの里","PrefName":"岡山県","HasGorira":true},
 {"Name":"白鳥の森","PrefName":"静岡県","HasGorira":false}
]

```
