# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/runtime/v1/primitive.proto](#atomix_runtime_v1_primitive-proto)
    - [CreatePrimitiveRequest](#atomix-runtime-v1-CreatePrimitiveRequest)
    - [CreatePrimitiveResponse](#atomix-runtime-v1-CreatePrimitiveResponse)
    - [DeletePrimitiveRequest](#atomix-runtime-v1-DeletePrimitiveRequest)
    - [DeletePrimitiveResponse](#atomix-runtime-v1-DeletePrimitiveResponse)
    - [GetPrimitiveRequest](#atomix-runtime-v1-GetPrimitiveRequest)
    - [GetPrimitiveResponse](#atomix-runtime-v1-GetPrimitiveResponse)
    - [Primitive](#atomix-runtime-v1-Primitive)
    - [PrimitiveId](#atomix-runtime-v1-PrimitiveId)
    - [PrimitiveMeta](#atomix-runtime-v1-PrimitiveMeta)
    - [PrimitiveMeta.LabelsEntry](#atomix-runtime-v1-PrimitiveMeta-LabelsEntry)
    - [RequestHeaders](#atomix-runtime-v1-RequestHeaders)
    - [ResponseHeaders](#atomix-runtime-v1-ResponseHeaders)
    - [UpdatePrimitiveRequest](#atomix-runtime-v1-UpdatePrimitiveRequest)
    - [UpdatePrimitiveResponse](#atomix-runtime-v1-UpdatePrimitiveResponse)
  
    - [PrimitiveManager](#atomix-runtime-v1-PrimitiveManager)
  
- [Scalar Value Types](#scalar-value-types)



<a name="atomix_runtime_v1_primitive-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/runtime/v1/primitive.proto



<a name="atomix-runtime-v1-CreatePrimitiveRequest"></a>

### CreatePrimitiveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primitive | [Primitive](#atomix-runtime-v1-Primitive) |  |  |






<a name="atomix-runtime-v1-CreatePrimitiveResponse"></a>

### CreatePrimitiveResponse







<a name="atomix-runtime-v1-DeletePrimitiveRequest"></a>

### DeletePrimitiveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primitive_id | [PrimitiveId](#atomix-runtime-v1-PrimitiveId) |  |  |






<a name="atomix-runtime-v1-DeletePrimitiveResponse"></a>

### DeletePrimitiveResponse







<a name="atomix-runtime-v1-GetPrimitiveRequest"></a>

### GetPrimitiveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primitive_id | [PrimitiveId](#atomix-runtime-v1-PrimitiveId) |  |  |






<a name="atomix-runtime-v1-GetPrimitiveResponse"></a>

### GetPrimitiveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primitive | [Primitive](#atomix-runtime-v1-Primitive) |  |  |






<a name="atomix-runtime-v1-Primitive"></a>

### Primitive



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meta | [PrimitiveMeta](#atomix-runtime-v1-PrimitiveMeta) |  |  |






<a name="atomix-runtime-v1-PrimitiveId"></a>

### PrimitiveId



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="atomix-runtime-v1-PrimitiveMeta"></a>

### PrimitiveMeta



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primitive_id | [PrimitiveId](#atomix-runtime-v1-PrimitiveId) |  |  |
| labels | [PrimitiveMeta.LabelsEntry](#atomix-runtime-v1-PrimitiveMeta-LabelsEntry) | repeated |  |






<a name="atomix-runtime-v1-PrimitiveMeta-LabelsEntry"></a>

### PrimitiveMeta.LabelsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="atomix-runtime-v1-RequestHeaders"></a>

### RequestHeaders



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primitive_id | [PrimitiveId](#atomix-runtime-v1-PrimitiveId) |  |  |






<a name="atomix-runtime-v1-ResponseHeaders"></a>

### ResponseHeaders







<a name="atomix-runtime-v1-UpdatePrimitiveRequest"></a>

### UpdatePrimitiveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primitive | [Primitive](#atomix-runtime-v1-Primitive) |  |  |






<a name="atomix-runtime-v1-UpdatePrimitiveResponse"></a>

### UpdatePrimitiveResponse






 

 

 


<a name="atomix-runtime-v1-PrimitiveManager"></a>

### PrimitiveManager


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetPrimitive | [GetPrimitiveRequest](#atomix-runtime-v1-GetPrimitiveRequest) | [GetPrimitiveResponse](#atomix-runtime-v1-GetPrimitiveResponse) |  |
| CreatePrimitive | [CreatePrimitiveRequest](#atomix-runtime-v1-CreatePrimitiveRequest) | [CreatePrimitiveResponse](#atomix-runtime-v1-CreatePrimitiveResponse) |  |
| UpdatePrimitive | [UpdatePrimitiveRequest](#atomix-runtime-v1-UpdatePrimitiveRequest) | [UpdatePrimitiveResponse](#atomix-runtime-v1-UpdatePrimitiveResponse) |  |
| DeletePrimitive | [DeletePrimitiveRequest](#atomix-runtime-v1-DeletePrimitiveRequest) | [DeletePrimitiveResponse](#atomix-runtime-v1-DeletePrimitiveResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |
