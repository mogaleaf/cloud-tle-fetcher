# Cloud TLE fetcher solution
Be notified when a new TLE has been issued

## tle_fetcher_solution
Be notified when a new tle for your satellite is available.  
Deploy the solution on aws.  
Use the client to test the solution  


![image](schema.png)

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

## cloud

- dynamodb
- lambdas
- api gateway
- cloud watch
