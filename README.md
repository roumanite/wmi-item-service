### Where's My Item? Item Service

Where's My Item? is a social network mobile app created to help people develop a healthier lifestyle in their shopping
choices and space management by keeping track of what they possess and where the items are, how
much they spent on new items, and comparison of purchase vs discard every month.

### Package structure
The packages are structured according to hexagonal architecture that I learnt from this article: https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3

## How to deploy
ECS, S3, and RDS are covered in AWS Free Tier.

### Run docker
```
docker build -f ./docker/Dockerfile . -t wmi:dev --build-arg UID=$UID
docker run -d --name=wmi-development --publish 8080:8080 <image id>
```

### Debug
You can connect to ECS and RDS from your local machine just to play around. Ideally, RDS should only be accessed by the ECS.

#### Connect to ECS
```
ssh -i "file.pem" <instance address>
```

#### Connect to RDS
```
psql --host=... --username=... --password
```
Follow the explanation here if there's operation timed out error: https://medium.com/overlander/connecting-to-rds-from-local-over-tcp-operation-timed-out-5cfc819f402c

### Running the server manually
Can be accomplished using task definition on ECS.
