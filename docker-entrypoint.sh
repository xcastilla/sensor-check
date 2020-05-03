# https://medium.com/@bouffard.jul/use-case-add-an-init-script-to-the-docker-official-mongo-image-be58cb2dff25
for f in /docker-entrypoint-initdb.d/*; do
	case "$f" in
		*.sh) echo "$0: running $f"; . "$f" ;;
		*.js) echo "$0: running $f"; "${mongo[@]}" "$MONGO_INITDB_DATABASE" "$f"; echo ;;
		*)    echo "$0: ignoring $f" ;;
	esac
	echo
done