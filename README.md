# unikerneltests

Various tests of the Nanovms unikernel in the cloud. At the moment the cloud
being tested is the Google Cloud. This project is a way for me to learn about
how to use a unikernel instance in the GCP, doing things like logging, obtaining
an instance's public IP, private IP, and determining whether an instance is
running in a managed instance group.

One possible thing that could be done if for instance, a managed instance group
of a clustered NATS messaging service was running, each instance hopefully would
be able to ascertain the IPs of the other instances in the group, thus forming
an ad-hoc cluster. A client could use the same technique with sufficient
privileges, to obtain the IPs of the group's instances and thus join a NATS
cluster.

I expect that much of what I am exploring here is already known. The goal is for
me to get to know it. I would like to find ways to avoid things like building a
Kubernetes cluster when Google already offers managed clusters.