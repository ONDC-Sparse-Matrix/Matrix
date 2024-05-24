let request: IDBOpenDBRequest;
let db: IDBDatabase;
let version = 1;

export interface MerchantData {
  id?: number;
  name: string;
  email: string;
  pincodes: string[];
}

export enum Stores {
  Merchants = "merchants",
}

export const initDB = (): Promise<boolean> => {
  return new Promise((resolve) => {
    request = indexedDB.open("merchants", version);

    request.onupgradeneeded = () => {
      db = request.result;
      if (!db.objectStoreNames.contains(Stores.Merchants)) {
        console.log("Creating merchants store");
        const objectStore = db.createObjectStore(Stores.Merchants, { keyPath: "id", autoIncrement: true });
        // Add other properties if needed
        objectStore.createIndex("name", "name", { unique: false });
        objectStore.createIndex("email", "email", { unique: true });
      }
    };

    request.onsuccess = () => {
      db = request.result;
      version = db.version;
      console.log("request.onsuccess - initDB", version);
      resolve(true);
    };

    request.onerror = () => {
      resolve(false);
    };
  });
};

export const addData = <T>(
  storeName: string,
  data: T
): Promise<T | string | null> => {
  return new Promise((resolve) => {
    const tx = db.transaction([storeName], "readwrite"); // Ensure database is already opened before using db.transaction
    const store = tx.objectStore(storeName);
    const request = store.add(data);
    
    tx.oncomplete = () => {
      console.log("Data added successfully");
      resolve(data);
    };

    tx.onerror = () => {
      const error = tx.error?.message;
      if (error) {
        resolve(error);
      } else {
        resolve("Unknown error");
      }
    };
  });
};


export const getStoreData = <T>(storeName: Stores): Promise<T[]> => {
  return new Promise((resolve) => {
    request = indexedDB.open("merchants");

    request.onsuccess = () => {
      console.log("request.onsuccess - getAllData");
      db = request.result;
      const tx = db.transaction(storeName, "readonly");
      const store = tx.objectStore(storeName);
      const res = store.getAll();
      res.onsuccess = () => {
        resolve(res.result);
      };
    };
  });
};
