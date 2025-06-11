export class FakeTrip {
  distance: number;
  direction: string;
  subscription: string;
  totalPrice: number;
  isUberX: boolean;
  constructor(
    distance: number,
    direction: string,
    subscription: string,
    isUberX = false,
  ) {
    if (isUberX && distance < 3) {
      throw new Error('cannot UberX under 3 km');
    }
    this.distance = distance;
    this.direction = direction;
    this.subscription = subscription;
    this.isUberX = isUberX;
  }
  async calculatePriceByDistance(): Promise<number> {
    if (this.subscription == 'PREMIUM') {
      const price = this.distance <= 5 ? 0 : (this.distance - 5) * 0.5;
      return Promise.resolve(price);
    }
    return Promise.resolve(this.distance * 0.5);
  }

  async calculateBasePrice(): Promise<number> {
    const basePrice = this.handleUberX();
    switch (this.direction) {
      case 'INTRA_MUROS':
        return Promise.resolve(basePrice + 30);
      case 'EXTRA_MUROS':
        return Promise.resolve(basePrice + 20);
      case 'PARIS':
        return Promise.resolve(basePrice + 10);
      default:
        throw new Error('direction error');
    }
  }

  handleUberX() {
    return this.isUberX ? 10 : 0;
  }

  async getTotalFee(): Promise<number> {
    const basePrice = await this.calculateBasePrice();
    const price = await this.calculatePriceByDistance();
    this.totalPrice = basePrice + price;
    return this.totalPrice;
  }
}
