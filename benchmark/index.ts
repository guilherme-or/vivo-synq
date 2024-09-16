import type { ProductOrError } from "./entity";

const API_SECRET = "vivo-synq";

async function findUserProducts(userID: number): Promise<ProductOrError> {
  const url = `http://localhost:8080/users/${userID}/products`;
  const response = await fetch(url, {
    method: "GET",
    headers: {
      "X-Secret": API_SECRET,
    },
  });

  let products: ProductOrError;
  products = await response.json();

  return products;
}

const rounds = 1000;
const times: number[] = [];

for (let i = 0; i < rounds; i++) {
  const  now = Date.now();
  const randomID = Math.floor(Math.random() * 1000);
  try {
    // console.log(JSON.stringify(await findUserProducts(randomID), null, 2));
    await findUserProducts(randomID);
  } catch (error) {
    console.error(error);
  }
  const elapsed = Date.now() - now;
  console.log(`Time elapsed: ${elapsed} ms`);
  times.push(elapsed);
}

const maxTime = Math.max(...times);
const minTime = Math.min(...times);
const averageTime = times.reduce((sum, time) => sum + time, 0) / times.length;

console.log(`Maximum time elapsed p/r: ${maxTime} ms`);
console.log(`Minimum time elapsed p/r: ${minTime} ms`);
console.log(`Average time elapsed p/r: ${averageTime} ms`);