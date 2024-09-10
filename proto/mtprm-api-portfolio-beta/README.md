# API Portfolio (beta)

## Authentication

Use metadata key `jwt`.

### Authentication Example (Typescript)

```ts
import { Service } from '@buf/mtprm_mtprm-api-portfolio-beta.connectrpc_es/mtprm/api/portfolio/beta/resources/entities/v1/service_connect'
import { type PromiseClient, createPromiseClient } from '@connectrpc/connect'
import { createPromiseClient } from '@connectrpc/connect'
import { createGrpcTransport } from '@connectrpc/connect-node'

const METADATA_KEY = 'jwt'
const jwtValue = '123455...'

const transport = createGrpcTransport({
  baseUrl: 'https://api-portfolio.whitehawk-mtprm.com',
  httpVersion: '2',
  interceptors: [
    (next) => async (request) => {
      request.header.set(METADATA_KEY, jwtValue)

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
