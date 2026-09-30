# skvdmt-e2e-tests

## Translations
[En](./README.md)

## Описание
Набор end-to-end тестов главной страницы сайта https://skvdmt.ru/.

Осуществляется с использованием библиотеки [Chrome](https://github.com/skvdmt/chrome).

По ID селекторам ищуются узлы из dom модели документа и проводится сравнение их текстовых содержимых с указаными в тестах параметрами.

В случае не соответствия приложение закрывается со статусом 1.

## Установка
```sh
git clone https://github.com/skvdmt/skvdmt-e2e-tests
```
