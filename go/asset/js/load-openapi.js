async function main() {
  const rawOpenapi = await fetch('openapi/merged.swagger.json')
  const rawOpenapiJson = await rawOpenapi.json()

  const openapi = convert(rawOpenapiJson)

  document.getElementById('api-reference').text = JSON.stringify(openapi)

  const scalar = document.createElement('script')
  scalar.type = 'text/javascript'
  scalar.src = 'https://cdn.jsdelivr.net/npm/@scalar/api-reference'
  document.body.appendChild(scalar)
}

function convert(input) {
  const paths = {}

  for (const curr of Object.entries(input.paths)) {
    const [key, value] = curr

    if (key.startsWith('/mtprm.api.portfolio.resources')) {
      const parts = key.split('/')
      const [, fullName, method] = parts

      const split = fullName.split('.')
      const tag = getTag(split[4])

      paths[key] = Object.fromEntries(
        Object.entries(value).map(([key2, value2]) => [
          key2,
          {
            ...value2,
            operationId: method,
            tags: [tag],
          },
        ])
      )
    }
  }

  return {
    ...input,
    info: {
      title: 'MTPRM Portfolio API',
      termsOfService: 'https://www.whitehawk.com/terms-conditions',
      contact: {
        email: 'info@whitehawk.com',
      },
      version: '',
    },
    securityDefinitions: {
      ApiKeyHeader: {
        type: 'apiKey',
        name: 'Grpc-Metadata-api-key',
        in: 'header',
      },
    },
    security: [{ ApiKeyHeader: [] }],
    paths,
    tags: [{ name: 'Inquiries' }, { name: 'Entities' }, { name: 'Reports' }],
  }
}

function getTag(value) {
  switch (value) {
    case 'entities':
      return 'Entities'
    case 'entities__0__reports__combined__xlsx':
    case 'entities__0__reports__summary':
      return 'Reports'
    case 'entity_inquiries':
      return 'Inquiries'
  }
}

main()
