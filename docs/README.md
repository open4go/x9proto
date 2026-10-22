# x9proto — gRPC 协议定义与生成代码

> 仓库：https://github.com/open4go/x9proto ｜ 模块：`github.com/open4go/x9proto` ｜ Go 1.22.3 ｜ 最近提交：2026-10-20（最新 tag v0.0.30）

## 项目简介
存放各服务之间 gRPC 接口的 `.proto` 文件，以及用 `protoc-gen-go v1.36.5`、`protoc-gen-go-grpc v1.5.1`（protoc v5.29.3）生成的 Go 代码。

## 目录结构与服务
```
x9proto/
├── proto/  cst / finance / hello / member / merchant / order / payment / wxpay .proto
├── pb/     对应的 *.pb.go 和 *_grpc.pb.go
├── cc.sh   遍历 proto/ 调用 protoc 生成到 pb/
└── push.sh version.sh
```
| 文件 | 服务 | RPC |
|---|---|---|
| order.proto | OrderService | FetchOrderDetail、FetchOrderPaymentInfo、FetchOrderIdentityInfo、UpdateOrderStatus |
| member.proto | MembershipService | RegisterMember、FetchMemberByID、FetchMemberByOpenID、FetchMemberIdentityInfo |
| payment.proto | PayService | PreparePayment、RefundApply |
| finance.proto | FinanceService | FetchFinanceKey |
| merchant.proto | **FianceService**（拼写有误） | FetchMerchantDetail、FetchMerchantFinance |
| wxpay.proto | WxPayService | FetchOrderPayInfo |
| hello.proto | Greeter | SayHello |

## 使用示例（x8rpc/call/order.go）
```go
clientCreator := func(conn *grpc.ClientConn) order.OrderServiceClient { return order.NewOrderServiceClient(conn) }
resp, err := x8rpc.CallGrpcService(ctx, "order", clientCreator, serviceMethod)
```

## 生成代码
```bash
sh cc.sh   # 需要 protoc、protoc-gen-go、protoc-gen-go-grpc
```

## 构建
`go build`、`go vet` 通过。没有 CI。

## 与其他仓库的关系
被 `x8rpc`（v0.0.27）以及 r2day 下的 gRPC 服务端（order/member/pay/finance 等）依赖。
