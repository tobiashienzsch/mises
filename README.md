# Mises

Golang helper library for http microservices. Uses mux (router) & pop for database connections.

## Status

|                                                                 License                                                                  |                                                          Unit test                                                          |                                                              Coverage                                                              |                                     Issues                                      |
| :--------------------------------------------------------------------------------------------------------------------------------------: | :-------------------------------------------------------------------------------------------------------------------------: | :--------------------------------------------------------------------------------------------------------------------------------: | :-----------------------------------------------------------------------------: |
| [![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://github.com/tobiashienzsch/mises/blob/master/LICENSE) | [![Build Status](https://travis-ci.org/tobiashienzsch/mises.svg?branch=master)](https://travis-ci.org/tobiashienzsch/mises) | [![codecov](https://codecov.io/gh/tobiashienzsch/mises/branch/master/graph/badge.svg)](https://codecov.io/gh/tobiashienzsch/mises) | ![GitHub issues](https://img.shields.io/github/issues/tobiashienzsch/mises.svg) |

## Structure

| package       | description         | comments |
| ------------- | ------------------- | -------- |
| mises/        | common service code |          |
| mises/cmd     | cobra bindings      |          |
| mises/config  | global config       |          |
| mises/runtime | version info        |          |

## ToDo

- Base:
  - JWT middleware
  - Test setup (database cleaning?)
