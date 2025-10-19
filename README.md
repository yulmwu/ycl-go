# YCL implemented in GO

# 개요: JSON과 YAML, 무슨 문제가 있었죠?

표준적인 데이터 포맷으로 JSON, 그리고 JSON의 불편함을 해소하기 위한 슈퍼셋 포맷인 YAML이 널리 사용되고 있습니다.
하지만 JSON의 불편함을 개선한 YAML 또한 몇가지 문제를 가지고 있습니다.

특히 그 중 들여쓰기 기반의 문법과 문법적 자유도가 높지만, 이는 종종 혼란을 야기할 수 있습니다. 물론 이는 개발자의 숙련도에 따라 다르겠지만, 휴먼 에러의 가능성을 높입니다.

그리고 두 포맷은 데이터 전송(직렬화)에 있어 최적화가 되어있고 그에 맞게 설계가 되어있지만, 어떠한 구성 또는 설정을 표현하는데 있어서는 다소 불편한 점이 있습니다.

예를 들어 쿠버네티스 YAML을 작성할 때, `metadata.name` 필드와 `metadata.version` 필드를 합쳐 환경 변수로 주입한다고 해봅시다. (Downward API)
아래와 같은 코드가 동작할까요?

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

정답은 "동작하지 않는다" 입니다. YAML은 단순히 데이터를 표현하는 포맷이기 때문에, 위와 같은 동적 참조나 템플릿 기능을 제공하지 않습니다.
이를 처리하기 위해선 별도의 템플릿 엔진이나 애플리케이션 레벨에서 처리해야 합니다.

그리고 우리는 컨셉을 다르게 하여, 단순히 데이터를 표현하는 포맷이 아닌 프로그래밍적 요소를 포함하는 JSON, YAML의 슈퍼셋 포맷을 원했습니다.

# YCL: Yulmwu Configuration Language

우리는 HashiCorp의 HCL(HashiCorp Configuration Language)에서 영감을 받았습니다. 
HCL은 IaC(Infrastructure as Code) 도구인 Terraform에서 사용되는 DSL(Domain Specific Language)로, JSON과 호환되는 언어입니다.

하지만 HCL은 근본적으로 DSL, 즉 Terraform 생태계예 tightly coupled 되어있기 때문에 범용적으로 사용하기에는 제한이 있습니다.

이제 YCL(Yulmwu Configuration Language)을 소개합니다. YCL은 JSON 및 YAML, 그리고 HCL의 장점을 결합한 프로그래밍적 선언적 구성 Configuration Language 입니다.

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

// ... and if you evaluate this, you can use it exactly like the JSON result below.
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

# 왜 첫번째 구현에 Go 언어를 선택했나요?

(TODO)

---

(TODO)

- Human friendly, readable syntax
- Programmable
- Composable
- Extensible
- Superset of JSON/YAML
- Schema validation, type checking


