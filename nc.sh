docker run -d \
    -v /media/user/7ae97150-eb9a-4dc9-8421-fc8b7f819d6c/nextcloud:/var/www/html \
    -v /media/user/7ae97150-eb9a-4dc9-8421-fc8b7f819d6c/apps:/var/www/html/custom_apps \
    -v /media/user/7ae97150-eb9a-4dc9-8421-fc8b7f819d6c/config:/var/www/html/config \
    -v /media/user/7ae97150-eb9a-4dc9-8421-fc8b7f819d6c/data:/var/www/html/data \
    nextcloud

