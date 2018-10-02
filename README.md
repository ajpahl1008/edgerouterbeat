# Edgerouterbeat

The purpose of this project is to be able to take your statistics from your Ubiquiti Edgerouter X SFP and ship the data to the Elastic Stack.  Below is an example of the dashboards we'll be able to create when you install this beat on your Edgerouter.

Preview of Upcoming Dashboards (https://github.com/ajpahl1008/edgerouterbeat/issues/1)
![alt_text](https://github.com/ajpahl1008/edgerouterbeat/blob/master/images/Dashboard.png)


## Why Edgerouterbeat?

With zero documented remote APIs for the Edgerouter, using a lightweight framework provided by Elastic's libbeat protocol made creating a beat and gathering metrics easy.  The on-disk footprint is about 30MB and with a modest sample rate of 10-15 seconds, has very little impact to the routers CPU uzilization.


## Getting Started with Edgerouterbeat

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/<your_id>/edgerouterbeat`

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Building for EdgeRouter X SFP
```
env GOOS=linux GOARCH=mipsle go build 
```
