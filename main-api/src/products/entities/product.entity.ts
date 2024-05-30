import { Description } from './description.entity';
import { Price } from './price.entity';
import { ProductType } from '../enums/product-type.enum';
import { SubscriptionType } from '../enums/subscription-type.enum';

export class Product {
  id: string;
  product_name: string;
  product_type: ProductType;
  subscription_type: SubscriptionType;
  identifiers: string[];
  status: string;
  start_date: Date;
  end_date?: Date;
  descriptions: Description[];
  prices: Price[];
  sub_products: Product[];
}
