export class Ride {
  constructor(
    public id: string,
    public userId: string,
    public startAddr: string,
    public endAddr: string,
    public totalPrice: number,
    public createdAt: Date,
    public updatedAt: Date,
  ) {}
}
