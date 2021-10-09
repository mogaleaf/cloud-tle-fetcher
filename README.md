# satellite-leo-planning
Plan your low orbit satellite 

## tle_fetcher
Fetch tle from celestrak to register the last value in the db.

## deploying

This project is set to deploy the infra on aws with terraform.  
You need :
- aws cli
- terraform

First time
```
./init.sh
```

To deploy
```
./deploy.sh
```
