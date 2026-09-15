"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
class Shape {
    color;
    filled;
    constructor(color = "red", filled = true) {
        this.color = color;
        this.filled = filled;
    }
    getColor() {
        return this.color;
    }
    setColor(color) {
        this.color = color;
    }
    isFilled() {
        return this.filled;
    }
    setFilled(filled) {
        this.filled = filled;
    }
    toString() {
        return `Shape[color=${this.color},filled=${this.filled}]`;
    }
}
class Circle extends Shape {
    radius;
    constructor(radius = 1.0, color = "red", filled = true) {
        super(color, filled);
        this.radius = radius;
    }
    getRadius() {
        return this.radius;
    }
    setRadius(radius) {
        this.radius = radius;
    }
    getArea() {
        return Math.PI * this.radius * this.radius;
    }
    getPerimeter() {
        return 2 * Math.PI * this.radius;
    }
    toString() {
        return `Circle[${super.toString()},radius=${this.radius}]`;
    }
}
class Rectangle extends Shape {
    width;
    length;
    constructor(width = 1.0, length = 1.0, color = "red", filled = true) {
        super(color, filled);
        this.width = width;
        this.length = length;
    }
    getWidth() {
        return this.width;
    }
    setWidth(width) {
        this.width = width;
    }
    getLength() {
        return this.length;
    }
    setLength(length) {
        this.length = length;
    }
    getArea() {
        return this.width * this.length;
    }
    getPerimeter() {
        return 2 * (this.width + this.length);
    }
    toString() {
        return `Rectangle[${super.toString()},width=${this.width},length=${this.length}]`;
    }
}
class Square extends Rectangle {
    constructor(side = 1.0, color = "red", filled = true) {
        super(side, side, color, filled);
    }
    getSide() {
        return this.width;
    }
    setSide(side) {
        this.width = side;
        this.length = side;
    }
    setWidth(side) {
        this.setSide(side);
    }
    setLength(side) {
        this.setSide(side);
    }
    toString() {
        return `Square[${super.toString()}]`;
    }
}
// ================ TestShape ================
const s1 = new Circle(5.5, "red", false);
console.log(s1.toString());
console.log(s1.getArea());
console.log(s1.getPerimeter());
console.log(s1.getColor());
console.log(s1.isFilled());
console.log(s1.getRadius());
const s2 = new Rectangle(1.0, 2.0, "red", false);
console.log(s2.toString());
console.log(s2.getArea());
console.log(s2.getPerimeter());
console.log(s2.getColor());
console.log(s2.getLength());
const s3 = new Square(6.6);
console.log(s3.toString());
console.log(s3.getArea());
console.log(s3.getPerimeter());
console.log(s3.getColor());
console.log(s3.getSide());
s3.setSide(7.7);
console.log(s3.toString());
s3.setWidth(8.8);
console.log(s3.toString());
s3.setLength(9.9);
console.log(s3.toString());
//# sourceMappingURL=task6-1.js.map