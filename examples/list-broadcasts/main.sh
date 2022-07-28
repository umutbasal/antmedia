if [ -f .env ]; then
	export $(echo $(cat .env | sed 's/#.*//g' | xargs) | envsubst)
fi

go run antmedia/cmd/cli broadcast_rest_service getBroadcastList --hostname ${ANTMEDIA_HOST} --base ${ANTMEDIA_BASE} --offset 0 --size 1
00
