import { openDB, DBSchema, IDBPDatabase } from "idb";
import { PincodeData, PincodeDataCache, MerchantData } from "./types";

const DB_NAME = "cached-merchants";
const STORE_NAME = "merchants-list";
const STORE_KEY = "cache";

interface MyDB extends DBSchema {
  "merchants-list": {
    key: string;
    value: { cache: PincodeDataCache; id: string };
  };
}

const initDB = async (): Promise<IDBPDatabase<MyDB>> => {
  return openDB<MyDB>(DB_NAME, 1, {
    upgrade(db) {
      if (!db.objectStoreNames.contains(STORE_NAME)) {
        db.createObjectStore(STORE_NAME, { keyPath: "id" });
      }
    },
  });
};

export const updateCache = async (
  newCache: PincodeDataCache
): Promise<void> => {
  const { cache: prevCache } = (await getCache()) || {
    cache: undefined,
    id: STORE_KEY,
  };
  const db = await initDB();
  const tx = db.transaction(STORE_NAME, "readwrite");
  const store = tx.objectStore(STORE_NAME);

  let updatedCache: PincodeDataCache;
  if (prevCache) {
    updatedCache = mergePincodeDataCaches(prevCache, newCache);
    console.log("Merged cache", updatedCache);
  } else {
    updatedCache = newCache;
    console.log("No previous cache found, assigning new cache, ", updatedCache);
  }

  //@ts-ignore
  await store.put({ cache: updatedCache, id: STORE_KEY });
  await tx.done;
};

function mergePincodeDataCaches(
  previousCache: PincodeDataCache,
  newCache: PincodeDataCache
): PincodeDataCache {
  const pincodeMap: Map<string, MerchantData[]> = new Map();
  const addToPincodeMap = (cache: PincodeDataCache) => {
    if (!Array.isArray(cache)) {
      console.error("cache is not an array", cache);
      return;
    }

    for (const entry of cache) {
      if (!pincodeMap.has(entry.pincode)) {
        pincodeMap.set(entry.pincode, entry.merchantList);
      } else {
        const existingMerchants = pincodeMap.get(entry.pincode)!;
        for (const merchant of entry.merchantList) {
          if (!existingMerchants.some((m) => m.email === merchant.email)) {
            existingMerchants.push(merchant);
            console.log("Added merchant", merchant);
          }
        }
      }
    }
  };
  addToPincodeMap(previousCache);
  addToPincodeMap(newCache);

  const updatedCache: PincodeDataCache = [];
  for (const [pincode, merchantList] of pincodeMap) {
    updatedCache.push({ pincode, merchantList });
  }

  return updatedCache;
}

export function searchCache(pincode: string) {
  return new Promise<PincodeData["current"] | undefined>(async (resolve) => {
    const { cache } = (await getCache()) || { cache: undefined, id: STORE_KEY };
    if (!cache) {
      resolve(undefined);
      return;
    }
    const result = cache.find((entry) => entry.pincode === pincode);
    resolve(result ? result : undefined);
  });
}

export const getCache = async (): Promise<{
  cache: PincodeDataCache;
  id: string;
}> => {
  const db = await initDB();
  const tx = db.transaction(STORE_NAME, "readonly");
  const store = tx.objectStore(STORE_NAME);
  const cache = await store.get(STORE_KEY);
  await tx.done;
  return cache as { cache: PincodeDataCache; id: string };
};

export const clearCache = async (): Promise<void> => {
  const db = await initDB();
  const tx = db.transaction(STORE_NAME, "readwrite");
  const store = tx.objectStore(STORE_NAME);
  await store.clear();
  await tx.done;
};
