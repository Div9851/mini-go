#!/bin/bash
assert() {
    expected="$1"
    input="$2"

    ./bin/mgo "$input" > tmp.s
    cc -o tmp tmp.s
    ./tmp
    actual="$?"

    if [ "$actual" = "$expected" ]; then
        echo "$input => $actual"
    else
        echo "$input => $expected expected, but got $actual"
        exit 1
    fi
}

assert 0 "0;"
assert 42 "42;"
assert 3 "1+2;"
assert 2 "5-3;"
assert 36 "4*9;"
assert 5 "15/3;"
assert 0 "(1+2+3)*6-9*4;"
assert 1 "((1+2)*4-5)/7;"
assert 30 "a:=10;b:=20;a+b;"
assert 20 "a:=3*4+5;b:=a+3;b;"
assert 9 "a:=6; b:=9; a,b = b,a; a;"
assert 4 "a,b := 3,4; b;"
echo OK
