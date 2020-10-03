# metrics-collector-tado
collecting metrics from tado° thermostates and pushes them to an InfluxDB

## state of this project
- work in progress
- beta
- this is my first GOLANG project, so please do not judge me too hard

## what is does 
- It will connect to your tado° account using the credentials defined in the config.json
- It will fetch the current_temperature, requested_temperature and humidity from tado°'s API
- It will push the above metric to an InfluxDB using a DSN (url) defined in the config.json

## how to use it ?

copy the config.example.json file as config.json, open config.json an complete it with your settings

```
git clone git@github.com:sseidenthal/metrics-collector-tado.git
cp config.example.json config.json
```

### example config.json
```
{
    "tado_username": "me@example.com",
    "tado_password": "my-password",
    "tado_client_id": "tado-web-app",
	"tado_client_secret": "wZaRN7rpjn3FoNyF5IFuxg9uMzYJcvOoQ8QWiIqS3hfk6gLhVlG57j5YNoZL2Rtc",
    "interval": 60,
    "influxdb_dsn" : "http://username:password@uri:port/write?db=database_name"
}
```

### build & run
```
go build -o metrics-collector-tado main.go
./metrics-collector-tado
```
