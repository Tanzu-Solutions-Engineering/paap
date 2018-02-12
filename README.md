# paap - Platform as a Product CLI

Command line interface demo harness for Pivotal's Platform Engineer workshop. 
The CLI is designed to promote the following concepts:

1. Really good platform operators use golang to automate PCF operations. This isn't a
requirement but golang is the modern age BASH for sysadmins.

2. Most common sysadmin functions can be wrapped in a few "cf" calls. 

3. There are other sysadmin tools like "om" that can be used to automate via the Ops Manager API.

4. There are yet other sysadmin tools like "kubectl" that can be used to automate k8s operations.

5. The best way to explain how to do something is via a "demo". This tool is designed to be used primarily as a teach via demo tool.

## The GOAL of "paap" is to codeify and showcase common cloud operations activities and commands. 

## Getting Started

### Clone the github repo and run
```
go install
paap help
``` 

### Prerequisites

* [Golang](https://golang.org/doc/install)
* [Cloud Foundry CLI](https://github.com/cloudfoundry/cli#downloads)
* [PCF Dev](https://pivotal.io/pcf-dev) - for local testing. You can target any CF environment 
* Set environment variables denoted in "paap help". Reference for PCF Dev listed below:

```
export CF_USER=admin
export CF_PASS=admin
export CF_API=https://api.local.pcfdev.io
export CF_ORG=demo
export CF_SPACE=development
export CF_NETWORK_TOKEN=[YOUR_TOKN]
```

* Install CF CLI Plugins

```paap install-plugins```


### Running

1. Open a new console tab and source "development" environment variables

```source env_development```

2. Login to Cloud Foundry 

```paap login```

3. Create Line of Business(LOB) tenant

```paap create-lob```

4. Deploy application to "development"

```paap deploy-app```

6. Open a new console tab and start a "development" environment smoke test

```paap run-smoketest```

7. Open a new console tab and source "production" environment variables

```source env_production```

8. Deploy application to "production"

```paap deploy-app```

9. Upgrade middleware in "development" space only

```paap upgrade-middleware```

10. Notice that the "development" space has upgraded middleware and 
smoketest endpoint never returned an error. 
We upgraded the middlware without dropping requests.

11. Teardown LOB so you can run the demo in a clean environment. 

```paap teardown```

12. Teardown LOB and CF plugins in case you want to 



