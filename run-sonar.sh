#!/bin/sh
docker run --platform linux/x86_64 --rm --network host \
	-e SONAR_HOST_URL="http://localhost:9000" \
	-e SONAR_LOGIN="sqp_36fb102a27ef53f9788eb6ef63dad08c323c1770" \
	-v "${PWD}:/usr/src" sonarsource/sonar-scanner-cli