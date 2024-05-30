import { createClient } from 'redis';
import { Module } from '@nestjs/common';

const REDIS_CLIENT = 'REDIS_CLIENT';
const REDIS_OPTIONS = 'REDIS_OPTIONS';

@Module({
  providers: [
    {
      provide: REDIS_OPTIONS,
      useValue: {
        url: 'redis://localhost:6379',
      },
    },
    {
      inject: [REDIS_OPTIONS],
      provide: REDIS_CLIENT,
      useFactory: async (options: { url: string }) => {
        const client = createClient(options);
        await client.connect();
        return client;
      },
    },
  ],
  exports: [REDIS_CLIENT],
})
class RedisModule {}

export { REDIS_CLIENT, REDIS_OPTIONS, RedisModule };
