import { Description } from './description.entity';
import { Price } from './price.entity';
import { ProductType } from '../enums/product-type.enum';
import { SubscriptionType } from '../enums/subscription-type.enum';

export class Product {
  id: string;
  productName: string;
  productType: ProductType;
  subscriptionType: SubscriptionType;
  identifiers: string[];
  status: string;
  startDate: Date;
  endDate?: Date;
  descriptions: Description[];
  prices: Price[];
  subProducts: Product[];
}
