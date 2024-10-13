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
  const { fixedPaths, uniqueTags } = fixPaths(input.paths)

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
    paths: fixedPaths,
    tags: uniqueTags,
  }
}

function fixPaths(paths) {
  const fixedPaths = {}

  const tagHelper = new TagHelper()

  for (const curr of Object.entries(paths)) {
    // e.g.
    //
    // - pathName = `/mtprm.api.portfolio.resources.entities.v1.Service/List`
    // - pathValue = `{ ... }`
    const [pathName, pathValue] = curr

    if (pathName.startsWith('/mtprm.api.portfolio.resources')) {
      const parts = pathName.split('/')
      const [, fullName, serviceMethod] = parts

      const split = fullName.split('.')

      const resourceName = split[4]
      const serviceVersion = split[5]

      const tagName = tagHelper.add(resourceName, serviceVersion)

      fixedPaths[pathName] = Object.fromEntries(
        Object.entries(pathValue).map(([httpMethod, operationDefinition]) => [
          httpMethod,
          {
            ...operationDefinition,
            operationId: serviceMethod,
            tags: [tagName],
          },
        ])
      )
    }
  }

  return {
    fixedPaths,
    uniqueTags: tagHelper.sortedTags(),
  }
}

// Tracks seen tags and deals with proper sorting
class TagHelper {
  static GROUP_NAME_TO_ORDER = Object.fromEntries(
    ['Inquiries', 'Reports', 'Entities'].map((it, index) => [it, index])
  )

  fullNameToTag = new Map()

  add(resourceName, serviceVersion) {
    const groupName = this._groupName(resourceName)

    const fullName = `${groupName} (${serviceVersion})`

    this.fullNameToTag.set(fullName, { groupName, fullName })

    return fullName
  }

  sortedTags() {
    const uniqueValues = Array.from(this.fullNameToTag.values())
    uniqueValues.sort((a, b) => {
      const aGroup = TagHelper.GROUP_NAME_TO_ORDER[a.groupName] || -1
      const bGroup = TagHelper.GROUP_NAME_TO_ORDER[b.groupName] || -1

      const groupCompare = aGroup - bGroup

      if (groupCompare != 0) {
        return groupCompare
      }

      return a.fullName.localeCompare(b.fullName)
    })

    return uniqueValues.map((it) => ({ name: it.fullName }))
  }

  _groupName(resourceName) {
    switch (resourceName) {
      case 'entities':
        return 'Entities'
      case 'entities__0__reports__combined__xlsx':
      case 'entities__0__reports__summary':
        return 'Reports'
      case 'entity_inquiries':
        return 'Inquiries'
      default:
        return resourceName
    }
  }
}

main()
