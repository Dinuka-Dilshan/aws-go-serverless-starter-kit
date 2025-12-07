## 🚀 Go Lambda + AWS CDK Starter Kit: Seamless Serverless Deployment

This template provides a **minimal** structure for rapidly deploying high-performance **Go-based AWS Lambda functions** using the **TypeScript AWS CDK (Infrastructure as Code)**.

### Core Components & Layout

| Component                 | Path              | Description                                                                                         |
| :------------------------ | :---------------- | :-------------------------------------------------------------------------------------------------- |
| **Go Source & Artifacts** | `app/`            | Contains **`main.go`** and **`go.mod`**. The compiled `bootstrap` binary is placed in `app/build/`. |
| **CDK Application**       | `infrastructure/` | Contains the TypeScript code (`*.ts`) that defines all AWS resources (Lambda, Function URL, etc.).  |

---

### Prerequisites Checklist ✅

Ensure these tools are installed to begin building your serverless application:

- **Go** (version 1.20+)
- **Node.js** (and **npm/pnpm**)
- **AWS CLI** (with valid credentials configured)
- **AWS CDK v2**

### Single-Command Workflow ⚡

The entire process building the optimized Go binary and deploying the cloud infrastructure is consolidated into one command:

```bash
npm run deploy
```

This ensures your infrastructure code and application code are always deployed in the correct, sequential order.
