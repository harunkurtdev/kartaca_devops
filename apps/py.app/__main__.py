from typing import Union
from fastapi import FastAPI
import requests
import uvicorn
import random

class App(FastAPI):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.setup_routes()

    def setup_routes(self):
        @self.get("/")
        def read_root():
            return "Merhaba Python!"

        @self.get("/staj")
        async def get_random_city():
            try:
                response = requests.get('http://localhost:9200/cities/_search?size=100')
                data = response.json()
                cities = [hit['_source'] for hit in data['hits']['hits']]
                random_city = random.choice(cities)
                return random_city
            except Exception as e:
                return {'error': str(e)}
                
if "__main__"==__name__:
    uvicorn.run(App(),port=4444,host="0.0.0.0")