import { Inject, Injectable } from '@nestjs/common';
import { ViewProductDTO } from './dto/view-product.dto';
import {
  RedisClientType,
  RedisFunctions,
  RedisModules,
  RedisScripts,
} from 'redis';

@Injectable()
export class ProductsService {
  constructor(
    @Inject('REDIS_CLIENT')
    private readonly redis: RedisClientType<
      RedisModules,
      RedisFunctions,
      RedisScripts
    >,
  ) {}

  async findById(id: number): Promise<ViewProductDTO> {
    const p = new ViewProductDTO();
    p.id = String(id);

    const r = await this.redis.get('mykey');
    console.log('redis.get', r, typeof r);

    return p;
  }
}
