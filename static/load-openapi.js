async function main() {
  const rawOpenapi = await fetch(
    'gen-gateway-go/openapi/merged.swagger.json'
  )
  const rawOpenapiJson = await rawOpenapi.json()

  const openapi = convert(rawOpenapiJson)
  console.log(openapi)

  Redoc.init(
    openapi,
    {
      expandResponses: '200,400',
    },
    document.getElementById('redoc-container')
  )
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
        Object.entries(value).map(([key, value]) => [
          key,
          {
            ...value,
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
    paths,
    tags: [
      {name: 'Entities'},
      {name: 'Reports'},
      {name: 'Inquiries'},
    ],
  }
}

function getTag(value) {
  switch(value) {
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
