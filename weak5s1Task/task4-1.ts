


// ================ 4-1 ================

class Circle {
  private radius: number;
  private color: string;

  constructor(radius: number, color: string) {
    this.radius = radius;
    this.color = color;
  }

  getRadius(): number {
    return this.radius;
  }
  
  setRadius(radius: number): void {
    this.radius = radius;
  }
  getColor(): string {
    return this.color;
  }
  setColor(color: string): void {
    this.color = color;
  }
  getArea(): number {
    return Math.PI * this.radius * this.radius;
  }
  toString(): string {
    return `Circle[radius=${this.radius}, color=${this.color}]`;
  }
}

class Cylinder extends Circle {
  private height: number;
  
  constructor(radius: number, color: string, height: number) {
    super(radius, color);
    this.height = height;
  }


  getHeight(): number {
    return this.height;
  }
  setHeight(height: number): void {
    this.height = height;
  }
  getVolume(): number {
    return super.getArea() * this.height;
  }
  toString(): string {
    return `Cylinder[height=${this.height}, radius=${this.getRadius()}, color=${this.getColor()}]`;
  }
}

// ================ TestCylinder ================

function TestCylinder(): void {
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

