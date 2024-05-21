import random
from pymongo import MongoClient
from bson import ObjectId

client = MongoClient('mongodb://localhost:27017')
db = client['remoteServer_A']

map_collection = db['map']
merchants_collection = db['merchants']

def generate_random_count():
    return random.randint(5, 15)

for map_document in map_collection.find():
    random_count = generate_random_count()

    random_merchants = random.sample(list(merchants_collection.find()), random_count)

    merchant_ids = [str(merchant['_id']) for merchant in random_merchants]

    map_collection.update_one(
        {'_id': map_document['_id']},
        {'$set': {'merchant_ids': merchant_ids}}
    )

print("Operation completed successfully.")
