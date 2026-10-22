# x9proto 优化方案

| # | 优先级 | 问题 | 位置 | 建议 | 工作量 |
|---|---|---|---|---|---|
| 1 | 高 | 金额类型不统一：`order.proto:114` 商品单价用 `double price`，`:137` 属性价格用 `int32 price`，payment/refund 用 `int64`（分） | `proto/order.proto` | 统一用 `int64 xxx_fen`；在新字段号上新增字段，旧字段标记 `[deprecated=true]`，避免破坏兼容 | M |
| 2 | 中 | 拼写错误：`FianceService`、`FianceKeyRequest`、`FianceKeyRsp`、`FianceRequest`；merchant.proto 里的服务名还和财务服务的含义混在一起 | `proto/merchant.proto:9`、`finance.proto` | 新增 `MerchantService`，旧名称保留一个版本后删除 | S |
| 3 | 中 | `go_package = "./order"` 用的是相对路径，换个目录生成会得到不同的导入路径 | 全部 .proto | 改成 `github.com/open4go/x9proto/pb/order;order` | S |
| 4 | 中 | 没有 lint 和破坏性变更检查，也没有 CI | 仓库 | 引入 `buf`：`buf lint`、`buf breaking --against .git#branch=main`，再加 GitHub Actions | S |
| 5 | 中 | govulncheck：`google.golang.org/grpc v1.70.0`（GO-2026-6348、GO-2026-6061）、x/net v0.32.0 | `go.mod` | 升级 grpc 到 v1.83.1+、x/net 到 v0.60.0 | S |
| 6 | 低 | Go 1.22.3 太旧，和其他仓库（1.24 或 1.25）不一致 | `go.mod` | 统一到组织基线版本 | S |
| 7 | 低 | `hello.proto` 是示例 | `proto/hello.proto` | 移到 examples/ 或删除 | S |
| 8 | 低 | `cc.sh` 没有固定 protoc 和插件版本 | `cc.sh` | 用 `buf generate` 加 `buf.gen.yaml` 固定插件版本 | S |
