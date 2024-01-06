const http = require('http');

const options = {
    hostname: 'localhost',
    port: 8080,
    path: '/?count=3',
    method: 'GET'
};

const request = http.request(options, (res) => {
    console.log(`response headers: ${JSON.stringify(res.headers)}`);
    console.log(`response statusCode: ${res.statusCode}`);
    res.on('data', (chunk) => {
        console.log(`BODY: ${chunk}`);
    });

    res.on('end' , () => {
       console.log('No more data in response.');
    });

    res.on('error', (error)=> {
        console.log(`Error: ${error.message}`);
    })
});

request.end()