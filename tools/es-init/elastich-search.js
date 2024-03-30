const { cities, countries } = require('./variables');

const sleep = (ms) => new Promise(resolve => setTimeout(resolve, ms));

async function clearIndex(index) {
    try {
        await fetch(`http://elasticsearch1:9200/${index}`, {
            method: 'DELETE'
        });
        console.log(`Index '${index}' cleared.`);
        return true;
    } catch (error) {
        console.error(`Error clearing index '${index}':`, error.message);
        return false;
    }
}


async function indexData() {
    await sleep(5000);

    
    try {
        await fetch('http://elasticsearch1:9200/cities', {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                settings: {},
                mappings: {
                    properties: {
                        il: { type: 'text' },
                        nufus: { type: 'integer' },
                        ilceler: { type: 'keyword' }
                    }
                }
            })
        });

        for (const key in cities) {
            if (Object.hasOwnProperty.call(cities, key)) {
                const city = cities[key];
                await fetch('http://elasticsearch1:9200/cities/_doc', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        il: city.il,
                        nufus: city.nufus,
                        ilceler: city.ilceler
                    })
                });
            }
        }

        // Index countries
        await fetch('http://elasticsearch1:9200/countries', {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                settings: {},
                mappings: {
                    properties: {
                        ulke: { type: 'text' },
                        nufus: { type: 'integer' },
                        baskent: { type: 'text' }
                    }
                }
            })
        });

        for (const key in countries) {
            if (Object.hasOwnProperty.call(countries, key)) {
                const country = countries[key];
                await fetch('http://elasticsearch1:9200/countries/_doc', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        ulke: country.ulke,
                        nufus: country.nufus,
                        baskent: country.baskent
                    })
                });
            }
        }

        console.log('Data indexing completed.');

    } catch (error) {
        // recursive function
        await indexData();
        console.error('Error indexing data:', error.message);
    }
}

indexData();
