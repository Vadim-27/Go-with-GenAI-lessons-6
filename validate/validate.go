// Package validate надає прості синтаксичні валідатори для поширених
// форматів користувацького вводу.
//
// Домашнє завдання — Завдання 2 (Урок 6: Робота з файлами, JSON та тестування):
// Реалізуйте ValidateEmail та/або ValidatePhone нижче (ментор може
// попросити лише одну з них) і розширте таблиці тестів у validate_test.go
// щонайменше до 8 випадків кожна, включно з крайніми випадками.
package validate

import (
	"strings"
	"unicode"
)

// ValidateEmail повідомляє, чи є s синтаксично коректною email-адресою.
//
// Прийняті рішення:
//   - порожній рядок, пробільні символи та довжина понад 40 байтів відхиляються;
//   - має бути рівно один "@" з непорожніми локальною частиною та доменом;
//   - локальна частина: літери (включно з unicode), цифри та символи
//     "._+-%"; не може починатися чи закінчуватися крапкою і містити ".."
//   - домен: лише ASCII-літери, цифри та "-"; мінімум дві мітки через крапку
//     (тобто потрібен TLD); мітка не буває порожньою (це відсікає крапку
//     на початку, в кінці та послідовні крапки) і не починається/не
//     закінчується дефісом.
func ValidateEmail(s string) bool {
	if s == "" || len(s) > 40 {
		return false
	}
	if strings.IndexFunc(s, unicode.IsSpace) >= 0 {
		return false
	}
	if strings.Count(s, "@") != 1 {
		return false
	}
	local, domain, _ := strings.Cut(s, "@")
	return validEmailLocal(local) && validEmailDomain(domain)
}

func validEmailLocal(local string) bool {
	if local == "" || strings.HasPrefix(local, ".") || strings.HasSuffix(local, ".") ||
		strings.Contains(local, "..") {
		return false
	}
	for _, r := range local {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !strings.ContainsRune("._+-%", r) {
			return false
		}
	}
	return true
}

func validEmailDomain(domain string) bool {
	labels := strings.Split(domain, ".")
	if len(labels) < 2 {
		return false
	}
	for _, label := range labels {
		if label == "" || label[0] == '-' || label[len(label)-1] == '-' {
			return false
		}
		for _, r := range label {
			isAlnum := r < unicode.MaxASCII && (unicode.IsLetter(r) || unicode.IsDigit(r))
			if !isAlnum && r != '-' {
				return false
			}
		}
	}
	return true
}

// ValidatePhone повідомляє, чи є s синтаксично коректним номером телефону.
//
// TODO: реалізуйте цю функцію. Вирішіть, які формати ви приймаєте
// (наприклад, "+380501234567", "050-123-4567"), і задокументуйте
// це рішення тут. Щонайменше вона має відхиляти порожній рядок і будь-яке
// значення, що містить літери.
//
// Прийняте рішення: допустимі лише ASCII-цифри, роздільники "-", "(", ")",
// "." та необов'язковий "+" на самому початку. Пробіли відхиляються.
// Кількість цифр має бути від 10 до 15 (на кшталт E.164).
func ValidatePhone(s string) bool {
	digits := 0
	for i, r := range s {
		switch {
		case r >= '0' && r <= '9':
			digits++
		case r == '+':
			if i != 0 {
				return false
			}
		case r == '-', r == '(', r == ')', r == '.':
		default:
			return false
		}
	}
	return digits >= 10 && digits <= 15
}
