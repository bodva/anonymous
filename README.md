#Anonymous Slack message from bot

##Config
Create your bot and place bot-token in .env file

##Instalation
```
go run anonymous.go -h
```

##Docker install
```
docker build -t go-anonymous-app .
docker run -ti go-anonymous-app anonymous -h
docker run -ti go-anonymous-app anonymous -channel="@bodva" -message="hello"
```