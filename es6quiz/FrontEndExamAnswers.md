# Front End Exam Solutions

YouTube clone task skipped for now.

## Question 1

1. Hoisting means JavaScript moves declarations to the top of their scope during compilation. `var` declarations are hoisted with `undefined`; function declarations are hoisted with their body; `let` and `const` are hoisted but stay in the temporal dead zone until initialized.

2. `super` is used inside a child class to call the parent constructor or parent methods.

3. `let` and `const` are block-scoped. `let` can be reassigned, `const` cannot be reassigned. `var` is function-scoped, can be redeclared, and is hoisted with `undefined`.

4. Rest parameters collect remaining function arguments into an array. Example: `function sum(...nums) {}`. Arrow functions are shorter function expressions and keep lexical `this`. They are useful for callbacks and small functions. They differ from normal functions because they do not have their own `this`, `arguments`, or `prototype`, and cannot be used as constructors.

5. `readonly` means the user cannot edit the textarea value, but the value is still submitted with the form. `disabled` means the field cannot be edited or focused, and its value is not submitted.

6. CSS units can be absolute or relative. Examples: `px`, `cm`, `mm`, `%`, `em`, `rem`, `vw`, `vh`, `vmin`, `vmax`.

7. The CSS property for changing the font face is `font-family`.

8. Two ways to center a div inside another div:

```css
.parent {
  display: flex;
  justify-content: center;
  align-items: center;
}
```

```css
.child {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
```

## Question 2

1. False
2. False
3. True
4. False
5. False
6. True
7. True
8. True

## Question 3

1. JavaScript is a synchronous, blocking, single-threaded language.
2. Encapsulation.

## Question 4

1. Promise output:

```txt
Error: The Fails!
The Fails!
The Fails!
```

2. `member.getFullName()` output:

```txt
TypeError: member.getFullName is not a function
```

3. Curried sum output:

```txt
6
8
```

4. Chameleon output:

```txt
TypeError: freddie.colorChange is not a function
```

5. Prompt/type output:

```txt
boolean
string
```

6. `setTimeout` output order:

```txt
0
1
4
2
3
```

7. Counter output:

```txt
ReferenceError: i is not defined
```

8. Dynamic object key output:

```txt
hello world
10
```

9. Euros reduce output:

```txt
[59.52, 83.7, 93]
```

10. Names slice/spread output:

```txt
["Batman", "Bane"]
```

## Question 5

Full code is in `frontEndExam.js`, `frontEndExam.html`, and `frontEndExam.css`.

9. Better implementations for the bad constructor script:

```js
function MyObject(name, message) {
  this.name = String(name);
  this.message = String(message);
}

MyObject.prototype.getName = function () {
  return this.name;
};

MyObject.prototype.getMessage = function () {
  return this.message;
};
```

```js
class BetterObject {
  constructor(name, message) {
    this.name = String(name);
    this.message = String(message);
  }

  getName() {
    return this.name;
  }

  getMessage() {
    return this.message;
  }
}
```
