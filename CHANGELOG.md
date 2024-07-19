# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

## [15.0.0](https://github.com/mtprm/mtprm-proto/compare/14.3.0...15.0.0) (2024-07-18)

### ⚠ BREAKING CHANGES

- **api-admin:** remove service `organizations__0__reports__combined__template_xlsx.v1`

### Features

- **api-admin:** add service `organizations__0__entities__0__reports__combined__xlsx.v1` ([40118c2](https://github.com/mtprm/mtprm-proto/commit/40118c26b255f3f5d11f8c0a6f948c0d8457cc68))
- **api-admin:** remove service `organizations__0__reports__combined__template_xlsx.v1` ([b537444](https://github.com/mtprm/mtprm-proto/commit/b537444fb91d7e6635f31c88808da7f00b8264b3))

## [14.3.0](https://github.com/mtprm/mtprm-proto/compare/14.2.0...14.3.0) (2024-07-18)

### Features

- **api-admin:** add service `organizations__0__reports__combined__template_download_info.v1` ([cc61c9b](https://github.com/mtprm/mtprm-proto/commit/cc61c9b1611ca2cdc76dec0c865ec8f1bc66aa8c))
- **api-admin:** add service `organizations__0__reports__combined__template_upload_info.v1` ([db66140](https://github.com/mtprm/mtprm-proto/commit/db661405a6f6bba175fd13aaf706095725c70c69))

## [14.2.0](https://github.com/mtprm/mtprm-proto/compare/14.1.0...14.2.0) (2024-07-16)

### Features

- **api-admin:** add service `organizations__0__reports__combined__template_xlsx.v1` ([793504a](https://github.com/mtprm/mtprm-proto/commit/793504a2728db2a6ddbd220663f0ff2f57f09e7c))
- **proto:** add messages for `indirect_file` ([0617a75](https://github.com/mtprm/mtprm-proto/commit/0617a75511d402b4723961679b39a6474545bc6f))

## [14.1.0](https://github.com/MTPRM/mtprm-proto/compare/14.0.0...14.1.0) (2024-07-11)

### Features

- **api-admin:** add `organizations__0__reports__combined__zip` service ([66698e0](https://github.com/MTPRM/mtprm-proto/commit/66698e09d7966d53f2f615464bc4d23301ee3380))
- remove unused `proto-old` (i.e. `mtprm-proto`) ([332998b](https://github.com/MTPRM/mtprm-proto/commit/332998bda9717a37d72b4b82706ec224ab01a259))

## [14.0.0](https://github.com/MTPRM/mtprm-proto/compare/13.0.0...14.0.0) (2024-06-24)

### ⚠ BREAKING CHANGES

- **api-audit:** remove extra/unused `risk_relationship` + `overall_grade_letter`

### Features

- **api-audit:** remove extra/unused `risk_relationship` + `overall_grade_letter` ([6fe7b07](https://github.com/MTPRM/mtprm-proto/commit/6fe7b079668d759f46e7d219fb343a5feaf8cd6e))

## [13.0.0](https://github.com/MTPRM/mtprm-proto/compare/12.3.0...13.0.0) (2024-06-24)

### ⚠ BREAKING CHANGES

- **api-audit:** mirror `risk_config` data model

### Features

- **api-audit:** mirror `risk_config` data model ([700bffd](https://github.com/MTPRM/mtprm-proto/commit/700bffd9b3ec670fbab7a64e85fab6f808796183))

## [12.3.0](https://github.com/MTPRM/mtprm-proto/compare/12.2.0...12.3.0) (2024-06-18)

### Features

- **api-audit:** prefer array of string ([bd35e7d](https://github.com/MTPRM/mtprm-proto/commit/bd35e7d5585fc1d1f42417329979df6ba697d552))

## [12.2.0](https://github.com/MTPRM/mtprm-proto/compare/12.1.0...12.2.0) (2024-06-18)

### Features

- **api-audit:** capture `risk_relationship` ([0607c60](https://github.com/MTPRM/mtprm-proto/commit/0607c6095f2140fe8965f5b860707add090a3823))

## [12.1.0](https://github.com/MTPRM/mtprm-proto/compare/12.0.0...12.1.0) (2024-06-13)

### Features

- **api-admin:** add `organizations__0__reports__issues__xlsx.v1` ([aa032f7](https://github.com/MTPRM/mtprm-proto/commit/aa032f72dd106c0a4658d6261358a680955bbd04))

## [12.0.0](https://github.com/MTPRM/mtprm-proto/compare/11.0.5...12.0.0) (2024-06-04)

### ⚠ BREAKING CHANGES

- **api-admin:** use imported `entity` message

### Features

- **api-admin:** add `organizations__0__entities__0__reports__issues__xlsx` resource ([82ae469](https://github.com/MTPRM/mtprm-proto/commit/82ae4693b5cfffc4e749f738c769dd248e49bdd5))
- **api-admin:** add `organizations__0__entities__0__reports__summary__docx` resource ([b6c3acc](https://github.com/MTPRM/mtprm-proto/commit/b6c3acc48d9bd809800aa654b8822bb1c4dc4fc9))
- **api-admin:** add `organizations__0__entities__0__reports__summary__pdf` resource ([0e6e838](https://github.com/MTPRM/mtprm-proto/commit/0e6e838a721475bcbed163c194749682e8df84c2))
- **api-admin:** add `organizations__0__entities__0__reports__summary` resource ([557dd55](https://github.com/MTPRM/mtprm-proto/commit/557dd558a62e456fef2236de647083b5cb7021ce))
- **api-admin:** use imported `entity` message ([7563cd9](https://github.com/MTPRM/mtprm-proto/commit/7563cd9b5fc28f6c23a69da2e3830db58c0912d2))
- **api-audit:** add `organizations__0__entities` resource ([f5b5f61](https://github.com/MTPRM/mtprm-proto/commit/f5b5f611c87d4a3bdc6c8b8d45b7a692947076f5))
- **core:** add `entity` message ([c2e85d7](https://github.com/MTPRM/mtprm-proto/commit/c2e85d75d42822fffbbef347a2ae3a6ef1b84b00))
- **core:** add `summary_report` message ([bed4a4e](https://github.com/MTPRM/mtprm-proto/commit/bed4a4ee93d4d9bfa4a47dda86c49dd1cd3396ee))
- **core:** add `time` message ([ec70ac2](https://github.com/MTPRM/mtprm-proto/commit/ec70ac2bab08f32fa1c3a8ceb63212ef1c3b11ec))

### [11.0.5](https://github.com/mtprm/mtprm-proto/compare/11.0.4...11.0.5) (2024-06-04)

### [11.0.4](https://github.com/mtprm/mtprm-proto/compare/11.0.3...11.0.4) (2024-06-04)

### [11.0.3](https://github.com/mtprm/mtprm-proto/compare/11.0.2...11.0.3) (2024-06-04)

### [11.0.2](https://github.com/mtprm/mtprm-proto/compare/11.0.1...11.0.2) (2024-06-04)

### [11.0.1](https://github.com/mtprm/mtprm-proto/compare/11.0.0...11.0.1) (2024-06-04)

## [11.0.0](https://github.com/MTPRM/mtprm-proto/compare/10.3.0...11.0.0) (2024-06-04)

### ⚠ BREAKING CHANGES

- move `proto-2` to `proto`
- move `proto` to `proto-old`

### Features

- **api-admin:** add `organizations__0__entities` resource ([09f35dd](https://github.com/MTPRM/mtprm-proto/commit/09f35dd763674df7a3de0c16ef31f9d2b7830766))
- **core:** add `uuid.v1` message ([726541a](https://github.com/MTPRM/mtprm-proto/commit/726541a94c078ebc13234faef517c5c85a480702))
- move `proto-2` to `proto` ([9af0001](https://github.com/MTPRM/mtprm-proto/commit/9af00018788a1693c7ea91fefdc885872ee00225))
- move `proto` to `proto-old` ([45a8565](https://github.com/MTPRM/mtprm-proto/commit/45a8565b9c2e3a22b7b7e34c146d2c986ac203d1))

### Bug Fixes

- **api-admin:** remove rating from entity data ([78ff829](https://github.com/MTPRM/mtprm-proto/commit/78ff8294b94484793fa5fa6ed70c6efee99bb2ba))

## [10.3.0](https://github.com/MTPRM/mtprm-proto/compare/10.2.0...10.3.0) (2024-06-04)

### Features

- **api-admin:** add rating to entity data ([2edf5ef](https://github.com/MTPRM/mtprm-proto/commit/2edf5ef227f5588d9fe5bdae156a1fe0d0169e37))

## [10.2.0](https://github.com/MTPRM/mtprm-proto/compare/10.1.0...10.2.0) (2024-05-16)

### Features

- **api-admin:** add `organizations__0__entities__0__reports__control_mappings__xslx.v1` resource ([f35638f](https://github.com/MTPRM/mtprm-proto/commit/f35638f43b882da3e3a11fc733929ddc17925b20))

## [10.1.0](https://github.com/MTPRM/mtprm-proto/compare/10.0.0...10.1.0) (2024-05-02)

### Features

- **api-admin:** add `organizations__0__entities__0__reports__security__docx/v1` resource ([0a912c5](https://github.com/MTPRM/mtprm-proto/commit/0a912c5ff4f3b21c3400e86089795eb6dfcf26b8))
- **api-admin:** add `organizations__0__entities__0__reports__security__pdf/v1` resource ([c4c2243](https://github.com/MTPRM/mtprm-proto/commit/c4c2243bba3ee2bdfd46bfe0516e65af7be751e9))
- **api-admin:** add `organizations__0__entities__0__reports__security/v1` resource ([9ec6780](https://github.com/MTPRM/mtprm-proto/commit/9ec6780ef5a6f49563ac3d2a1bfa0633f04593ee))
- **api-admin:** add `organizations__0__entities/v1` resource ([7fb0835](https://github.com/MTPRM/mtprm-proto/commit/7fb0835ba78030a4b1050c68e0478a1a37ddb3e1))
- **api-admin:** add `organizations__0__reports__security__docx/v1` resource ([3a0c14d](https://github.com/MTPRM/mtprm-proto/commit/3a0c14d59239aa0cc307398ab0e90e9e3a26031a))
- **api-admin:** add `organizations__0__reports__security__pdf/v1` resource ([4d6fd70](https://github.com/MTPRM/mtprm-proto/commit/4d6fd70d7c6e216453b78e3ded2fc0b71b74a146))
- **api-admin:** add `organizations__0__reports__security/v1` resource ([6b637b4](https://github.com/MTPRM/mtprm-proto/commit/6b637b4eac0dfe9dd961087c8f67137ca2c13d61))
- **api-audit:** add `organizations__0__entities__0__reports/v1` resource ([9d1cb87](https://github.com/MTPRM/mtprm-proto/commit/9d1cb8799c2d2d226614600b761a837578cf11bd))
- **api-audit:** add `organizations__0__entities/v1` resource ([bc84f84](https://github.com/MTPRM/mtprm-proto/commit/bc84f8430eab659d0a1a151a7213ac43a8186a0d))
- **api-audit:** add `organizations__0__transactions/v1` resource ([6871fd5](https://github.com/MTPRM/mtprm-proto/commit/6871fd5176e43bd8f5f9635618fc2dfcc1b93c05))
- **api-audit:** add `organizations__0__usage/v1` resource ([fb28ad2](https://github.com/MTPRM/mtprm-proto/commit/fb28ad25419690c1dd9be7e3511e443cd4279826))
- remove custom generation for `java` ([32c54f0](https://github.com/MTPRM/mtprm-proto/commit/32c54f04678e62d7ca8e3b91684a5cf9010844aa))
- remove custom generation for `web` ([95c2026](https://github.com/MTPRM/mtprm-proto/commit/95c202661737bc045a01b5ef4351084343981796))

## [10.0.0](https://github.com/mtprm/mtprm-proto/compare/9.0.3...10.0.0) (2024-04-09)

### ⚠ BREAKING CHANGES

- **api:** remove unused message
- **api:** remove unused message
- **api:** remove unused message
- **api:** remove unused message

### Features

- **api-admin:** prefer messages from `proto` directory ([da01442](https://github.com/mtprm/mtprm-proto/commit/da0144248eb132180ef9493efa85d51aa0e3442a))
- **api-admin:** prefer messages from `proto` directory ([d25fccb](https://github.com/mtprm/mtprm-proto/commit/d25fccb12f563a7b852ac5973aba305d50315dd7))
- **api-admin:** prefer messages from `proto` directory ([758fbdb](https://github.com/mtprm/mtprm-proto/commit/758fbdbcb40bc8ff82b1e7e7547ace178d89fde2))
- **api-admin:** prefer messages from `proto` directory ([44654ca](https://github.com/mtprm/mtprm-proto/commit/44654cae866540f0e8485454c576c27fe0fd5875))
- **api-admin:** prefer messages from `proto` directory ([e77bc86](https://github.com/mtprm/mtprm-proto/commit/e77bc86cf603a869283ff172f8f32f1d628f6190))
- **api-audit:** prefer messages from `proto` directory ([8bba3d8](https://github.com/mtprm/mtprm-proto/commit/8bba3d8c7153cbecc3e89e899820af88f8fa1cb8))
- **api-audit:** prefer messages from `proto` directory ([89cb950](https://github.com/mtprm/mtprm-proto/commit/89cb9507bfd156fa8f7a79ae7d165bbeb04d4bae))
- **api-audit:** prefer messages from `proto` directory ([3b8475b](https://github.com/mtprm/mtprm-proto/commit/3b8475bc721eac1fc5cfcfcc87c8bd77ea3bb8de))
- **api-audit:** prefer messages from `proto` directory ([2ef2d65](https://github.com/mtprm/mtprm-proto/commit/2ef2d65cf44e6fd9b8232f785af9dce8db44a96d))
- **api:** remove unused message ([f3cdbcb](https://github.com/mtprm/mtprm-proto/commit/f3cdbcbd57a82746952abf531d8173d40e9c9cb3))
- **api:** remove unused message ([84afdb5](https://github.com/mtprm/mtprm-proto/commit/84afdb590cb332a3fce7b0a1847a8253a48a78a2))
- **api:** remove unused message ([fe1ec92](https://github.com/mtprm/mtprm-proto/commit/fe1ec92c26ec9b68630a200e3f75b62056b51c42))
- **api:** remove unused message ([077761b](https://github.com/mtprm/mtprm-proto/commit/077761b47710b53fa38c0c636448f5a5090edbdf))
- **proto:** add `grade_letter` message ([0c0bfba](https://github.com/mtprm/mtprm-proto/commit/0c0bfbaa91a274285bd17bdc1ae6b2e56cc0a9e1))
- **proto:** add `security_report` message ([fc76cfb](https://github.com/mtprm/mtprm-proto/commit/fc76cfbf15a48aa585cc0d6a92b05815f40380ee))
- **proto:** add `time` message ([12333e7](https://github.com/mtprm/mtprm-proto/commit/12333e73bee1a64ab76a6cb073e07500c6d209e9))
- **proto:** add `uuid` message ([ec5471b](https://github.com/mtprm/mtprm-proto/commit/ec5471bf6241b4e30a2e66efc83c799ca2fd91b6))

### [9.0.3](https://github.com/mtprm/mtprm-proto/compare/9.0.2...9.0.3) (2024-04-03)

### [9.0.2](https://github.com/mtprm/mtprm-proto/compare/9.0.1...9.0.2) (2024-04-03)

### [9.0.1](https://github.com/mtprm/mtprm-proto/compare/9.0.0...9.0.1) (2024-04-03)

## [9.0.0](https://github.com/mtprm/mtprm-proto/compare/8.0.1...9.0.0) (2024-04-03)

### ⚠ BREAKING CHANGES

- **proto:** change module name

### Features

- **proto:** change module name ([489294a](https://github.com/mtprm/mtprm-proto/commit/489294a8c62acf90e89fe96f505b7a55f7f6b8a9))

### [8.0.1](https://github.com/mtprm/mtprm-proto/compare/8.0.0...8.0.1) (2024-04-03)

## [8.0.0](https://github.com/MTPRM/mtprm-proto/compare/7.0.0...8.0.0) (2024-04-02)

### ⚠ BREAKING CHANGES

- **api-admin:** avoid using `EntityData` in `CreateRequest`

### Features

- **api-admin:** add `requested_domain` to `EntityData` ([78cb3d0](https://github.com/MTPRM/mtprm-proto/commit/78cb3d0e2d65850361145810dc7b3d4a8a639b86))
- **api-admin:** prefer to return `stream` ([9b3aa22](https://github.com/MTPRM/mtprm-proto/commit/9b3aa227fc082ecbb7837f36f4e219d995d33e75))
- **api-admin:** prefer to return `stream` ([650c757](https://github.com/MTPRM/mtprm-proto/commit/650c757affc372760ed4b9a77a45a95761ad34de))
- **api-admin:** prefer to return `stream` ([7d40935](https://github.com/MTPRM/mtprm-proto/commit/7d40935fb09ddc0caf8b7f27b8880e16f6388e7a))
- **api-admin:** prefer to return `stream` ([e1f02cf](https://github.com/MTPRM/mtprm-proto/commit/e1f02cf5cbf62510747ef7b1bf95bf36d5c8ba6e))
- **dev:** add `eclipse-formatter.xml` ([dee99e0](https://github.com/MTPRM/mtprm-proto/commit/dee99e07df1b82bef8a58d78ed6986da7162b3ef))

### Bug Fixes

- **api-admin:** avoid using `EntityData` in `CreateRequest` ([1983b87](https://github.com/MTPRM/mtprm-proto/commit/1983b8796f115cba42b0216402c1ef915f76c405))

## [7.0.0](https://github.com/MTPRM/mtprm-proto/compare/6.0.0...7.0.0) (2024-04-02)

### ⚠ BREAKING CHANGES

- **api-audit:** prefer only `List` service, returns needed entity info
- **api-admin:** use correct package name
- **api-admin:** move resource to `api-admin`
- **api-admin:** move resource to `api-admin`
- **api-admin:** move resource to `api-admin`
- **api-admin:** move resource to `api-admin`
- **api-admin:** move resource to `api-admin`
- **api-admin:** move resource to `api-admin`
- **api-admin:** move resource to `api-admin`

### Features

- **api-admin:** move resource to `api-admin` ([13eb522](https://github.com/MTPRM/mtprm-proto/commit/13eb5225d5027c91d1bb894fdf6f588ef4d602f7))
- **api-admin:** move resource to `api-admin` ([ac0f107](https://github.com/MTPRM/mtprm-proto/commit/ac0f1073f9123dfe355b846ee4bdaf152e7cb0f8))
- **api-admin:** move resource to `api-admin` ([7469394](https://github.com/MTPRM/mtprm-proto/commit/74693941d240dec2fc3672ded697407fa1c00a1b))
- **api-admin:** move resource to `api-admin` ([ca7578c](https://github.com/MTPRM/mtprm-proto/commit/ca7578c95ac1ff2e840220592ea93e2a21095e61))
- **api-admin:** move resource to `api-admin` ([c084c8b](https://github.com/MTPRM/mtprm-proto/commit/c084c8b2a62ff4ef9835283ce2786dae876edb7e))
- **api-admin:** move resource to `api-admin` ([77196a4](https://github.com/MTPRM/mtprm-proto/commit/77196a47e5a64a2c45e4247aa2c985b267c8f828))
- **api-admin:** move resource to `api-admin` ([596ede2](https://github.com/MTPRM/mtprm-proto/commit/596ede2eaf29716e81bd0d7ad75f69d4f0431682))
- **api-audit:** prefer only `List` service, returns needed entity info ([3365547](https://github.com/MTPRM/mtprm-proto/commit/3365547749b31685719164fdb1db66ac6983f83e))

### Bug Fixes

- **api-admin:** use correct package name ([53ec777](https://github.com/MTPRM/mtprm-proto/commit/53ec7771ab90d9edd0fe43b6e71f6538d2a90f88))

## [6.0.0](https://github.com/MTPRM/mtprm-proto/compare/5.0.0...6.0.0) (2024-04-02)

### ⚠ BREAKING CHANGES

- **api-audit:** prefer to use imported message
- **api-audit:** prefer to use imported message

### Features

- **api-audit:** add `entities__0__reports` service ([ed8ec3b](https://github.com/MTPRM/mtprm-proto/commit/ed8ec3b32b1ecc5f3881721af9c80733e6517524))
- **api-audit:** add `transactions/v1` service ([5605cb8](https://github.com/MTPRM/mtprm-proto/commit/5605cb8232c9c119a8cf9f386e0ec883aebaab3b))
- **api-audit:** add `usage/v1` service ([0afc106](https://github.com/MTPRM/mtprm-proto/commit/0afc106c82ae874ff89a7a7f82e4b7e4d84bc01d))
- **api-audit:** prefer to use imported message ([9edbdea](https://github.com/MTPRM/mtprm-proto/commit/9edbdea1dc14bc237d55f74c4d174df35c8a85d9))
- **api-audit:** prefer to use imported message ([881d7d2](https://github.com/MTPRM/mtprm-proto/commit/881d7d2335234d2d27a3dc3af15f5296206ce616))
- **api:** add `security_report/v1` message ([689e62a](https://github.com/MTPRM/mtprm-proto/commit/689e62a70986080879d5a63840b26a2e21e3c2de))

## [5.0.0](https://github.com/MTPRM/mtprm-proto/compare/4.0.0...5.0.0) (2024-03-28)

### ⚠ BREAKING CHANGES

- **api-audit:** prefer to use imported `GradeLetter`

### Features

- **api-audit:** add `entities__0__reports__security__docx.v1` resource ([c0d866c](https://github.com/MTPRM/mtprm-proto/commit/c0d866cebb18ece35815079f2010a3b38647a358))
- **api-audit:** add `entities__0__reports__security__pdf.v1` resource ([1aff1c5](https://github.com/MTPRM/mtprm-proto/commit/1aff1c5b629495690e7b94ed993d63dfc7ef801d))
- **api-audit:** add `entities__0__reports__security.v1` resource ([82737c8](https://github.com/MTPRM/mtprm-proto/commit/82737c8e523f372dbce9da5b7c335249bb490ee7))
- **api-audit:** add `reports__security__docx.v1` resource ([792fa1d](https://github.com/MTPRM/mtprm-proto/commit/792fa1d5582fe3321c976366f662fc8f13611527))
- **api-audit:** add `reports__security__pdf.v1` resource ([40237c4](https://github.com/MTPRM/mtprm-proto/commit/40237c4af06d6b52e64591e8f53373f76e0a9872))
- **api-audit:** add `reports__security.v1` resource ([69ee3da](https://github.com/MTPRM/mtprm-proto/commit/69ee3dab37fbb16c3ee39b2238d848753d2a2906))
- **api-audit:** prefer to use imported `GradeLetter` ([a2aa1be](https://github.com/MTPRM/mtprm-proto/commit/a2aa1bee67405c0172e3ac98ec4e468047349c5b))
- **api:** add `grade_letter` message ([759cc65](https://github.com/MTPRM/mtprm-proto/commit/759cc65343fb213e210017c1e022af3eece25751))
- **api:** add `time` message ([030736e](https://github.com/MTPRM/mtprm-proto/commit/030736e74fd0fb324dcdb215b0d9269c3912dd3f))

### Bug Fixes

- **api-audit:** use correct import ([e223435](https://github.com/MTPRM/mtprm-proto/commit/e22343513a4851703b694e66f24029f8270c5fd7))

## [4.0.0](https://github.com/MTPRM/mtprm-proto/compare/3.0.0...4.0.0) (2024-03-28)

### ⚠ BREAKING CHANGES

- **app:** remove `app` project

### Features

- **api-audit:** add `entities/v1` resource ([2701f81](https://github.com/MTPRM/mtprm-proto/commit/2701f81dd010014e9423b80ea881b741cc60aab1))
- **api:** add `uuid` message ([1619585](https://github.com/MTPRM/mtprm-proto/commit/161958524d33ccf1f6978f1ee88068e9bc683afc))
- **app:** remove `app` project ([e6e9410](https://github.com/MTPRM/mtprm-proto/commit/e6e94102992c422a671f1059cf72cef1632fccfb))

## [3.0.0](https://github.com/MTPRM/mtprm-proto/compare/2.0.0...3.0.0) (2024-03-27)

### ⚠ BREAKING CHANGES

- **app-admin:** remove example `.proto`

### Features

- **app-admin:** remove example `.proto` ([ba3bae1](https://github.com/MTPRM/mtprm-proto/commit/ba3bae1a716ab9e7bc44ea25a5a0109efed06755))

## 2.0.0 (2024-03-27)

### ⚠ BREAKING CHANGES

- **java:** generate code without "com." package prefix

### Features

- add example `.proto` ([b88c140](https://github.com/MTPRM/mtprm-proto/commit/b88c140405ae880cde494ff09cf796d4823a1b2a))
- **app-admin:** add `entities/v1` resource ([c86a843](https://github.com/MTPRM/mtprm-proto/commit/c86a843c51162f43a37299ace2beac7674a4e877))
- **app:** add `time` message ([bc1de7b](https://github.com/MTPRM/mtprm-proto/commit/bc1de7b500cf38a1da662d6aa17adb5ccc73462e))
- **app:** add `uuid` message ([82cb00f](https://github.com/MTPRM/mtprm-proto/commit/82cb00f3f7cd05ac4917e84714eb6e965068a3b1))
- **ci:** update permission for execution ([62549e0](https://github.com/MTPRM/mtprm-proto/commit/62549e02abeb79068bf6994b92ffb79e7916e19a))
- **java:** add project files ([2b91f94](https://github.com/MTPRM/mtprm-proto/commit/2b91f948dab54aba3995b190f55cb715aa773420))
- **java:** generate code without "com." package prefix ([8fb0f23](https://github.com/MTPRM/mtprm-proto/commit/8fb0f23e5684b19cee0674e32fd00497eefd72a8))
- **web:** add `web` project ([c34308f](https://github.com/MTPRM/mtprm-proto/commit/c34308f9be9fa5759798af252b85341e13c1bb98))

### Bug Fixes

- **web:** add `connectrpc` dependencies ([6c2b266](https://github.com/MTPRM/mtprm-proto/commit/6c2b266399d7199f954c3f71f037b784d3dfecff))
- **web:** use correct name ([ebee853](https://github.com/MTPRM/mtprm-proto/commit/ebee8531b55b393189364f95360dc7d9f7ed4cf5))
- **web:** use correct name ([3400f83](https://github.com/MTPRM/mtprm-proto/commit/3400f831877db3b395652eca2fa52449c4e74a40))
