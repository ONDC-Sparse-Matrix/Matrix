import {getZipcodes} from "fakexy";
import fs from 'fs';

console.time("Time taken start");

for(let i=0;i<1000;i++){
    getZipcodes({countryCode:'in'}).then((data) => {
        Object.keys(data).forEach((key)=>{
            fs.appendFile('zipcodes.txt', key + '\n', function (err) {
                if (err) throw err;
            }
        );
    })
}
);
}
console.timeEnd("Time taken start");