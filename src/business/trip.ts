export interface ITrip {
  calculate(): Promise<number>;
}

export class FakeTrip implements ITrip {
  public distance: number;
  constructor(distance: number) {
    this.distance = distance;
  }
  calculate(): Promise<number> {
    return Promise.resolve(this.distance);
  }
}
