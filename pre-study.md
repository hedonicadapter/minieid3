### 1. **From access keys to managed identity**

This feature transitions us from using key vault access keys to using Azure managed identities for authentication to Redis. It improves security and simplifies credential management by letting Azure handle authentication automatically.

---

### 2. **Background**

Access keys require manual rotation and management, leading to operational overhead and security vulnerabilities. With managed identities, Azure handles credential issuance and rotation automatically, reducing the risk of compromised keys and removing the burden of manual management.

Go-redis' support for managed identity is a work in progress at the time of writing. For this to be implemented, the minimum requirement for us is to upgrade go-redis to v9, and/or potentially wait for an experimental go-redis feature ([StreamingCredentialsProvider](https://github.com/redis/go-redis/pull/3320)) to go live. I've compared the relevant go-redis features below.

#### CredentialsProvider -- since go-redis ~v9.0.0-beta.1
Credentials are set once at the time of initialization, without context and proper erroring.

#### CredentialsProviderContext -- since go-redis v9.5.2
Supports contexts and erroring, and credentials are determined at the time of each operation.

#### StreamingCredentialsProvider -- experimental and unreleased feature
Supports contexts and erroring, and credentials are dynamically updated during the connection lifecycle.

---

### 3. **Goal**

- remove need to manually rotate and manage access keys

---

### 4. **Prerequisites**

#### Infrastructure
Azure Cache for Redis with system or user assigned managed identity enabled and a role assignment of Redis Cache Contributor at the minimum scope of resource group.

#### go and go-redis
The rest of the prerequisites depend on which of the three solutions we want.

- CredentialsProvider
  - go [1.17](https://github.com/reactive-go/redis/blob/cfd0933930162b9b9836de80306a11b1e9950382/go.mod#L3)
  - go-redis 9.0.0-beta.1

- CredentialsProviderContext
  - go [1.17](https://github.com/redis/go-redis/blob/53c26684c137e21b019c40062744ef8d015f115a/go.mod#L3)
  - go-redis 9.5.2

- StreamingCredentialsProvider
  - go [1.23 and 1.24 are tested, but go.mod is set at 1.18. They are planning to bump to 1.24 in an upcoming release](https://github.com/redis/go-redis/blob/45e5ee96a1e093095dc6d3dd7e0e29ed7f29d4d5/README.md#supported-versions).
  - go-redis 9.5+

---

### 5. **Risks / Considerations**

CredentialsProvider has clear risks in silent erroring. From there on out it's unclear what implications dynamic credentials have. I can imagine a connection becoming invalid during an operation, but consequences would depend on how Azure handles connection invalidation and how go-redis responds to that.

---

### 6. **Action Plan outline**

1. Enable managed identity on relevant Azure Cache for Redis resource(s)
  - [Add system assigned identity to existing Azure Cache for Redis](https://learn.microsoft.com/en-us/azure/azure-cache-for-redis/cache-managed-identity#add-system-assigned-identity-to-an-existing-cache)
  - [Add user assigned identity to existing Azure Cache for Redis](https://learn.microsoft.com/en-us/azure/azure-cache-for-redis/cache-managed-identity#add-a-user-assigned-identity-to-an-existing-cache)
2. Assign the identity the role of "Redis Cache Contributor" at the minimum scope of resource group
3. Get hostname and port from Azure Cache for Redis instance(s) for go application to use
  Example for local development with docker-compose:
  ```docker-compose
  services:
    server:
      image: golang-app-that-connects-to-redis # needs to have Azure CLI
      working_dir: /app
      volumes:
        - ./server:/app
        - ~/.azure:/root/.azure # mount developer's own Azure config for authentication
      environment:
        - REDIS_DATABASE_URL=rediss://my-azure-cache-for-redis-instance.redis.cache.windows.net:6380/0
  ```
4. Set up authentication logic:
  - [For CredentialsProviderContext](https://github.com/hedonicadapter/minieid3/blob/e9b17fb426d9ac65a93899358a7a7a74d3e74e31/server/config/redis.go#L74)
  - [For StreamingCredentialsProvider](https://github.com/redis/go-redis-entraid?tab=readme-ov-file#minimal-example)
