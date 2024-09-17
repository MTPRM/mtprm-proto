# API Portfolio (beta)

## Authentication

Use metadata key `api-key`.

### Authentication Example (Typescript)

```ts
import { Service } from '@buf/mtprm_mtprm-api-portfolio-beta.connectrpc_es/mtprm/api/portfolio/beta/resources/entities/v1/service_connect'
import { type PromiseClient, createPromiseClient } from '@connectrpc/connect'
import { createPromiseClient } from '@connectrpc/connect'
import { createGrpcTransport } from '@connectrpc/connect-node'

const METADATA_KEY = 'api-key'
const apiKeyValue = 'mak_...'

const transport = createGrpcTransport({
  baseUrl: 'https://api-portfolio.whitehawk-mtprm.com',
  httpVersion: '2',
  interceptors: [
    (next) => async (request) => {
      request.header.set(METADATA_KEY, apiKeyValue)

      return next(request)
    },
  ],
})

const client = createPromiseClient(Service, transport)

client.create({
  entity: {
    name: '1231',
    // other fields...
  },
})
```
