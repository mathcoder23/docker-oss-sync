docker run -e TASK_CRON="0 2 * * *" \
-e OSS_ENDPOINT=oss-cn-beijing.aliyuncs.com \
-e OSS_KEY_ID=xxx \
-e OSS_KEY_SECRET=xxx \
-e OSS_BUCKET_NAME=xxx \
-e OSS_PREFIX=backups \
-v xxx:/sync \
docker-oss-sync:1.0

