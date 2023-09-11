Authentication Link
===================

http://www.strava.com/oauth/authorize?client_id=113411&response_type=code&redirect_uri=http://api.parkrun.online/connect&approval_prompt=force&scope=read_all,profile:read_all,activity:read_all

API Endpoint
============

1. Challenges

- Create Challenge: POST /v1/challenges
- Update Challenge: PUT /v1/challenges/{id}
- Get Challenge: GET /v1/challenges/{id}
- List Challenge: GET /v1/challenges
- Join Challenge: PUT /v1/challenges/{id}/join
- UnJoin Challenge: DELETE /v1/challenges/{id}/unjoin
- List Gamer per challenge: GET /v1/challenges/{id}/gamers
- List Longest run per day: GET /v1/challenges/{id}/longest-run-per-day

2. Activity

- Get Activity: GET /v1/activity/{id}

API Sample
==========

- Create Challenge

   ```curl -X POST -H "Authorization: Bearer v2.local.caHTS-lKZRmqQ-s1-bkgQytELrMXGoyG3jbT077bXXo6DvDfRW8snm8ogpex6_qXZ98QQ0FRLvcq8LAPivAjlATSDG8LKdFp1-aa23bitUtqTA.bnVsbA" -H "Content-Type: application/json" http://localhost:8080/v1/challenges -d '{"name": "park run2", "rules": "haha"}'```

- Join Game

   ```curl -X PUT -H "Authorization: Bearer v2.local.caHTS-lKZRmqQ-s1-bkgQytELrMXGoyG3jbT077bXXo6DvDfRW8snm8ogpex6_qXZ98QQ0FRLvcq8LAPivAjlATSDG8LKdFp1-aa23bitUtqTA.bnVsbA" -H "Content-Type: application/json" http://localhost:8080/v1/challenges/1/join -d '{"start_date": 1692476825, "end_date": 1695155225, "target": 30}'```

- Unjoin Game

   ```curl -X DELETE -H "Authorization: Bearer v2.local.caHTS-lKZRmqQ-s1-bkgQytELrMXGoyG3jbT077bXXo6DvDfRW8snm8ogpex6_qXZ98QQ0FRLvcq8LAPivAjlATSDG8LKdFp1-aa23bitUtqTA.bnVsbA" -H "Content-Type: application/json" http://localhost:8080/v1/challenges/1/unjoin```

- GetActivity

   ```curl -X GET -H "Authorization: Bearer v2.local.caHTS-lKZRmqQ-s1-bkgQytELrMXGoyG3jbT077bXXo6DvDfRW8snm8ogpex6_qXZ98QQ0FRLvcq8LAPivAjlATSDG8LKdFp1-aa23bitUtqTA.bnVsbA" -H "Content-Type: application/json" http://localhost:8080/v1/activity/112078641```

- LongestRunPerDay

   ```curl -X GET -H "Authorization: Bearer v2.local.caHTS-lKZRmqQ-s1-bkgQytELrMXGoyG3jbT077bXXo6DvDfRW8snm8ogpex6_qXZ98QQ0FRLvcq8LAPivAjlATSDG8LKdFp1-aa23bitUtqTA.bnVsbA" -H "Content-Type: application/json" http://localhost:8080/v1/challenges/1/longest-run-per-day```