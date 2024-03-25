# Controller-Manager

The controller runs on the master node and permanently:

1. Watches Status
2. Remediate Situation

The controller is a process continuously monitores the state of various components within the system and works towards bringing the whole system to the desired state.

- Node Controller
  - Monitores the status of the nodes
  - Takes the status of the nodes every 5 seconds
  - Takes action to keep the app running
  - It does this through the api-server

- Replication Controller
  - Monitors the status of replicasets
  - Ensures that the desired number of pods is given at all points in time within the set.

There are many more controllers like:

- Deployment Controller
- Namespace Controller
- Job Controller
- Endpoint Controller
- PV-Protection Controller
- PV-Binder Controller
- CronJob
- Stateful-Set
