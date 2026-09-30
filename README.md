# skvdmt-e2e-tests

## Translations
[Русский](./README_ru.md)

## Description
A set of end-to-end tests for the homepage of the website https://skvdmt.ru/.

Implemented using the [Chrome](https://github.com/skvdmt/chrome) library.

Nodes are located in the document's DOM model using ID selectors, and their text content is compared against the parameters specified in the tests.

In the event of a mismatch, the application terminates with status 1.

## Instalation
```sh
git clone https://github.com/skvdmt/skvdmt-e2e-tests
```