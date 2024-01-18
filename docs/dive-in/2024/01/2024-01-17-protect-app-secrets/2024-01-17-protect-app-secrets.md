# Protect app secrets

- Author: Cedric Gautier
- Date : 2024-01-17

## Summary

Protect application secrets in our environnements by safely fetching secrets in a secure vault.

## Context

We are currently building a product in which we use secrets to be able to connect to differennt services. It could be to connect to a database, an email service or a payment service like Stripe.

All these different services provide tokens to connect to them.

## Solution

**The right place to store secrets.**
To store the secrets in the right place, we need to find a wayy to store our them securely in vault that allows potential identity-based secrets and encryption management system.

For this we have a few choices :

- HashiCorp Vault is a tool that allows you to safely manage secrets. By secrets, we mean sensitive information like digital certificates, database credentials, passwords, and API encryption keys.

- AWS Secrets Manager is a secrets management service that helps you protect access to your applications, services, and IT resources. This service enables you to easily rotate, manage, and retrieve database credentials, API keys, and other secrets throughout their lifecycle.

We'll be using HashiCorp Vault as the setup is easier and the documentation is great. Very easy to read and understand.

### Hashicorp Vault

On our account, I have made an application on the Hashicorp Vault called Airneis.

To be able to login to Hashicorp services, we are going to create an An HCP Service Principal and the associated Client_ID and Client_Secret will be used for non-human access of HCP APIs from machines, applications, or system services.

#### How to authenticate:

The HCP API requires a valid Access Token. By authenticating to HCP with a user or Service Principal, you can retrieve a short-lived Access Token to call the HCP API.

Will have to use the equivalent of this request in Golang :

```sh
HCP_API_TOKEN=$(curl --location 'https://auth.hashicorp.com/oauth/token' \
--header 'content-type: application/json' \
--data '{
"audience": "https://api.hashicorp.cloud",
"grant_type": "client_credentials",
"client_id": "'$HCP_CLIENT_ID'",
"client_secret": "'$HCP_CLIENT_SECRET'"
}' | jq -r .access_token)
```

We'll have `$HCP_CLIENT_ID` & `HCP_CLIENT_SECRET` stored in the the local environnement of one of our instances.

#### How to fetch secrets in our vault:

##### From the Airneis app:

To fetch application secrets, we need to make a function in the Hashicorp Vault client that we'll be making which we'll be able to make a HTTP request that will GET on this endpoint: `https://api.cloud.hashicorp.com/secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/apps/{app_name}/secrets/{secret_name}`

Reference : https://developer.hashicorp.com/hcp/api-docs/vault-secrets#GetAppSecret

In one of our configuration files, we'll be placing : `location.organization_id`, `location.project_id` , `app_name` and all secret names.

It should look something like this :

```yaml
env:
  organization_id: blablabla
  project_id: blablabla
  app_name: blablabla
  secrets:
    STRIPE_ID: stripe_id
    STRIPE_TOKEN: stripe_token
    SENGDRID_ID: sendgrid_id
    SENDGRID_TOKEN: sendgrid_token
```

##### From your laptop

If needed to get an application secret from our laptop you can use the CLI Tool furnished by Hashicorp.

To install the Hashicorp's CLI Tool, follow these steps with the credentials that have been shared to each member of our team.
![CLI Tool](https://file%252B.vscode-resource.vscode-cdn.net/Users/cedric/DevPerso/sdv/airneis/docs/dive-in/2024/01/2024-01-17-protect-app-secrets/Screenshot%25202024-01-18%2520at%252018.51.59.png?version%253D1705600337373)
