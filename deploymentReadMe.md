**To run : WebApp on EC2**

1. ssh -i CreatedKey.pem ec2-user@54.175.125.107
2. sudo service docker start
3. docker run -p 80:3000 anirudhrv1234/reactjs

**EC2 - GO API IP:GET IP WHEN INSTANCE IS RUNNING**

1. Create an EC2 instance and allow HTTP:80 connections in the security options
2. chmod 400 goapikey.pem
3. ssh -i goapikey.pem ec2-user@<IP-Address>
4. sudo yum update -y
5. sudo yum install -y docker
6. sudo service docker start
7. sudo usermod -a -G docker ec2-user
8. exit
9. ssh -i goapikey.pem ec2-user@54.197.42.159
10. docker run -p 80:8080 anirudhrv1234/goapi
    _Important :_
    update atlas with the public IP of Go API for Permissions

**To run : Update atlas with the public key for Permissions**

1. ssh -i goapikey.pem ec2-user@34.227.160.149
2. sudo service docker start
3. docker run -p 80:8080 anirudhrv1234/goapi

# Resources :

1. https://docs.docker.com/docker-for-mac/
2. https://docs.docker.com/compose/gettingstarted/
3. https://medium.com/travis-on-docker/how-to-dockerize-your-go-golang-app-542af15c27a2
4. https://hostadvice.com/how-to/how-to-use-docker-containers-with-aws-ec2/
5. https://medium.com/fredwong-it/aws-cloudfront-url-return-accessdenied-in-front-of-s3-react-app-a4bad7d3e4f2
