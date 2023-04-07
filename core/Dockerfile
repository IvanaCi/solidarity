
FROM php:fpm-alpine3.16

RUN apk update

# for laravel lumen run smoothly
RUN apk add php-fpm \
php-pdo \
php-mbstring \
php-openssl

# for our code run smoothly
RUN apk --no-cache add \
php-json \
php-dom \
curl \
php-curl

# for swagger run smoothly
RUN apk --no-cache add \
php-tokenizer

# for composer & our project depency run smoothly
RUN apk --no-cache add \
php-phar \
php-xml \
php-xmlwriter

# if need composer to update plugin / vendor used
RUN php -r "copy('http://getcomposer.org/installer', 'composer-setup.php');" && \
php composer-setup.php --install-dir=/usr/bin --filename=composer && \
php -r "unlink('composer-setup.php');"

# copy all of the file in folder to /core
COPY . /app
WORKDIR /app

RUN composer update

# run the php server service
# move this command to -> docker-compose.yml
# CMD php -S 0.0.0.0:00 public/index.php