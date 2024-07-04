# Changelog

## [1.2.0](https://github.com/mattgialelis/dutycontroller/compare/v1.1.1...v1.2.0) (2024-07-04)


### Features

* Try return on the finalizer addtion to reduce the mulitple calls ([e6af51d](https://github.com/mattgialelis/dutycontroller/commit/e6af51d0a52d895ce69b25f1d24cc587d1b3e4ff))
* Update to add condtions on Service and business services ([d37584d](https://github.com/mattgialelis/dutycontroller/commit/d37584dcdff6a6578278a82b612c515a49940baa))

## [1.1.1](https://github.com/mattgialelis/dutycontroller/compare/v1.1.0...v1.1.1) (2024-07-03)


### Bug Fixes

* Add requeue ([941f686](https://github.com/mattgialelis/dutycontroller/commit/941f6868d6012da76373966ca49d8ddb876b0678))

## [1.1.0](https://github.com/mattgialelis/dutycontroller/compare/v1.0.3...v1.1.0) (2024-07-03)


### Features

* BREAKING moved business service to cluster scoped ([53a4f0d](https://github.com/mattgialelis/dutycontroller/commit/53a4f0d6f686e542061b3d6248b8d0fbab46818e))


### Bug Fixes

* Fix the chart monitoring ports ([e6dc646](https://github.com/mattgialelis/dutycontroller/commit/e6dc64696b1ceea47b2211e38af23dfd9cd84287))
* helm fixes to better fit the the controller ([fbf1f3b](https://github.com/mattgialelis/dutycontroller/commit/fbf1f3bea701cb65d143cec7a612076036a7c396))
* Pipline to ensure not triggering when not really needed ([1e5ddcf](https://github.com/mattgialelis/dutycontroller/commit/1e5ddcfa7023c6309a416c03c48ccb7995dfa41a))
* working helm chart ([23d459a](https://github.com/mattgialelis/dutycontroller/commit/23d459a6d171c576692c86a187fb1ef2f42f90af))


### Miscellaneous Chores

* release 1.1.0 ([516204f](https://github.com/mattgialelis/dutycontroller/commit/516204fce52015265d01d2e5dbbe37d08dc8f134))

## [1.0.3](https://github.com/mattgialelis/dutycontroller/compare/v1.0.2...v1.0.3) (2024-07-02)


### Bug Fixes

* try update to use mike to see if that works for managing the repos fiels ([06f53d3](https://github.com/mattgialelis/dutycontroller/commit/06f53d3d370fc6b45763018304e591d0ec7282f2))

## [1.0.2](https://github.com/mattgialelis/dutycontroller/compare/v1.0.1...v1.0.2) (2024-07-02)


### Bug Fixes

* try fix pipeline to ensure that mkdocs and other doesnt overide ([29f2c12](https://github.com/mattgialelis/dutycontroller/commit/29f2c12d67bc97eabdcbfb35537328e9b1748442))

## [1.0.1](https://github.com/mattgialelis/dutycontroller/compare/v1.0.0...v1.0.1) (2024-07-01)


### Bug Fixes

* pipelines and bump release ([8918ccc](https://github.com/mattgialelis/dutycontroller/commit/8918ccc8ece528bc6de42783aaae5a0bd08c87c2))

## 1.0.0 (2024-06-29)


### Features

* Init code to main branch ([#1](https://github.com/mattgialelis/dutycontroller/issues/1)) ([52fa18e](https://github.com/mattgialelis/dutycontroller/commit/52fa18e95f309cb0406358f686484dfdaa55880a))


### Bug Fixes

* Attempt fix of chart relaser ([9c89d8a](https://github.com/mattgialelis/dutycontroller/commit/9c89d8a674a48a2c6e4ded1e1c354ffb6bb7cc46))
* Change discription to bump chart to test action ([94e4b23](https://github.com/mattgialelis/dutycontroller/commit/94e4b2302000c413218c418fe9f8babbccc14245))
* Fix the helm chart add some features for passing in the Token ([154a378](https://github.com/mattgialelis/dutycontroller/commit/154a3786d344938c957e1476a0c60a7ba4890360))
* Fixed charts_dir on helm releaser ([b312a36](https://github.com/mattgialelis/dutycontroller/commit/b312a364a08059a233b2f5c008ed4cdaf0660173))
* Fixed spacing in deploy.yml breaking pipeline ([e94f5e4](https://github.com/mattgialelis/dutycontroller/commit/e94f5e44f45669694cb1e426cbcd4af8dfa56d5e))
* Make docs only publish on release ([58c7920](https://github.com/mattgialelis/dutycontroller/commit/58c792036038af9ea473cb5c40100b60fd7763d5))
* set the releaser chart to have the correct path and add more perms for release please ([6a0b0bb](https://github.com/mattgialelis/dutycontroller/commit/6a0b0bb268f4ac3962e3707df7339684be7ab619))
* The pipelines to do the first build ([#2](https://github.com/mattgialelis/dutycontroller/issues/2)) ([7b0ab0a](https://github.com/mattgialelis/dutycontroller/commit/7b0ab0ad984712a8273406931a3f595684f65f5c))
