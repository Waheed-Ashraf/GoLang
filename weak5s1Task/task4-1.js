"use strict";
// ================ 4-1 ================
Object.defineProperty(exports, "__esModule", { value: true });
class Circle {
    radius;
    color;
    constructor(radius, color) {
        this.radius = radius;
        this.color = color;
    }
    getRadius() {
        return this.radius;
    }
    setRadius(radius) {
        this.radius = radius;
    }
    getColor() {
        return this.color;
    }
    setColor(color) {
        this.color = color;
    }
    getArea() {
        return Math.PI * this.radius * this.radius;
    }
    toString() {
        return `Circle[radius=${this.radius}, color=${this.color}]`;
    }
}
class Cylinder extends Circle {
    height;
    constructor(radius, color, height) {
        super(radius, color);
        this.height = height;
    }
    getHeight() {
        return this.height;
    }
    setHeight(height) {
        this.height = height;
    }
    getVolume() {
        return super.getArea() * this.height;
    }
    toString() {
        return `Cylinder[height=${this.height}, radius=${this.getRadius()}, color=${this.getColor()}]`;
    }
}
// ================ TestCylinder ================
function TestCylinder() {
    const c1 = new Cylinder(2.0, "red", 10.0);
    console.log(c1.toString());
    console.log("radius is: " + c1.getRadius());
    console.log("color is: " + c1.getColor());
    console.log("height is: " + c1.getHeight());
    console.log("base area is: " + c1.getArea());
    console.log("volume is: " + c1.getVolume());
    c1.setRadius(3.0);
    c1.setColor("blue");
    c1.setHeight(5.0);
    console.log(c1.toString());
    console.log("radius is: " + c1.getRadius());
    console.log("color is: " + c1.getColor());
    console.log("height is: " + c1.getHeight());
    console.log("base area is: " + c1.getArea());
    console.log("volume is: " + c1.getVolume());
}
TestCylinder();
//# sourceMappingURL=task4-1.js.map