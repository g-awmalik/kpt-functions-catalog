

// Code generated by "mdtogo"; DO NOT EDIT.
package generated

var SetAnnotationShort = `Add a list of annotations to all resources.`
var SetAnnotationLong = `
Configured using a ConfigMap with key-value pairs in ` + "`" + `data` + "`" + ` field in ` + "`" + `ConfigMap` + "`" + `
resource.

For example: To add an annotation ` + "`" + `color: orange` + "`" + ` to all resources:

  apiVersion: v1
  kind: ConfigMap
  metadata:
    name: my-config
  data:
    color: orange

To add 2 annotations ` + "`" + `color: orange` + "`" + ` and ` + "`" + `fruit: apple` + "`" + ` to all resources:

  apiVersion: v1
  kind: ConfigMap
  metadata:
    name: my-config
  data:
    color: orange
    fruit: apple

You can use key ` + "`" + `fieldSpecs` + "`" + ` to specify the resource selector you want to use.
By default, the function will not only add or update the annotations in
` + "`" + `metadata/annotations` + "`" + ` but also a bunch of different places where have
references to the annotations. These field specs are defined in
https://github.com/kubernetes-sigs/kustomize/blob/master/api/konfig/builtinpluginconsts/commonannotations.go#L6

You need to use a custom resource to specify additional information.

For example: To add an annotation ` + "`" + `color: orange` + "`" + ` to path ` + "`" + `data.selector` + "`" + ` in
` + "`" + `MyOwnKind` + "`" + ` resource:

  apiVersion: fn.kpt.dev/v1alpha1
  kind: SetAnnotationConfig
  metadata:
  name: my-config
  annotations:
  color: orange
  fieldSpecs:
  - path: data/selector
    kind: MyOwnKind
    create: true

To support your own CRDs you will need to add more items to fieldSpecs list.
Your own specs will be used with the default ones.

Field spec has following fields:

- group: Select the resources by API version group. Will select all groups if
  omitted.
- version: Select the resources by API version. Will select all versions if
  omitted.
- kind: Select the resources by resource kind. Will select all kinds if omitted.
- path: Specify the path to the field that the value will be updated. This field
  is required.
- create: If it's set to true, the field specified will be created if it doesn't
  exist. Otherwise, the function will only update the existing field.

For more information about fieldSpecs, please see
https://kubectl.docs.kubernetes.io/guides/extending_kustomize/builtins/#arguments-3
`
var SetAnnotationExamples = `
https://github.com/GoogleContainerTools/kpt-functions-catalog/tree/master/examples/set-annotation/
`
