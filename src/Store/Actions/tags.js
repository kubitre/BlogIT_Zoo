import { GET_TAGS_FROM_BACKEND_SUCCESS, GET_TAGS_FROM_BACKEND_LOADING, GET_TAGS_FROM_BACKEND_FAILED } from "../Constants/tags";
import { ConfigureRequest } from "../API/requestBuilder";
import { TAGS_URL } from "../API/main";

export function getTagsFromBackendSucces(data){
    return {
        "type": GET_TAGS_FROM_BACKEND_SUCCESS,
        "payload": data,
    }
}

export function getTagsFromBackendLoading(status){
    return {
        "type": GET_TAGS_FROM_BACKEND_LOADING,
        "payload": status,
    }
}

export function getTagsFromBackendWithError(status){
    return {
        "type": GET_TAGS_FROM_BACKEND_FAILED,
        "payload": status,
    }
}

export function startFetchingTagsFromBackend(){
    return dispatch => {
        dispatch(getTagsFromBackendLoading(true))

        let request_packet = ConfigureRequest(TAGS_URL, "GET", null);

        fetch(request_packet)
        .then(response => {
            if(!response.ok){
                dispatch(getTagsFromBackendWithError(true));
                throw Error("error request");
            }
            dispatch(getTagsFromBackendLoading(false))
            return response
        })
        .then(reponse => reponse.json())
        .then(jsonresp => dispatch(getTagsFromBackendSucces(jsonresp)))
        .catch(err => getTagsFromBackendWithError(true))
    }
}