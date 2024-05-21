import requests
import concurrent.futures
from pymongo import MongoClient

def make_api_call(api_url):
    response = requests.get(api_url)
    return response.json()

def save_to_database(response_text):
    client = MongoClient('mongodb://localhost:27017/')
    db = client['remoteServer_A']
    collection = db['merchants']
    data = {
        'name': response_text
    }
    collection.insert_one(data)
    client.close()

N = 2000
batch_size = 50

api_url = 'https://randomuser.me/api/'

def process_batch(batch_indices):
    for index in batch_indices:
        try:
            api_response = make_api_call(api_url)
            response_text = f"{api_response['results'][0]['name']['first']} {api_response['results'][0]['name']['last']}"
            save_to_database(response_text)
        except Exception as e:
            print(f"Error processing index {index}: {str(e)}")

num_batches = (N + batch_size - 1) // batch_size
batch_indices_list = [range(i * batch_size, min((i + 1) * batch_size, N)) for i in range(num_batches)]

with concurrent.futures.ThreadPoolExecutor() as executor:
    executor.map(process_batch, batch_indices_list)

print(f"{N} API calls completed and saved to the database.")
