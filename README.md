# kubemirror

## Scope

Copy entire namespace with all objects as is, even same secret values.
Copy objects by extracting specifications (YAML's) and corresponding images.
Currently no data is being copied.


## Objects & Treatment

- pod - ignore
- replicaset - ignore
- deployment - copy
- daemonset - copy
- statefulset - copy
- job - copy
- cronjob - copy
- endpoint - ignore
- service - copy
- ingress - copy
- configmap - copy
- secret - copy
- persistentVolumeClaim - copy
- persistentVolume - copy ? (definition only, no data)
