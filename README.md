# Golang Payment Gateway

![continuous integration](https://github.com/jefflssantos/payment-gateway/actions/workflows/continuous_integration.yml/badge.svg)
[![License](http://poser.pugx.org/jefflssantos/payment-gateway/license)](https://packagist.org/packages/jefflssantos/payment-gateway)

## Golang Payment Gateway

Developed for study purposes only.

## Getting Started

### Clone the project
```bash
git clone git@github.com:jefflssantos/payment-gateway.git
```
### Start docker (docker compose must be intalled) and install composer dependencies
```bash
docker run --rm \
    -u "$(id -u):$(id -g)" \
    -v $(pwd):/var/www/html \
    -w /var/www/html \
    laravelsail/php81-composer:latest \
    composer install --ignore-platform-reqs
```

### Copy the ```.env.example```  to  ```.env```
```bash
cp .env.example .env
```