import axios from 'axios';
const baseUrl = 'https://api.mangadex.org';

class Manga {
    constructor(id, attributes) {
        this.id = id;
        this.attribute = attributes;
    }
}

axios.get(`${baseUrl}/manga`, {
    params: {
        title: "A boyish girlfriend in high humidity"
    }
}).then(response => {
    const obj = response.data
    console.log(obj)
    }).catch(error => {
        console.error('error fetching data', error)
    });




