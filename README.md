AuthURL
=======
http://www.strava.com/oauth/authorize?client_id=113411&response_type=code&redirect_uri=http://localhost:8080/connect&approval_prompt=force&scope=read_all,profile:read_all,activity:read_all

API sample:
1. Join Game
curl -X PUT -H "x-user-info: 112078641" -H "Content-Type: application/json" http://localhost:8080/v1/game/join -d '{"start_date": 1694348963, "end_date": 1695212963, "target": 15000}'
2. Unjoin Game 
curl -X DELETE -H "x-user-info: 112078641" -H "Content-Type: application/json" http://localhost:8080/v1/game/unjoin
3. GetActivity