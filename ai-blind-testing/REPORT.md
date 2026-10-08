# Task 3 — AI blind test generation

> Homework — Task 3 (Lesson 6: File I/O, JSON and Testing).
> This task is **not auto-graded** by CI — it's reviewed by your mentor.
> The workflow only checks that you actually filled this file in
> (see the "Task 3 — AI report present" step in the Actions log).

## Instructions

1. Pick one function you wrote for Task 1 or Task 2 (e.g. `LoadTodos` or
   `ValidateEmail`).
2. Give an AI assistant **only the function signature** — no explanation
   of the logic, no existing test file. Ask it to generate a full
   table-driven test suite for that signature.
3. Run the AI-generated tests against your implementation.
4. Fill in the sections below.

---

## Function under test

<!-- e.g. `func ValidateEmail(s string) bool` from validate/validate.go -->

## Prompt you gave the AI

<!-- Paste the exact prompt you used. -->

## Edge cases the AI found that you had missed

<!-- List them, and say whether they exposed a real bug. -->

## Edge cases you had that the AI missed

<!-- List them, and say why you think the AI didn't think of them. -->

## Cases where the AI's expected output was wrong

<!-- AI can be confidently wrong about what the "correct" output should
     be — did that happen here? -->

## What you'd change about your own test-writing process after this

<!-- A few sentences of reflection. -->


---

# Завдання 3 — сліпа генерація тестів за допомогою ШІ

> Домашнє завдання — Завдання 3 (Урок 6: Робота з файлами, JSON та тестування).
> Це завдання **не перевіряється автоматично** в CI — його перевіряє ваш ментор.
> Робочий процес лише перевіряє, що ви заповнили цей файл
> (див. крок «Task 3 — AI report present» у журналі Actions).

## Інструкції

1. Виберіть одну функцію, яку ви написали для Завдання 1 або 2 (наприклад,
   `LoadTodos` або `ValidateEmail`).
2. Дайте ШІ-асистенту **лише сигнатуру функції** — без пояснення логіки та
   без наявного файлу тестів. Попросіть згенерувати повний табличний
   набір тестів для цієї сигнатури.
3. Запустіть згенеровані ШІ тести на вашій реалізації.
4. Заповніть розділи нижче.

---

## Функція, що тестується

`func ValidateEmail(s string) bool` з `validate/validate.go`.

## Промпт, який ви дали ШІ

> для функції ValidateEmail напиши прості табличні тести

ШІ отримав лише назву й сигнатуру функції, без реалізації та без мого
`validate_test.go`. Згенеровані тести збережено у
`validate/validate_ai_test.go` (33 випадки).

## Крайні випадки, які знайшов ШІ, а ви пропустили

Із 33 згенерованих випадків 11 впали на моїй реалізації. Більшість із них
указують на реальні прогалини у `ValidateEmail`:

`user@example` (домен без крапки/TLD) — функція повертає `true`;
послідовні крапки: `us..er@example.com` і `user@example..com`;
крапка на початку або в кінці локальної частини: `.user@example.com`, `user.@example.com`;
крапка на початку домену: `user@.example.com`;
дефіс на початку домену: `user@-example.com`;
недопустимі символи: `us(er@example.com`, `user@exa_mple.com`.

Усе це реальні помилки: моя валідація перевіряла лише порожній рядок,
пробіли, рівно один `@`, непорожні частини та крапку в кінці домену. Раніше
я навіть не подумав про такі випадки, хоча в коментарі до функції сам
згадував крапки як те, що треба задокументувати.

## Крайні випадки, які були у вас, а ШІ пропустив

Обмеження загальної довжини в 40 символів (мій власний ліміт). ШІ
орієнтувався на стандарт (локальна частина ≤ 64, адреса ≤ 254) і не міг
знати про моє рішення, бо бачив лише сигнатуру.
Межа 40/41 символів: ШІ не тестував саме цю границю з тієї самої причини.
Табуляція й пробіли були в обох наборах, тому тут розбіжності немає.

## Випадки, де очікуваний результат ШІ був хибним

Є два випадки, де ШІ «помилився» лише відносно моїх рішень, а не відносно
стандарту:
`max valid local part 64` — ШІ очікував `true` для адреси з 64-символьною
локальною частиною. За RFC це справді валідно, але моя функція навмисно
обмежує всю адресу 40 символами, тому повертає `false`. Тут слід змінити
не код, а очікування (або ліміт).




Загалом ШІ впевнено обирав очікуваний результат за «типовим» суворим
стандартом email, не знаючи вимог конкретного проєкту.

## Що ви змінили б у власному процесі написання тестів після цього


Я писав тести лише ті які по завданню ну і яки швидко прийшли в на думку а ШІ охопив всі випадки на які його наренували.
Що я змінив одразу у ШІ питати найкращі практики по валідації щоб були охоплені всі випадки окрім зазаначених наприклад мною згідно ТЗ як допустимі

