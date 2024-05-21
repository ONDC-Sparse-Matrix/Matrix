import random
from pymongo import MongoClient

def generate_random_data():
    pin_code = f"{random.randint(202000,210500)}"

    merchant_ids = get_random_merchant_ids()

    array_length = random.randint(1, 10)

    selected_ids = random.sample(merchant_ids, array_length)

    return {
        "pin_code": pin_code,
        "merchant_ids": selected_ids
    }

def get_random_merchant_ids():
    client = MongoClient('mongodb://localhost:27017/')
    db = client['remoteServer_C']

    merchant_ids = [str(doc['_id']) for doc in db['merchants'].find({}, {'_id': 1})]

    client.close()
    print(merchant_ids)
    return merchant_ids

def insert_data_into_map_collection(data):
    client = MongoClient('mongodb://localhost:27017/')
    db = client['remoteServer_C']

    db['map'].insert_one(data)

    client.close()

if __name__ == "__main__":
    num_entries = 10

    for _ in range(num_entries):
        data_entry = generate_random_data()
        insert_data_into_map_collection(data_entry)

    print(f"{num_entries} entries inserted into the 'map' collection.")
