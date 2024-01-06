const axios = require('axios');

async function main() {
    try{
        const response = await axios.get('http://localhost:8080/?count=3');
        console.log(`response headers: ${JSON.stringify(response.headers)}`);
        console.log(`response statusCode: ${response.status}`);
        console.log(`response data: ${response.data}`);
    } catch (error) {
        console.log(`Error: ${error.message}`);
    }
}

main();