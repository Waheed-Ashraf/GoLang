abstract class Shape {
  protected color: string;
  protected filled: boolean;

  constructor(color: string = "red", filled: boolean = true) {
    this.color = color;
    this.filled = filled;
  }

  getColor(): string {
    return this.color;
  }

  setColor(color: string): void {
    this.color = color;
  }

  isFilled(): boolean {
    return this.filled;
  }

  setFilled(filled: boolean): void {
    this.filled = filled;
  }

  abstract getArea(): number;

  abstract getPerimeter(): number;

  toString(): string {
    return `Shape[color=${this.color},filled=${this.filled}]`;
  }
}

class Circle extends Shape {
  protected radius: number;

  constructor(radius: number = 1.0, color: string = "red", filled: boolean = true) {
    super(color, filled);
    this.radius = radius;
  }

  getRadius(): number {
    return this.radius;
  }

  setRadius(radius: number): void {
    this.radius = radius;
  }

  override getArea(): number {
    return Math.PI * this.radius * this.radius;
  }

  override getPerimeter(): number {
    return 2 * Math.PI * this.radius;
  }

  override toString(): string {
    return `Circle[${super.toString()},radius=${this.radius}]`;
  }
}

class Rectangle extends Shape {
  protected width: number;
  protected length: number;

  constructor(
    width: number = 1.0,
    length: number = 1.0,
    color: string = "red",
    filled: boolean = true,
  ) {
    super(color, filled);
    this.width = width;
    this.length = length;
  }

  getWidth(): number {
    return this.width;
  }

  setWidth(width: number): void {
    this.width = width;
  }

  getLength(): number {
    return this.length;
  }

  setLength(length: number): void {
    this.length = length;
  }

  override getArea(): number {
    return this.width * this.length;
  }

  override getPerimeter(): number {
    return 2 * (this.width + this.length);
  }

  override toString(): string {
    return `Rectangle[${super.toString()},width=${this.width},length=${this.length}]`;
  }
}

class Square extends Rectangle {
  constructor(side: number = 1.0, color: string = "red", filled: boolean = true) {
    super(side, side, color, filled);
  }

  getSide(): number {
    return this.width;
  }

  setSide(side: number): void {
    this.width = side;
    this.length = side;
  }

  override setWidth(side: number): void {
    this.setSide(side);
  }

  override setLength(side: number): void {
    this.setSide(side);
  }

  override toString(): string {
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
