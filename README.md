This project is a textbook example of Hexagonal Architecture (a.k.a. Ports & Adapters) for a simple arithmetic gRPC service backed by MySQL. The key principle: the core business logic has zero knowledge of frameworks, databases, or transport protocols — it only depends on interfaces (ports).

The 4 Ports (Interfaces) — internal/ports/
Each port is an interface that defines a contract. Nothing in internal/ports/ imports any concrete implementation.

Port	File	Purpose	Direction
ArithmeticPort	core.go	Math operations: Addition, Substraction, Multiplication, Division	Inward (driven by app)
APIPort	app.go	Orchestration layer: GetAddition, GetSubstraction, GetMultiplication, GetDivision	Inward (exposed to framework-left)
GRPCPort	framework_left.go	gRPC transport: Run() + 4 RPC handlers using OperationParameters/Answer protos	Outward (driving side)
DbPort	framework_right.go	Persistence: CloseDbConnection(), AddToHistory(answer, operation)	Outward (driven side)
The 4 Adapters (Implementations) — internal/adapters/
1. arithmetic.Adapter → satisfies ArithmeticPort
internal/adapters/core/arithmetic/arithmetic.go

Pure math — no IO, no errors (except division by zero)
Addition(a,b) → a+b
Substraction(a,b) → a-b
Multiplication(a,b) → a*b
Division(a,b) → a/b or error if b==0

2. api.Adapter → satisfies APIPort
internal/adapters/app/api.go

The orchestrator/Use Case layer. Holds references to ArithmeticPort (core logic) and DbPort (persistence).
Each Get* method: calls ArithmeticPort to compute → calls DbPort.AddToHistory() to persist → returns the answer.
This is where business flow is coordinated, not the math itself.

3. grpc.Adapter → satisfies GRPCPort and ArithmeticServiceServer (gRPC interface)
internal/adapters/framework/left/grpc/server.go + rpc.go

server.go: Defines Adapter struct with api APIPort field. Run() starts a gRPC server on port 9000.
rpc.go: Each method validates input (rejects zero params) → delegates to APIPort → returns a protobuf Answer.
Also implicitly satisfies the generated ArithmeticServiceServer gRPC interface (the 4 methods match exactly).

4. db.Adapter → satisfies DbPort
internal/adapters/framework/right/db/db.go

Opens a MySQL connection via database/sql + go-sql-driver/mysql.
AddToHistory() inserts rows into an arith_history table using Masterminds/squirrel as a query builder.
CloseDbConnection() closes the DB handle.
Interface Satisfaction Map
ArithmeticPort  ←── arithmetic.Adapter       (core math)
APIPort         ←── api.Adapter              (orchestration, depends on ArithmeticPort + DbPort)
GRPCPort        ←── grpc.Adapter            (gRPC transport, depends on APIPort)
DbPort          ←── db.Adapter              (MySQL persistence)
Additionally, grpc.Adapter satisfies the protobuf-generated ArithmeticServiceServer interface.

Wiring in cmd/main.go
var core ports.ArithmeticPort = arithmetic.NewAdapter()
var dbAdapter ports.DbPort   = db.NewAdapter("mysql", dsn)
var api   ports.APIPort      = api.NewAdapter(core, dbAdapter)
var grpc  *grpc.Adapter       = grpc.NewAdapter(api)
grpc.Run()

Data flow for a request:
gRPC Client
  → grpc.Adapter.GetAddition()          [framework-left: transport]
    → api.Adapter.GetAddition()          [app: orchestration]
      → arithmetic.Adapter.Addition()   [core: business logic]
      → db.Adapter.AddToHistory()       [framework-right: persistence]
  ← Answer protobuf
Dependency Rule
The dependency direction always points inward:

framework/left → app → core ← framework/right
(grpc)           (api)  (arithmetic)  (db)
core imports nothing from the project (pure Go).
app imports ports only (depends on abstractions).
framework/* import ports + their respective framework libraries.
cmd/main.go is the only place concrete types are coupled together.
This is the essence of hexagonal architecture: everything depends on ports (interfaces), never on concrete adapters.
