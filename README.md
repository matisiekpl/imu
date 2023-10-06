## IMU Dataset Receiver
Tiny microservice that catches CSV files and stores in FS.
### Endpoints:
```
GET /
POST /submit/:category
GET /download
```

### Installation
```bash
go install github.com/matisiepl/imu@latest
imu
```