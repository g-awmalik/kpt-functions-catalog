

// Code generated by "mdtogo"; DO NOT EDIT.
package generated

var ListSettersShort = `Lists information about [setters] like setter name, values and count.

Refer to the [create-setters] function documentation for information about creating new setters or [apply-setters] function documentation for information about updating the field values parameterized by setters.`
var ListSettersLong = `
## Usage

` + "`" + `list-setters` + "`" + ` function is expected to be executed imperatively like

  $ kpt fn eval -i list-setters:unstable

` + "`" + `list-setters` + "`" + ` function performs the following steps:

1. Searches for setter comments in input list of resources.
1. Lists discovered setters and related information.
`
var ListSettersExamples = `
### Listing setters in a package

Let's start with the input resource in a package

  # resources.yaml
  apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: my-nginx
  spec:
    replicas: 4 # kpt-set: ${nginx-replicas}
    selector:
      matchLabels:
        app: nginx
    template:
      metadata:
        labels:
          app: nginx
      spec:
        containers:
        - name: nginx
          image: "nginx:1.16.1" # kpt-set: nginx:${tag}
          ports:
          - protocol: TCP
            containerPort: 80

Invoke the function:

  $ kpt fn eval --image gcr.io/kpt-fn/list-setters:unstable

Output looks like the following:

    Results:
      [INFO] Name: nginx-replicas, Value: 4, Type: int, Count: 1
      [INFO] Name: tag, Value: 1.16.1, Type: str, Count: 1
`
