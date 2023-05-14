import React, {useEffect, useState} from 'react';
import './style.css';
import axios from "axios";

const baseURL = "http://localhost:8080"

function App() {
    const [request, setRequest] = useState('Paste link here')
    const [shortUrl, setShortURL] = useState('')

    function sendTheURL(value) {
        setRequest(value);
    }

    function setShortURLAxios(response) {
        setShortURL(response.url);
    }

    function sendRequest() {
        axios.post(`${baseURL}/short`, {
            url: request,
        }).then((response) => {
            setShortURLAxios(response.data);
        })
    }

    return (
        <div className="App">
            <h1>{shortUrl}</h1>
            <input type='text'
                   value={request}
                   onChange={event => sendTheURL(event.target.value)}
            />
            <button type='button'
                    onClick={sendRequest}

            >Send me</button>
        </div>
    );
}

export default App;