export class FakeTrip {
  distance: number;
  direction: string;
  subscription: string;
  totalPrice: number;
  constructor(distance: number, direction: string, subscription: string) {
    this.distance = distance;
    this.direction = direction;
    this.subscription = subscription;
  }
  async calculatePriceByDistance(): Promise<number> {
    if (this.subscription == 'PREMIUM') {
      const price = this.distance <= 5 ? 0 : (this.distance - 5) * 0.5;
      return Promise.resolve(price);
    }
    return Promise.resolve(this.distance * 0.5);
  }

  async calculateBasePrice(): Promise<number> {
    switch (this.direction) {
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

  async getTotalFee(): Promise<number> {
    const basePrice = await this.calculateBasePrice();
    const price = await this.calculatePriceByDistance();
    this.totalPrice = basePrice + price;
    return this.totalPrice;
  }
}
