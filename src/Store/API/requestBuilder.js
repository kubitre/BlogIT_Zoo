import { ARTICLES_URL } from "./main";
import { SUCCESS_GET_REQUEST, FAILED_GET_REQUEST } from "../Constants/api";

export function GetRequestToBackend(type){
    var myHeaders = new Headers({
        "Content-Type": "application/json",
        "Accept": "application/json",
        "Access-Control-Allow-Origin": "*",
        "Origin": "*",
    });

    var myInit = { 
        method: 'GET',
        headers: myHeaders,
        mode: "cors",
        cache: 'default' 
    };

    const request = new Request(ARTICLES_URL, 
        myInit
    )

    fetch(request)
    .then((result) => {
        result.json().then(values => {
            return {
                "type": SUCCESS_GET_REQUEST,
                "payload": {
                    values
                }
            }
        })
    })
    .catch(err => 
    {
        return {
            "type": FAILED_GET_REQUEST,
            "payload": err,
        }
    }
        )
}

export function ConfigureRequest(url, methodType, body){
    let myHeaders = new Headers({
        "Content-Type": "application/json",
        "Accept": "application/json",
        "Access-Control-Allow-Origin": "*",
        "Origin": "*",
    });

    let myInit = { 
        method: methodType,
        headers: myHeaders,
        mode: "cors",
        cache: 'default',
        body: body
    };

    let request = new Request(url, myInit)
    return request
}


export function ConfigureRequestWithToken(url, methodType, body, token){
    let myHeaders = new Headers({
        "Content-Type": "application/json",
        "Accept": "application/json",
        "Access-Control-Allow-Origin": "*",
        "Origin": "*",
        "Authorization": "Bearer " + token,
    });

    let myInit = { 
        method: methodType,
        headers: myHeaders,
        mode: "cors",
        cache: 'default',
        body: body
    };

    let request = new Request(url, myInit)
    return request
}