export interface ITrip {
  calculatePriceByDistance(distance: number): number;
  calculateBasePrice(direction: string): number;
}

export class FakeTrip {
  distance: number;
  direction: string;
  constructor(distance: number, direction: string) {
    this.distance = distance;
    this.direction = direction;
  }
  async calculatePriceByDistance(distance: number): Promise<number> {
    return Promise.resolve(distance * 0.5);
  }

  async calculateBasePrice(direction: string): Promise<number> {
    switch (direction) {
      case 'INTRA_MUROS':
        return Promise.resolve(30);
      case 'EXTRA_MUROS':
        return Promise.resolve(20);
      case 'PARIS':
        return Promise.resolve(10);
      default:
        throw new Error('direction error');
    }
  }

  async getTotalFee() {
    const basePrice = await this.calculateBasePrice(this.direction);
    const price = await this.calculatePriceByDistance(this.distance);
    return basePrice + price;
  }
}
