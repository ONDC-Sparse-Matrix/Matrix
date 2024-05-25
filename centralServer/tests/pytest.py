import requests
import threading
import time

# Replace 'http://example.com/api' with your actual endpoint
url = 'http://localhost:3001/'
num_requests = 5000

def send_request():
    try:
        response = requests.get(url)
        print(f"Status code: {response.status_code}")
    except Exception as e:
        print(f"Error: {e}")

# Run the requests in parallel using threads
threads = []

for _ in range(num_requests):
    thread = threading.Thread(target=send_request)
    threads.append(thread)

start_time = time.time()

for thread in threads:
    thread.start()

for thread in threads:
    thread.join()

end_time = time.time()

total_time = end_time - start_time
requests_per_second = num_requests / total_time

print(f"Total requests: {num_requests}")
print(f"Total time: {total_time} seconds")
print(f"Requests per second: {requests_per_second}")
