import { GET_ARTICLES_FROM_BACKEND_SUCCESS, ARTICLES_HAS_ERROR, ARITCLES_IS_LOADING } from "../Constants/articles";
import { ConfigureRequest } from "../API/requestBuilder";
import { ARTICLES_URL } from "../API/main";

export function getArticlesFromBackendSuccess(payload){
    return{
        type: GET_ARTICLES_FROM_BACKEND_SUCCESS,
        payload,
    }
}

export function getArticlesHasError(status){
    return {
        type: ARTICLES_HAS_ERROR,
        payload: status,
    }
}

export function getArticlesHasLoading(status){
    return {
        type: ARITCLES_IS_LOADING,
        payload: status,
    }
}

export function ArticlesFetchData(){
    return (dispatch) => {
        dispatch(getArticlesHasLoading(true));

        let request = ConfigureRequest(ARTICLES_URL, "GET", null);        

        console.log("start request to :", request);
        setTimeout(() => 
            fetch(request)
            .then((response) => {
                if (!response.ok) {
                    throw Error(response.statusText);
                }

                dispatch(getArticlesHasLoading(false));

                return response;
            })
            .then((response) => response.json())
            .then((items) => dispatch(getArticlesFromBackendSuccess(items)))
            .catch(() => {
                dispatch(getArticlesHasError(true));
                dispatch(getArticlesHasLoading(false));
            }),
            500    
        );
    }
}
