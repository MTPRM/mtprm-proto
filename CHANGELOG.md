# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

### [26.1.1](https://github.com/mtprm/mtprm-proto/compare/26.1.0...26.1.1) (2024-10-13)

## [26.1.0](https://github.com/mtprm/mtprm-proto/compare/26.0.2...26.1.0) (2024-10-13)


### Features

* **go:** generate openapi schema with fully qualified names for models ([dac314d](https://github.com/mtprm/mtprm-proto/commit/dac314d749f00fdc777c6d1ce2ff3bddaa5f789f))

### [26.0.2](https://github.com/MTPRM/mtprm-proto/compare/26.0.1...26.0.2) (2024-10-11)


### Bug Fixes

* **api-admin:** rename `Entity` field ([0f0d168](https://github.com/MTPRM/mtprm-proto/commit/0f0d1685e85a2691a27db72c57977e389a3e8020))

### [26.0.1](https://github.com/MTPRM/mtprm-proto/compare/26.0.0...26.0.1) (2024-10-11)


### Bug Fixes

* **core:** remove double import ([6ed5933](https://github.com/MTPRM/mtprm-proto/commit/6ed59339d09545ecaccd910e22071f064ccf5195))

## [26.0.0](https://github.com/MTPRM/mtprm-proto/compare/25.0.0...26.0.0) (2024-10-11)


### ⚠ BREAKING CHANGES

* **admin:** prefer `Entity` for `UpdateRequest`

### Features

* **admin:** prefer `Entity` for `UpdateRequest` ([d6d53b2](https://github.com/MTPRM/mtprm-proto/commit/d6d53b2ce9d43b8d4f47a7d9969998be9823fba5))

## [25.0.0](https://github.com/MTPRM/mtprm-proto/compare/24.1.0...25.0.0) (2024-10-11)


### ⚠ BREAKING CHANGES

* **core:** add `inherent_risk_rating` to `EntityData`

### Features

* **admin:** prefer `EntityData` for `UpdateRequest` ([5cd70e5](https://github.com/MTPRM/mtprm-proto/commit/5cd70e55e2a03a93484c3b6f14f18fc61adc1f07))
* **core:** add `inherent_risk_rating` to `EntityData` ([1f07e49](https://github.com/MTPRM/mtprm-proto/commit/1f07e4927a3bc8a7df6793be28f52f9303af7b42))

## [24.1.0](https://github.com/MTPRM/mtprm-proto/compare/24.0.0...24.1.0) (2024-10-09)

### Features

- **admin:** add `organizations__0__entities__0__reports__combined__xlsx.v2` ([1aa675b](https://github.com/MTPRM/mtprm-proto/commit/1aa675b47654505b0029d68889719767fa7281f5))
- **admin:** add `organizations__0__entities__0__reports__summary.v2` ([cde21b6](https://github.com/MTPRM/mtprm-proto/commit/cde21b61fd8ee67da1722e5b9ee6036eddc1141d))

## [24.0.0](https://github.com/MTPRM/mtprm-proto/compare/23.4.0...24.0.0) (2024-10-09)

### ⚠ BREAKING CHANGES

- **admin:** prefer to use `PrefixedId`

### Bug Fixes

- **admin:** prefer to use `PrefixedId` ([71ebbae](https://github.com/MTPRM/mtprm-proto/commit/71ebbae6a4f5cbebf8e7cb39e3da2a1e24a65df5))

## [23.4.0](https://github.com/MTPRM/mtprm-proto/compare/23.3.5...23.4.0) (2024-10-08)

### Features

- **admin:** add `organizations__0__entities.v2` ([8ae8fee](https://github.com/MTPRM/mtprm-proto/commit/8ae8fee7326ed49052447333bbc4c99c3a5b2668))
- **admin:** add `organizations__0__entity_inquiries.v1` ([324419b](https://github.com/MTPRM/mtprm-proto/commit/324419b64e4fa666baaff27dd27df9167f9fb1c0))
- **core:** add `EntityInquiry` v2 message ([2425a5a](https://github.com/MTPRM/mtprm-proto/commit/2425a5a1230d7f076a7cd1dd5e1980a6a8169696))
- **portfolio:** add `entity_inquiries.v2` ([b1de917](https://github.com/MTPRM/mtprm-proto/commit/b1de917f07ea5a33a22532504448b3e52ab10131))

### [23.3.5](https://github.com/MTPRM/mtprm-proto/compare/23.3.4...23.3.5) (2024-10-08)

### [23.3.4](https://github.com/MTPRM/mtprm-proto/compare/23.3.3...23.3.4) (2024-10-08)

### Bug Fixes

- **proto:** use correct update param ([f0abd44](https://github.com/MTPRM/mtprm-proto/commit/f0abd44a67e7c5f6580e3ce83840aab49ea2b936))

### [23.3.3](https://github.com/MTPRM/mtprm-proto/compare/23.3.2...23.3.3) (2024-10-07)

### Bug Fixes

- **proto:** missing import ([984819a](https://github.com/MTPRM/mtprm-proto/commit/984819a0a2a218a47eb1d4b02021abfde349ade4))

### [23.3.2](https://github.com/MTPRM/mtprm-proto/compare/23.3.1...23.3.2) (2024-10-07)

### Bug Fixes

- **proto:** correct types for dss level and irr ([8a62b88](https://github.com/MTPRM/mtprm-proto/commit/8a62b8847abe66c40b708d4197fcfc8996e3bdfd))

### [23.3.1](https://github.com/MTPRM/mtprm-proto/compare/23.3.0...23.3.1) (2024-10-07)

### Bug Fixes

- **proto:** entity update message ([eda1f8b](https://github.com/MTPRM/mtprm-proto/commit/eda1f8ba6f16e2c2d6583853414a1db84bf6226b))

## [23.3.0](https://github.com/MTPRM/mtprm-proto/compare/23.2.0...23.3.0) (2024-10-07)

### Features

- **proto:** add entity update service ([90191aa](https://github.com/MTPRM/mtprm-proto/commit/90191aa474d21553b126c668b7a5e53305d545d0))

## [23.2.0](https://github.com/MTPRM/mtprm-proto/compare/23.1.2...23.2.0) (2024-10-04)

### Features

- **proto:** entity update fields ([206b05f](https://github.com/MTPRM/mtprm-proto/commit/206b05fb4aadf775ea11f15c3f025734aa8c2fc5))
- **proto:** entity update service ([eafe317](https://github.com/MTPRM/mtprm-proto/commit/eafe317c5482b3b2c79aa729071c143710e71fd4))

### [23.1.2](https://github.com/mtprm/mtprm-proto/compare/23.1.1...23.1.2) (2024-10-04)

### [23.1.1](https://github.com/mtprm/mtprm-proto/compare/23.1.0...23.1.1) (2024-10-03)

## [23.1.0](https://github.com/mtprm/mtprm-proto/compare/23.0.0...23.1.0) (2024-10-03)

### Features

- **proto:** mark required fields ([0f7c947](https://github.com/mtprm/mtprm-proto/commit/0f7c94743253cc3312ab994c6e06bc928c2663ff))
- **proto:** mark required fields ([cf1910f](https://github.com/mtprm/mtprm-proto/commit/cf1910fa6339611b4717bb0fb10905f45172c517))
- **proto:** mark required fields ([880149f](https://github.com/mtprm/mtprm-proto/commit/880149f1394865888fbeca1a5b345dd61960cc61))
- **proto:** mark required fields ([20d72e3](https://github.com/mtprm/mtprm-proto/commit/20d72e3cf2c371c0d6b9575a78bf04a01ce1f78b))
- **proto:** mark required fields ([da56571](https://github.com/mtprm/mtprm-proto/commit/da56571424d3e91919e913bb6b6b93ac8233506d))
- **proto:** mark required fields ([7fd3b96](https://github.com/mtprm/mtprm-proto/commit/7fd3b961ee9ce77792b83db57272908e041f59ea))
- **proto:** mark required fields ([4ea2139](https://github.com/mtprm/mtprm-proto/commit/4ea2139ba500fd7eb5c47bf7a9c9202f1980a125))
- **proto:** mark required fields ([7170802](https://github.com/mtprm/mtprm-proto/commit/7170802034be5404c2214844f2c681f2d5c8c01c))
- **proto:** mark required fields ([5ac5e9b](https://github.com/mtprm/mtprm-proto/commit/5ac5e9b6efeb29c13a1ff5f1aa4cf83a7bb1e602))
- **proto:** mark required fields ([0b10aac](https://github.com/mtprm/mtprm-proto/commit/0b10aaca2b8e1911a61b848ab9e1c90d8b8f5036))
- **proto:** mark required fields ([aa05e12](https://github.com/mtprm/mtprm-proto/commit/aa05e12c76c746342f39a26872c0adfdbb6732b9))
- **proto:** mark required fields ([fa22d8a](https://github.com/mtprm/mtprm-proto/commit/fa22d8a6769e5ac3e55a9765fefeaf825b547f61))
- **proto:** mark required fields ([5d2c5da](https://github.com/mtprm/mtprm-proto/commit/5d2c5da7aecce8730b822a9e60c2d5bcb5b493ba))

## [23.0.0](https://github.com/mtprm/mtprm-proto/compare/22.0.0...23.0.0) (2024-10-03)

### ⚠ BREAKING CHANGES

- **go:** remove `api-portfolio-beta` endpoints

### Features

- **go:** remove `api-portfolio-beta` endpoints ([f314cbe](https://github.com/mtprm/mtprm-proto/commit/f314cbeeaa9bc30b0951673e288f9199d360e712))

## [22.0.0](https://github.com/MTPRM/mtprm-proto/compare/21.0.0...22.0.0) (2024-10-02)

### ⚠ BREAKING CHANGES

- **core:** prefer `dss_level` as optional

### Features

- **core:** prefer `dss_level` as optional ([1ae1f42](https://github.com/MTPRM/mtprm-proto/commit/1ae1f420c3e096669d4d052cc9168b480e693357))

## [21.0.0](https://github.com/MTPRM/mtprm-proto/compare/20.0.2...21.0.0) (2024-10-02)

### ⚠ BREAKING CHANGES

- **core:** prefer `string` for duns number

### Features

- **core:** prefer `string` for duns number ([bbc669e](https://github.com/MTPRM/mtprm-proto/commit/bbc669eea1d7abfdd3c23595c2b1b0a32dea0f1d))

### [20.0.2](https://github.com/MTPRM/mtprm-proto/compare/20.0.1...20.0.2) (2024-10-02)

### Bug Fixes

- **go:** re-order tags ([ec6b085](https://github.com/MTPRM/mtprm-proto/commit/ec6b085ecf4fb3d8bb903f9c3d45500e8adb5e8f))

### [20.0.1](https://github.com/MTPRM/mtprm-proto/compare/20.0.0...20.0.1) (2024-10-02)

### Bug Fixes

- **api-portfolio:** update description ([8838677](https://github.com/MTPRM/mtprm-proto/commit/8838677c90558f9f111c36891a72dc06fa5d24d0))

## [20.0.0](https://github.com/MTPRM/mtprm-proto/compare/19.0.0...20.0.0) (2024-10-02)

### ⚠ BREAKING CHANGES

- **api-portfolio:** remove optional `grade` param

### Features

- **api-portfolio:** remove optional `grade` param ([5d44301](https://github.com/MTPRM/mtprm-proto/commit/5d4430137785961997388db51fea92e59631e099))
- **core:** add `InherentRiskRating` message ([11157c4](https://github.com/MTPRM/mtprm-proto/commit/11157c4f6d7d54ba1fd536cdd853e252a56dbb46))
- **core:** add `risk_relationship` and `inherent_risk_rating` fields ([290e222](https://github.com/MTPRM/mtprm-proto/commit/290e2220908db00ea1003faaf5e2ba8ef3f5abe5))

## [19.0.0](https://github.com/MTPRM/mtprm-proto/compare/18.1.0...19.0.0) (2024-10-02)

### ⚠ BREAKING CHANGES

- **core:** include `GradeLetter`
- **core:** prefer `double`
- **core:** prefer `double` to `string`

### Features

- **core:** include `GradeLetter` ([b5c5f8d](https://github.com/MTPRM/mtprm-proto/commit/b5c5f8dc28367bf0964adf322421a46df328c4e1))
- **core:** prefer `double` ([6297247](https://github.com/MTPRM/mtprm-proto/commit/6297247679778d5e17afb394a91a3d9418769733))
- **core:** prefer `double` to `string` ([ab8b838](https://github.com/MTPRM/mtprm-proto/commit/ab8b8387d49094770bc1c3ad5012f1db53330a28))

## [18.1.0](https://github.com/mtprm/mtprm-proto/compare/18.0.4...18.1.0) (2024-10-01)

### Features

- **go:** remove now impossible situation ([a92977e](https://github.com/mtprm/mtprm-proto/commit/a92977ed16fc4a54bb6ae7eaa97b4923ebdbe552))
- **go:** support the new api-porfolio endpoints (non-beta) ([a6245b3](https://github.com/mtprm/mtprm-proto/commit/a6245b360d1c0098e1c283fa5590f63c7191e617))

### [18.0.4](https://github.com/MTPRM/mtprm-proto/compare/18.0.3...18.0.4) (2024-10-01)

### [18.0.3](https://github.com/MTPRM/mtprm-proto/compare/18.0.2...18.0.3) (2024-10-01)

### [18.0.2](https://github.com/MTPRM/mtprm-proto/compare/18.0.1...18.0.2) (2024-10-01)

### [18.0.1](https://github.com/MTPRM/mtprm-proto/compare/18.0.0...18.0.1) (2024-10-01)

### Bug Fixes

- **go:** send `application/json` upon 401 error ([0ab14dc](https://github.com/MTPRM/mtprm-proto/commit/0ab14dc7a750af3c6d27d0146b104bc7f36c153f))

## [18.0.0](https://github.com/MTPRM/mtprm-proto/compare/17.0.0...18.0.0) (2024-10-01)

### ⚠ BREAKING CHANGES

- **api-portfolio:** remove `optional` keyword
- **go:** switch from `redoc` to `scalar` for openapi UI
- move `index.html` under go grpc gateway
- **go:** prefer lowercase module name

### Features

- add `index.html` ([1c5bff7](https://github.com/MTPRM/mtprm-proto/commit/1c5bff736612b25ce1b51b808d2370f911ddcf76))
- add `load-openapi.js` ([f0c63a1](https://github.com/MTPRM/mtprm-proto/commit/f0c63a11307afc1b68868a0c3709ca083be65b5e))
- **api-portfolio:** add comments ([7f324a9](https://github.com/MTPRM/mtprm-proto/commit/7f324a9489523e632c5ea22d04c93dc5e57dbe9b))
- **api-portfolio:** prefer to use `oneof` ([f93c6ea](https://github.com/MTPRM/mtprm-proto/commit/f93c6ea9a17a76d1e386f73aa5b8907aa16425b8))
- **go:** add helper to point to known file assets ([3d6e6e3](https://github.com/MTPRM/mtprm-proto/commit/3d6e6e3b39c17cedda9f946523ed0712a280460f))
- **go:** dynamically determine endpoints with streaming server response ([eb9183f](https://github.com/MTPRM/mtprm-proto/commit/eb9183fa66046325720e9f84f80bb8763fa550e9))
- **go:** move assets to one folder ([9803337](https://github.com/MTPRM/mtprm-proto/commit/980333707629387284d1217c600dbfbf65ca0e27))
- **go:** prefer lowercase module name ([4a2acbd](https://github.com/MTPRM/mtprm-proto/commit/4a2acbdb53839d348334efbf8a67421771c37a06))
- **go:** remove `console.log` ([bb22d9e](https://github.com/MTPRM/mtprm-proto/commit/bb22d9e954684d917256cab0498f54989bfabc8e))
- **go:** remove `console.log` ([35223ca](https://github.com/MTPRM/mtprm-proto/commit/35223cac79cc6f86d0e77f86187603b1169c546a))
- **go:** switch from `redoc` to `scalar` for openapi UI ([82656d2](https://github.com/MTPRM/mtprm-proto/commit/82656d210eaeafb9271ea339903a283bd3b2ba9d))
- **go:** use embedded files rather than rely on file system ([09a04f7](https://github.com/MTPRM/mtprm-proto/commit/09a04f7620216ecc9adbd457c071d5685fc1c26c))
- grpc proxied ([d469254](https://github.com/MTPRM/mtprm-proto/commit/d4692548716786dd620e576e6ba4822501658e60))
- move `index.html` under go grpc gateway ([0c84698](https://github.com/MTPRM/mtprm-proto/commit/0c84698ceba7b9e86cb7ef5bb679947fc26f97ef))
- render the swagger yaml ([eaec851](https://github.com/MTPRM/mtprm-proto/commit/eaec851b0c53d8ac3cc9b54373fdb4239c327485))
- use clearer tags ([d9e9d66](https://github.com/MTPRM/mtprm-proto/commit/d9e9d66bdf680f5869a924e74fe3c8384426c3df))

### Bug Fixes

- **api-portfolio:** remove `optional` keyword ([6ddd8fe](https://github.com/MTPRM/mtprm-proto/commit/6ddd8feffd542ee6a0841708f42564b47e12d436))
- **api-portfolio:** update comments ([6503187](https://github.com/MTPRM/mtprm-proto/commit/65031873e16022f884c5ecd9773ed37ed564f6ab))
- **api-portfolio:** update comments ([8776146](https://github.com/MTPRM/mtprm-proto/commit/8776146fc1e3d4a8e5f2ac99b6a96be01ec893a7))
- **api-portfolio:** update comments ([996c774](https://github.com/MTPRM/mtprm-proto/commit/996c774455629fbc656924b15595240b207982bc))
- **go:** avoid duplicate names ([2f66cff](https://github.com/MTPRM/mtprm-proto/commit/2f66cff191f2c4b81de7bc5b0cbd6155603a19ee))
- **go:** avoid sending `application/x-ndjson` on error ([95e2379](https://github.com/MTPRM/mtprm-proto/commit/95e23791edb802c06ba10428efb51125fcd4bfb9))
- **go:** send `application/x-ndjson` for streaming response ([4d946a7](https://github.com/MTPRM/mtprm-proto/commit/4d946a7fecb286e9f9d333226ee402285ab407c5))
- **go:** update path to file ([1352195](https://github.com/MTPRM/mtprm-proto/commit/1352195e0ed5652a5741583d5288222b7f4e79f9))
- **go:** update script path ([1963ffd](https://github.com/MTPRM/mtprm-proto/commit/1963ffdc28468661598dbbfd08d1276114b8763e))
- prefer to hardcode tags ([457838a](https://github.com/MTPRM/mtprm-proto/commit/457838a386173568b4b719fae0625d8c64e8696f))

## [17.0.0](https://github.com/MTPRM/mtprm-proto/compare/16.4.1...17.0.0) (2024-09-30)

### ⚠ BREAKING CHANGES

- **api-portfolio:** remove unnecessary `Create/Update/Delete` endpoints for now

### Features

- **api-portfolio:** remove unnecessary `Create/Update/Delete` endpoints for now ([1bdd91c](https://github.com/MTPRM/mtprm-proto/commit/1bdd91cd094067aaabd73a420dc6df6b9acb8886))

### [16.4.1](https://github.com/MTPRM/mtprm-proto/compare/16.4.0...16.4.1) (2024-09-30)

## [16.4.0](https://github.com/MTPRM/mtprm-proto/compare/16.3.1...16.4.0) (2024-09-30)

### Features

- **api-portfolio:** add `entities__0__reports__combined__xlsx` v1 resource ([62a13d7](https://github.com/MTPRM/mtprm-proto/commit/62a13d7560ecd76ee48b98f53d366491f8bc1f39))
- **api-portfolio:** add `entities__0__reports__summary` v1 resource ([ecc5cde](https://github.com/MTPRM/mtprm-proto/commit/ecc5cde171b855599d08affb4fa9f659c5807f99))
- **api-portfolio:** add `entities` v1 resource ([58a5e1d](https://github.com/MTPRM/mtprm-proto/commit/58a5e1de4e85a70ba7aa4d4a6703c0c815e5c4b2))
- **api-portfolio:** add `entity_inquiries` v1 resource ([b5cc34c](https://github.com/MTPRM/mtprm-proto/commit/b5cc34cf982305dd656e8a72758c68439aebe7b3))
- **core:** add `Entity` v2 message ([c693113](https://github.com/MTPRM/mtprm-proto/commit/c693113c333dd77b806e2e57b26d83be2e7275ab))
- **core:** add `EntityRequest` v1 message ([46d6bad](https://github.com/MTPRM/mtprm-proto/commit/46d6badc63dc6ee2ea3a80649ac10ffd3b174bfb))
- **core:** rename from `EntityRequest` to `EntityInquiry`; add `created_at` field ([29f2c7f](https://github.com/MTPRM/mtprm-proto/commit/29f2c7f72ea4f4eeb6b5fa7a6eb7b8a950f510d6))
- **proto:** add `PrefixedId` ([72e047d](https://github.com/MTPRM/mtprm-proto/commit/72e047d7bfdf237cc92e7a324fed0cb5e3aef27d))

### [16.3.1](https://github.com/mtprm/mtprm-proto/compare/16.3.0...16.3.1) (2024-09-18)

## [16.3.0](https://github.com/MTPRM/mtprm-proto/compare/16.2.0...16.3.0) (2024-09-12)

### Features

- **portfolio-beta:** add a comment about what the service does ([b298dc6](https://github.com/MTPRM/mtprm-proto/commit/b298dc67f69b690afd6c7942972556d899d251ee))

## [16.2.0](https://github.com/MTPRM/mtprm-proto/compare/16.1.1...16.2.0) (2024-09-12)

### Features

- **portfolio-beta:** update comment about what the service does ([9419bdd](https://github.com/MTPRM/mtprm-proto/commit/9419bdd8ab379a2545137febfa86c333ca8d79d8))

### [16.1.1](https://github.com/MTPRM/mtprm-proto/compare/16.1.0...16.1.1) (2024-09-12)

## [16.1.0](https://github.com/MTPRM/mtprm-proto/compare/16.0.2...16.1.0) (2024-09-12)

### Features

- **portfolio-beta:** add a comment about what the service does ([4bb07a5](https://github.com/MTPRM/mtprm-proto/commit/4bb07a5d279e89fe99c40432632130ac18ca35ab))
- **portfolio-beta:** add a comment about what the service does ([443fc40](https://github.com/MTPRM/mtprm-proto/commit/443fc407d407682ebe0d2ab69531c1ea4507d381))
- **portfolio-beta:** add a comment about what the service does ([80ea454](https://github.com/MTPRM/mtprm-proto/commit/80ea454af4b68bf2e4459b56c977bf8b1df83e4a))

### [16.0.2](https://github.com/mtprm/mtprm-proto/compare/16.0.1...16.0.2) (2024-09-09)

### [16.0.1](https://github.com/mtprm/mtprm-proto/compare/16.0.0...16.0.1) (2024-09-05)

## [16.0.0](https://github.com/MTPRM/mtprm-proto/compare/15.2.3...16.0.0) (2024-08-27)

### ⚠ BREAKING CHANGES

- **portfolio-beta:** prefer `EntityData` for `CreateRequest`
- **portfolio-beta:** return `stream`

### Features

- **portfolio-beta:** prefer `EntityData` for `CreateRequest` ([b5dff98](https://github.com/MTPRM/mtprm-proto/commit/b5dff98939c5ce9d7965649288cba7949db9222e))

### Bug Fixes

- **portfolio-beta:** return `stream` ([0c9add4](https://github.com/MTPRM/mtprm-proto/commit/0c9add475eb24322aa8a0e1fd294bb37139fffe1))

### [15.2.3](https://github.com/MTPRM/mtprm-proto/compare/15.2.2...15.2.3) (2024-08-27)

### Bug Fixes

- **portfolio-beta:** use correct package ([ae90304](https://github.com/MTPRM/mtprm-proto/commit/ae903049d2761185d06aa996e803e793a412a816))
- **portfolio-beta:** use correct package ([e032bf7](https://github.com/MTPRM/mtprm-proto/commit/e032bf752c21912eb3ac6284dc70f4ea3aff0551))
- **portfolio-beta:** use correct package ([292670e](https://github.com/MTPRM/mtprm-proto/commit/292670ed56f18cb8ee8d6d3b321139b43c69ced0))
- **portfolio-beta:** use correct package ([092d66a](https://github.com/MTPRM/mtprm-proto/commit/092d66a8c28e3b8aa233e18d42c752ed486e981b))

### [15.2.2](https://github.com/MTPRM/mtprm-proto/compare/15.2.1...15.2.2) (2024-08-27)

### [15.2.1](https://github.com/MTPRM/mtprm-proto/compare/15.2.0...15.2.1) (2024-08-27)

## [15.2.0](https://github.com/MTPRM/mtprm-proto/compare/15.1.2...15.2.0) (2024-08-27)

### Features

- **portfolio-beta:** add `entities__0__reports__combined__xlsx/v1` service ([d769e19](https://github.com/MTPRM/mtprm-proto/commit/d769e1906257679768ad4c1fb2ab15d1482cdf60))
- **portfolio-beta:** add `entities__0__reports__summary/v1` service ([ed801c3](https://github.com/MTPRM/mtprm-proto/commit/ed801c3e5e941174237728a45ff60b4366f18901))
- **portfolio-beta:** add `entities/v1` services ([bc2d9b0](https://github.com/MTPRM/mtprm-proto/commit/bc2d9b01192bcbdd879fe2ff2f330e9e92ae6e02))
- **portfolio-beta:** add `reports__combined__zip/v1` service ([b51fd89](https://github.com/MTPRM/mtprm-proto/commit/b51fd89a728f05b8c79bf47851640c257b6e71ec))

### [15.1.2](https://github.com/mtprm/mtprm-proto/compare/15.1.1...15.1.2) (2024-08-07)

### [15.1.1](https://github.com/mtprm/mtprm-proto/compare/15.1.0...15.1.1) (2024-08-07)

## [15.1.0](https://github.com/mtprm/mtprm-proto/compare/15.0.0...15.1.0) (2024-08-07)

### Features

- **core:** add more fields to `EntityData` ([ffb7e7a](https://github.com/mtprm/mtprm-proto/commit/ffb7e7a6e3922a3c4b584d3536fbbf08d112e33a))
- **proto:** add messages for `mtprm.core.pci.dss.v1` ([d13a7eb](https://github.com/mtprm/mtprm-proto/commit/d13a7eba013769683177ced4516e1d619ba52341))

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
