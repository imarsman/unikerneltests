# unikerneltests

This repository is designed to be a collection of various tests of the Nanovms
unikernel in the cloud. At the moment the cloud being tested is the Google
Cloud. The AWS cloud looks to have similar overall supported instance API data
such as public and private ips, instance groups, etc. 

This project is a way for me to learn about how to use a unikernel instance in
the GCP, doing things like cloud logging, obtaining an instance's public IP,
private IP, and determining whether an instance is running in a managed instance
group.

One possible thing that could be done if for instance, a managed instance group
of a clustered NATS messaging service was running, each instance hopefully would
be able to ascertain the IPs of the other instances in the group, thus forming
an ad-hoc cluster. A client could use the same technique with sufficient
privileges, to obtain the IPs of the group's instances and thus join a NATS
cluster. I have not finished testing out setting up a cluster based on finding
IPs for an instance group at startup. We'll see.

I expect that much of what I am exploring here is already known. The goal is for
me to get to know it. I would like to find ways to avoid things like building a
Kubernetes cluster when Google already offers managed clusters.

## What works

- straightforward logging to GCP logs

## What is in progress

- core instance attributes like public/private IPs for GCE
  - use of interfaces works as does quickly testing if code running in cloud
- Abstraction of logging for supported clouds - at the moment is hard-coded for
  GCP, though that works fine.

## What does not work

- AWS support for instance information
- Any other cloud provider for instance information
- Setup of NATS cluster in context of instance group's IPs. This is the goal of
  the NATS part of the project. The idea is to avoid use of pre-known IPs when
  starting NATS cluster. This is new for me in terms of choosing a starting
  master instance. I am considering either just picking the instance with the
  earliest creation timestamp or using an algorithm such as the bully algorithm
  to choose a leader and use that to initiate the cluster. I am new to this.