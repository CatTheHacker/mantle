# Mantle
![loc](https://sloc.xyz/github/nektro/mantle)
[![license](https://img.shields.io/github/license/nektro/mantle.svg)](https://github.com/nektro/mantle/blob/master/LICENSE)
[![discord](https://img.shields.io/discord/551971034593755159.svg)](https://discord.gg/P6Y4zQC)
[![paypal](https://img.shields.io/badge/donate-paypal-009cdf)](https://paypal.me/nektro)
[![circleci](https://circleci.com/gh/nektro/mantle.svg?style=svg)](https://circleci.com/gh/nektro/mantle)
[![release](https://img.shields.io/github/v/release/nektro/mantle)](https://github.com/nektro/mantle/releases/latest)
[![goreportcard](https://goreportcard.com/badge/github.com/nektro/mantle)](https://goreportcard.com/report/github.com/nektro/mantle)
[![codefactor](https://www.codefactor.io/repository/github/nektro/mantle/badge)](https://www.codefactor.io/repository/github/nektro/mantle)

Easy and effective communication is the foundation of any successful team or community. That's where Mantle comes in, providing you the messaging platform that puts you in charge of both the conversation and the data.

## Getting Started
These instructions will help you get the project up and running and are required before moving on.

### Creating External Auth Credentials
In order to allow users to log in to Mantle, you will need to create an app on your Identity Provider(s) of choice. See the [nektro/go.oauth2](https://github.com/nektro/go.oauth2#readme) docs for more detailed info on this process on where to go and what data you'll need.

Here you can also fill out a picture and description that will be displayed during the authorization of users on your chosen Identity Provider. When prompted for the "Redirect URI" during the app setup process, the URL to use will be `http://mantle/callback`, replacing `mantle` with any origins you wish Mantle to be usable from, such as `example.com` or `localhost:800`.

Once you have finished the app creation process you should now have a Client ID and Client Secret. These are passed into Mantle through flags as well.

| Name | Type | Default | Description |
|------|------|---------|-------------|
| `--auth-{IDP-ID}-id` | `string` | none. | Client ID. |
| `--auth-{IDP-ID}-secret` | `string` | none. | Client Secret. |

The Identity Provider IDs can be found from the table in the [nektro/go.oauth2](https://github.com/nektro/go.oauth2#readme) documentation.



## Deployment
Pre-compiled binaries can be obtained from https://github.com/nektro/mantle/releases/latest.

## Built With
- https://github.com/gorilla/mux
- https://github.com/gorilla/sessions
- https://github.com/gorilla/websocket
- https://github.com/nektro/go-util
- https://github.com/nektro/go.dbstorage
- https://github.com/nektro/go.etc
- https://github.com/nektro/go.oauth2
- https://github.com/oklog/ulid
- https://github.com/spf13/pflag
- https://github.com/valyala/fastjson

## Contributing
[![issues](https://img.shields.io/github/issues/nektro/mantle.svg)](https://github.com/nektro/mantle/issues)

We listen to issues all the time right here on GitHub. Labels are extensively to show the progress through the fixing process. Question issues are okay but make sure to close the issue when it has been answered! Off-topic and '+1' comments will be deleted. Please use post/comment reactions for this purpose.

## Contact
- hello@nektro.net
- Meghan#2032 on discordapp.com
- https://twitter.com/nektro

## License
Apache 2.0
