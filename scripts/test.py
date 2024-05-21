import requests

def make_api_call(api_url):
    response = requests.get(api_url)
    return response.json()

api_url = 'https://randomuser.me/api/'

api_response = make_api_call(api_url)

full_name = f"{api_response['results'][0]['name']['first']} {api_response['results'][0]['name']['last']}"
print(f"{full_name}")
