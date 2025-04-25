# ip-country-checker

Reposistory for Avoxi coding challenge

## Design considerations

### Use of AI

First of all I would like to acknowledge that I used chatgpt to help me with this coding challenge, mainly for coding style/best practices for GO.

### Technical Choices

I downloaded the MMDB and am using the database locally instead of going with the web API route. For brevity and not having to deal with api keys are the main reasons for doing so.

I am using the unofficial official geoip2-golang MMDB reader. It is recommended in Maxmind docs to use a DB reader and although the go version isnt officially endorsed, it is listed in the Maxmind docs.

#### Extra notes

##### Extensions to project

Two things come to mind when thinking about extending features for this project. 

1. Adding caching to the ip look ups. MMDB lookups are fast and designed for scale but caching would solve for unnecessary repeated IP lookups.
2. Adding metrics with logging. Metrics and logging enhance observability, helping identify technical issues and giving the business insight to guide product direction.

##### Keeping MMDB up to date
There is a question posed in the coding requirements of "Documenting a plan for keeping the mapping data up to date.  Extra bonus points for implementing the solution."

For a production ready application, my recommendation would be to have the CI/CD rebuild the docker image we are deploying weekly with the docker build file downloading the latest MMDB.

For a locally maintained application, my recommendation would be to have a weekly script run and download the latest MMBD. I have included a sample powerscript file (omitting the needed API key)


# Project usage instructions

`docker compose up --build`

This service provides an HTTP endpoint to check whether a given IP address belongs to an allowed country.

## 📘 Endpoint

**POST** `/ip-authorization`

---

## 📄 Description

Checks whether a given IP address is from one of the allowed countries.

---

## 📥 Request

**Content-Type**: `application/json`

### Body Example

```json
{
  "ip": "8.8.8.8",
  "allowed_countries": ["US", "CA", "GB"]
}
```

### Fields

| Field             | Type       | Required | Description                                       |
|------------------|------------|----------|---------------------------------------------------|
| `ip`             | string     | ✅        | The IPv4 or IPv6 address to check.                |
| `allowed_countries` | string[] | ✅        | List of ISO country codes (e.g., "US", "CA").     |

---

## 📤 Response

**Content-Type**: `application/json`

### Body Example

```json
{
  "allowed": true
}
```

### Fields

| Field     | Type    | Description                                                 |
|-----------|---------|-------------------------------------------------------------|
| `allowed` | boolean | `true` if the IP’s country is in the allowed list; otherwise `false`. |

---

## ❌ Error Responses

| Code | Message           | Description                                |
|------|-------------------|--------------------------------------------|
| 400  | `Invalid request` | Malformed JSON or missing required fields. |
| 500  | `Internal error`  | Internal server or GeoIP lookup error.     |