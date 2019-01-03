module match-server

require (
	cloud.google.com/go v0.34.0 // indirect
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/go-redis/redis v6.15.0+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/builder v0.3.3
	github.com/go-xorm/core v0.6.0
	github.com/go-xorm/xorm v0.7.1
	github.com/google/go-cmp v0.2.0 // indirect
	github.com/smallnest/rpcx v0.0.0-20181228102150-00249c214b89
	github.com/streadway/amqp v0.0.0-20181205114330-a314942b2fd9
)

replace stathat.com/c/consistent => github.com/stathat/consistent v1.0.0
