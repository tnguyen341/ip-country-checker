# ip-country-checker

Reposistory for Avoxi coding challenge

## Design considerations

### Use of AI

First of all I would like to acknowledge that I used chatgpt to help me with this coding challenge, mainly for coding style/best practices for GO.

### Technical Choices

I downloaded the MMDB and am using the database locally instead of going with the web API route. For brevity and not having to deal with api keys are the main reasons for doing so.

I am using the unofficial official geoip2-golang MMDB reader. It is recommended in Maxmind docs to use a DB reader and although the go version isnt officially endorsed, it is listed in the Maxmind docs.

#### Extra notes

There is a question posed in the coding requirements of "Documenting a plan for keeping the mapping data up to date.  Extra bonus points for implementing the solution."

For a production ready application, my recommendation would be to have the CI/CD rebuild the docker image we are deploying weekly with the docker build file downloading the latest MMDB.

For a locally maintained application, my recommendation would be to have a weekly script run and download the latest MMBD. I have included a sample powerscript file (omitting the needed API key)
