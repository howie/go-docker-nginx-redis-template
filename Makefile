
run:
	# Go1.9
	dev_appserver.py  app.yaml
	# Go1.11
	#source ./dev_env.sh && go run main.go
deploy:
	gcloud config set project ${ptoject-id}
	gcloud app deploy
	gcloud app deploy cron.yaml
log:
	gcloud config set project ${ptoject-id}
	gcloud app logs tail -s default


