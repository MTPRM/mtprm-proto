# API Admin

## Authentication

Use metadata key `jwt`.

### Authentication Example (Typescript)

```ts
import { Service } from '@buf/mtprm_mtprm-api-admin.connectrpc_es/mtprm/api/admin/resources/organizations__0__entities/v1/service_connect'
import { type PromiseClient, createPromiseClient } from '@connectrpc/connect'
import { createPromiseClient } from '@connectrpc/connect'
import { createGrpcWebTransport } from '@connectrpc/connect-web'

const METADATA_KEY = 'jwt'
const jwtValue = '123455...'

const transport = createGrpcWebTransport({
  baseUrl: 'https://api.mtprm.com',
  useBinaryFormat: true,
  interceptors: [
    (next) => async (request) => {
      request.header.set(METADATA_KEY, jwtValue)

      return next(request)
    },
  ],
})

const client = createPromiseClient(Service, transport)

client.create({
  org_id: {
    value: '1231',
  },
  // ...
})
```
