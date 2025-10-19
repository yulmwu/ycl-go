# YCL implemented in GO

# Overview: What’s wrong with JSON and YAML?

JSON has long been the standard data format, and YAML was introduced as its more human-friendly superset.  
However, while YAML improves upon JSON’s verbosity, it introduces its own set of problems.

One major issue is its indentation-based syntax and high degree of syntactic freedom — this flexibility can often lead to confusion. While experienced developers might manage this well, YAML’s permissiveness increases the risk of human error.

Moreover, both JSON and YAML were designed and optimized primarily for data serialization and transmission — not necessarily for expressing configurations or compositions intuitively. This limitation becomes apparent when dealing with complex configuration scenarios.

For example, consider writing a Kubernetes YAML manifest where you want to combine the Pod’s `metadata.name` and `metadata.version` fields into an environment variable (using the Downward API).  
Would the following snippet work?

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-app
  version: v1
spec:
  containers:
    - name: my-container
      image: my-image:latest
      env:
        - name: APP_NAME
        # valueFrom:
        # fieldRef:
        #     fieldPath: metadata.name # Downward API
          value: "{{metadata.name}}-{{metadata.version}}" # or "${metadata.name}-${metadata.version}" ?, which one?
```

The answer is **“It doesn’t work.”**
YAML is merely a data representation format; it doesn’t support dynamic expressions, variable substitution, or template evaluation.
To achieve such behavior, you’d need a separate templating engine or application-level processing.

This limitation motivates a new direction: a format that’s not only a data representation language but also includes *programmable and composable* elements — a true superset of JSON and YAML.

---

# YCL: Yulmwu Configuration Language

We took inspiration from **HashiCorp Configuration Language (HCL)**, the DSL used by Terraform and other IaC (Infrastructure as Code) tools.
HCL is a JSON-compatible language, offering human readability and programmability.

However, HCL is fundamentally a domain-specific language — tightly coupled to Terraform’s ecosystem — which makes it less suitable as a general-purpose configuration language.

Enter **YCL (Yulmwu Configuration Language)** — a programmable, declarative configuration language that blends the best parts of JSON, YAML, and HCL.

```js
!import { foo, bar as _b } from 'path/to/other/file.ycl'
!schema 'path/to/schema/file.ycl'

my_value = 1234,
my_object {
    field1 = 'YCL',
    field2 = true,
    field3 = null,
    field4 = [ 'item1', 42, 'item3' ]
}

another_value   = `Hello $(my_object.field1) World!`
another_value_2 = @upper(my_object.field4[0]) + ' is the answer.'

// ... and when evaluated, this produces the following JSON:
```

```json
{
    "my_value": 1234,
    "my_object": {
        "field1": "YCL",
        "field2": true,
        "field3": null,
        "field4": [ "item1", 42, "item3" ]
    },
    "another_value": "Hello YCL World!",
    "another_value_2": "ITEM1 is the answer."
}
```

(TODO)

---

# Why was Go chosen for the first implementation?

(TODO)

---

(TODO)

* Human-friendly, readable syntax
* Programmable
* Composable
* Extensible
* Superset of JSON/YAML
* Schema validation and type checking

