// -----------------------------------------------------------------------------
// ZR Library                                              zr/[calendar_test.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package zr

import (
	"strings"
	"testing"
)

//  to test all items in calendar.go use:
//      go test --run Test_Calendar_
//
//  to generate a test coverage report use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

// go test --run Test_Calendar_1
func Test_Calendar_1(t *testing.T) {
	TBegin(t)
	//
	var ret Calendar
	ret.Set("2018-01-01", 20.01)
	ret.Set("2018-01-02", 21.02)
	ret.Set("2018-01-03", 22.03)
	ret.Set("2018-01-04", 23.04)
	ret.Set("2018-01-05", 24.05)
	ret.Set("2018-01-06", 10.06)
	ret.Set("2018-01-07", 11.07)
	ret.Set("2018-01-08", 12.08)
	ret.Set("2018-01-09", 13.09)
	ret.Set("2018-01-10", 9.10)
	ret.Set("2018-01-11", 8.11)
	ret.Set("2018-01-12", 7.12)
	ret.Set("2018-01-13", 6.13)
	ret.Set("2018-01-14", 5.14)
	ret.Set("2018-01-15", 4.15)
	ret.Set("2018-01-16", 3.16)
	ret.Set("2018-01-17", 2.17)
	ret.Set("2018-01-18", 1.18)
	ret.Set("2018-01-19", 0.19)
	ret.Set("2018-01-20", 20.10)
	ret.Set("2018-01-21", 19.20)
	ret.Set("2018-01-22", 18.30)
	ret.Set("2018-01-23", 17.40)
	ret.Set("2018-01-24", 16.50)
	ret.Set("2018-01-25", 15.60)
	ret.Set("2018-01-26", 14.70)
	ret.Set("2018-01-27", 13.80)
	ret.Set("2018-01-28", 12.90)
	ret.Set("2018-01-29", 11.11)
	ret.Set("2018-01-30", 10.22)
	ret.Set("2018-01-31", 9.33)
	ret.Set("2018-02-01", 8.44)
	ret.Set("2018-02-02", 7.55)
	ret.Set("2018-02-03", 6.66)
	ret.Set("2018-02-04", 5.77)
	ret.Set("2018-02-05", 4.88)
	ret.Set("2018-02-06", 3.99)
	ret.Set("2018-02-07", 2.15)
	ret.Set("2018-02-08", 1.54)
	ret.Set("2018-02-09", 0.00)
	ret.Set("2018-02-10", 1.00)
	ret.Set("2018-02-11", 2.00)
	ret.Set("2018-02-12", 3.00)
	ret.Set("2018-02-13", 4.00)
	ret.Set("2018-02-14", 5.00)
	ret.Set("2018-02-15", 6.00)
	ret.Set("2018-02-16", 7.00)
	ret.Set("2018-02-17", 8.00)
	ret.Set("2018-02-18", 9.00)
	ret.Set("2018-02-19", 10.00)
	got := ret.String()
	const expect = `
2018 JANUARY
*--------------------------------------------------------------*
|  Mon   |  Tue   |  Wed   |  Thu   |  Fri   |  Sat   |  Sun   |
|--------|--------|--------|--------|--------|--------|--------|
| 1      | 2      | 3      | 4      | 5      | 6      | 7      |
|  20.01 |  21.02 |  22.03 |  23.04 |  24.05 |  10.06 |  11.07 |
|--------|--------|--------|--------|--------|--------|--------|
| 8      | 9      | 10     | 11     | 12     | 13     | 14     |
|  12.08 |  13.09 |    9.1 |   8.11 |   7.12 |   6.13 |   5.14 |
|--------|--------|--------|--------|--------|--------|--------|
| 15     | 16     | 17     | 18     | 19     | 20     | 21     |
|   4.15 |   3.16 |   2.17 |   1.18 |   0.19 |   20.1 |   19.2 |
|--------|--------|--------|--------|--------|--------|--------|
| 22     | 23     | 24     | 25     | 26     | 27     | 28     |
|   18.3 |   17.4 |   16.5 |   15.6 |   14.7 |   13.8 |   12.9 |
|--------|--------|--------|--------|--------|--------|--------|
| 29     | 30     | 31     |        |        |        |        |
|  11.11 |  10.22 |   9.33 |        |        |        |        |
|--------|--------|--------|--------|--------|--------|--------|
|        |        |        |        |        |        |        |
|        |        |        |        |        |        |        |
*--------------------------------------------------------------*
382.06

2018 FEBRUARY
*--------------------------------------------------------------*
|  Mon   |  Tue   |  Wed   |  Thu   |  Fri   |  Sat   |  Sun   |
|--------|--------|--------|--------|--------|--------|--------|
|        |        |        | 1      | 2      | 3      | 4      |
|        |        |        |   8.44 |   7.55 |   6.66 |   5.77 |
|--------|--------|--------|--------|--------|--------|--------|
| 5      | 6      | 7      | 8      | 9      | 10     | 11     |
|   4.88 |   3.99 |   2.15 |   1.54 |      0 |      1 |      2 |
|--------|--------|--------|--------|--------|--------|--------|
| 12     | 13     | 14     | 15     | 16     | 17     | 18     |
|      3 |      4 |      5 |      6 |      7 |      8 |      9 |
|--------|--------|--------|--------|--------|--------|--------|
| 19     | 20     | 21     | 22     | 23     | 24     | 25     |
|     10 |        |        |        |        |        |        |
|--------|--------|--------|--------|--------|--------|--------|
| 26     | 27     | 28     |        |        |        |        |
|        |        |        |        |        |        |        |
|--------|--------|--------|--------|--------|--------|--------|
|        |        |        |        |        |        |        |
|        |        |        |        |        |        |        |
*--------------------------------------------------------------*
95.98
	`
	got = strings.TrimSpace(got)
	TEqual(t, got, strings.TrimSpace(expect))
} //                                                             Test_Calendar_1

// go test --run Test_Calendar_2
func Test_Calendar_2(t *testing.T) {
	TBegin(t)
	//
	var ret Calendar
	ret.Set("2020-11-01", 1.01)
	ret.Set("2020-11-02", 5.02)
	ret.Set("2020-11-03", 10.03)
	ret.Set("2020-11-04", 15.04)
	ret.Set("2020-11-05", 7.05)
	ret.Set("2020-11-06", 8.06)
	ret.Set("2020-11-30", 10.5)
	got := ret.String()
	const expect = `
2020 NOVEMBER
*--------------------------------------------------------------*
|  Mon   |  Tue   |  Wed   |  Thu   |  Fri   |  Sat   |  Sun   |
|--------|--------|--------|--------|--------|--------|--------|
|        |        |        |        |        |        | 1      |
|        |        |        |        |        |        |   1.01 |
|--------|--------|--------|--------|--------|--------|--------|
| 2      | 3      | 4      | 5      | 6      | 7      | 8      |
|   5.02 |  10.03 |  15.04 |   7.05 |   8.06 |        |        |
|--------|--------|--------|--------|--------|--------|--------|
| 9      | 10     | 11     | 12     | 13     | 14     | 15     |
|        |        |        |        |        |        |        |
|--------|--------|--------|--------|--------|--------|--------|
| 16     | 17     | 18     | 19     | 20     | 21     | 22     |
|        |        |        |        |        |        |        |
|--------|--------|--------|--------|--------|--------|--------|
| 23     | 24     | 25     | 26     | 27     | 28     | 29     |
|        |        |        |        |        |        |        |
|--------|--------|--------|--------|--------|--------|--------|
| 30     |        |        |        |        |        |        |
|   10.5 |        |        |        |        |        |        |
*--------------------------------------------------------------*
56.71`
	got = strings.TrimSpace(got)
	TEqual(t, got, strings.TrimSpace(expect))
} //                                                             Test_Calendar_2

// go test --run Test_Calendar_3
func Test_Calendar_3(t *testing.T) {
	TBegin(t)
	//
	var ret Calendar
	ret.SetWeekTotals(true)
	ret.Set("2022-04-01", 1.0)
	ret.Set("2022-04-02", 1.5)
	ret.Set("2022-04-03", 2.9)
	ret.Set("2022-04-04", 8.1)
	ret.Set("2022-04-08", 5.5)
	ret.Set("2022-04-10", 3.0)
	ret.Set("2022-04-12", 2.99)
	ret.Set("2022-04-14", 7.7)
	ret.Set("2022-04-16", 1.01)
	ret.Set("2022-04-18", 0.9)
	ret.Set("2022-04-20", 4.5)
	ret.Set("2022-04-22", 5.1)
	ret.Set("2022-04-24", 10.74)
	ret.Set("2022-04-26", 0.0)
	ret.Set("2022-04-28", 0.5)
	ret.Set("2022-04-30", 3.7)
	//
	got := ret.String()
	const expect = `
2022 APRIL
*-----------------------------------------------------------------------*
|  Mon   |  Tue   |  Wed   |  Thu   |  Fri   |  Sat   |  Sun   |  TOTAL |
|--------|--------|--------|--------|--------|--------|--------|--------|
|        |        |        |        | 1      | 2      | 3      |    (4) |
|        |        |        |        |      1 |    1.5 |    2.9 |    5.4 |
|--------|--------|--------|--------|--------|--------|--------|--------|
| 4      | 5      | 6      | 7      | 8      | 9      | 10     |   (16) |
|    8.1 |        |        |        |    5.5 |        |      3 |   16.6 |
|--------|--------|--------|--------|--------|--------|--------|--------|
| 11     | 12     | 13     | 14     | 15     | 16     | 17     |   (10) |
|        |   2.99 |        |    7.7 |        |   1.01 |        |   11.7 |
|--------|--------|--------|--------|--------|--------|--------|--------|
| 18     | 19     | 20     | 21     | 22     | 23     | 24     |   (19) |
|    0.9 |        |    4.5 |        |    5.1 |        |  10.74 |  21.24 |
|--------|--------|--------|--------|--------|--------|--------|--------|
| 25     | 26     | 27     | 28     | 29     | 30     |        |    (3) |
|        |      0 |        |    0.5 |        |    3.7 |        |    4.2 |
|--------|--------|--------|--------|--------|--------|--------|--------|
|        |        |        |        |        |        |        |        |
|        |        |        |        |        |        |        |        |
*-----------------------------------------------------------------------*
(52)
59.14
`
	got = strings.TrimSpace(got)
	TEqual(t, got, strings.TrimSpace(expect))
} //                                                             Test_Calendar_3

// end
