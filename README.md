# gen-struct

### Mục đích
- Dùng để tạo struct đơn giản trong golang từ json

### Nhược điểm
- Chỉ hỗ trợ tạo struct đơn giản, chưa hỗ trợ tạo struct lòng nhau

### Ví dụ

- Ta có chuỗi json sau:

```json
    {
        "address": "string",
        "id": 0,
        "idStr": "string",
        "isTransferPoint": true,
        "min_customer": 0,
        "name": "string",
        "pointId": 0,
        "realTime": "string",
        "surcharge": 0,
        "surcharge_type": 0,
        "time": 0,
        "transsPointId": "string",
        "unfixed_point": 0
    }

```

- Kết quả sau khi chạy tool sẽ cho ta một struct như sau

```go
    type Object struct{
        Name    string  `json:"name"`
        PointId int64   `json:"pointId"`
        RealTime        string  `json:"realTime"`
        UnfixedPoint    int64   `json:"unfixed_point"`
        Time    int64   `json:"time"`
        Address string  `json:"address"`
        Id      int64   `json:"id"`
        IdStr   string  `json:"idStr"`
        IsTransferPoint bool    `json:"isTransferPoint"`
        MinCustomer     int64   `json:"min_customer"`
        Surcharge       int64   `json:"surcharge"`
        SurchargeType   int64   `json:"surcharge_type"`
        TranssPointId   string  `json:"transsPointId"`
    }


```