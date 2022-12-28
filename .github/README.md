# Uptimes

A simple server status dashboard heavily inspired by [gatus](https://github.com/TwiN/gatus).

## Features

-   Easily customizable dashboard page
-   Rules for what categorizes as uptime
-   High concurrency

## Setup

This application is published to [ghcr.io](ghcr.io/Testausserveri/uptimes). Therefore, the setup with docker is simple:

1. Create a directory to store the application configurations. This can for example be called `uptimes` or `status-page`

    ```sh
    export UPTIMES_DIRECTORY=uptimes # replace "uptimes" with the name you want
    mkdir $UPTIMES_DIRECTORY && cd $UPTIMES_DIRECTORY
    ```

2. Download the docker compose file

    ```sh
    curl -Lso docker-compose.yml raw.githubusercontent.com/Testausserveri/uptimes/main/docker-compose.yml
    ```

3. Install dependencies if you haven't already

    ```sh
    # ubuntu
    sudo apt install docker docker-compose
    # arch
    sudo pacman -S docker docker-compose
    ```

4. [Configure the program](#configuration)

5. Start the server

    ```sh
    docker-compose up -d
    ```

### <a id="configuration">Configuration</a>

-   Configuration is done in several files to enhance modularity and minimize complexity of larger scale infrastructures.
-   The configuration files for the program are often located in `configs/`-directory (which needs to be manually created), but if you want to use an alternative name feel free to go for it.
    Just remember to change it in your docker compose file (which was in this documentation named to `docker-compose.yml `) as well.

```toml
# ServePath configures the path where this configuration is served.
# In this example this is "/", which basically means that the page for this configuration
# is served at localhost:<port>/
#
# Limits:
# - Two or more configuration files can't be served at same path
ServePath = "/"

# This is often unneeded, but can be helpful for example in customization where
# you'd like to have different looks for different configurations.
#
# - Templates are located in $UPTIMES_DIRECTORY/public/
# - Assets are located in $UPTIMES_DIRECTORY/public/assets/
# - Assets are served at localhost:<port>/assets
TemplateName = "index.gohtml"

# The domains are configured in the Domains-list below.
# For each domain you will create a new list element which looks like following.
# For requirements

[[ Domains ]]
Name = "localhost"
Interval = "10s"
Url = "https://some-api.com"

[ Domains.Requirements ]
ContentType = "application/json"
Status = 200

[[ Domains ]]
Name = "localhost-1"
Interval = "10s"
Url = "http://someotherdoma.in"

[ Domains.Requirements ]
ContentType = "text/html"
Status = 200

```