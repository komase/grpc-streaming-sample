const fetch = require('node-fetch');

async function main() {
    try{
        const response = await fetch('http://localhost:8080/?count=3');
        console.log(`response headers: ${JSON.stringify(response.headers)}`);
        console.log(`response statusCode: ${response.status}`);
        const data = await response.text();
        console.log(`response data: ${data}`);
    } catch (error) {
        console.log(`Error: ${error.message}`);
    }
}

main();