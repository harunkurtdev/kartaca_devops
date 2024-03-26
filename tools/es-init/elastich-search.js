// Function to sleep for a given number of milliseconds
const sleep = (ms) => new Promise(resolve => setTimeout(resolve, ms));

async function indexData() {
  await sleep(5000); // Sleep for 5 seconds

  try {
    // Create index "cities" with mapping
    await fetch('http://localhost:9200/cities', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        settings: {},
        mappings: {
          properties: {
            name: { type: 'text' },
            country: { type: 'text' }
          }
        }
      })
    });

    // Index 10 documents into "cities" index
    for (let i = 1; i <= 10; i++) {
      await fetch('http://localhost:9200/cities/_doc', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          name: `City${i}`,
          country: `Country${i}`
        })
      });
    }

    // Create index "countries" with mapping
    await fetch('http://localhost:9200/countries', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        settings: {},
        mappings: {
          properties: {
            name: { type: 'text' }
          }
        }
      })
    });

    // Index 10 documents into "countries" index
    for (let i = 1; i <= 10; i++) {
      await fetch('http://localhost:9200/countries/_doc', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          name: `Country${i}`
        })
      });
    }

    console.log('Data indexing completed.');
  } catch (error) {
    console.error('Error indexing data:', error.message);
  }
}

indexData();
