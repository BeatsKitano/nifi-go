
数据库一般都是公共访问，如有必要，data放置到app/your_app 目录下也可


docker run --rm --device /dev/net/tun --cap-add NET_ADMIN -ti -e PASSWORD=175544 -e URLWIN=1 -e DISABLE_PKG_VERSION_XML=1 -v /Users/xiaohui/Docker/ec/.ecdata:/root -p 127.0.0.1:5901:5901 -p 127.0.0.1:11080:1080 -p 127.0.0.1:18888:8888 hagb/docker-easyconnect:7.6.7


docker run --rm --device /dev/net/tun --cap-add NET_ADMIN -ti -e PASSWORD=175544 -e URLWIN=1 -e DISABLE_PKG_VERSION_XML=1 -v /Users/xiaohui/Docker/ec/.ecdata:/root -p 127.0.0.1:5901:5901 -p 127.0.0.1:11080:1080 -p 127.0.0.1:18888:8888 hagb/docker-easyconnect:7.6.3