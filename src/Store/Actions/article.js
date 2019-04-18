import { GET_ARTICLE_FROM_BACKEND_SUCCESS, GET_ARTICLE_LOADING, GET_ARTICLE_FAILED, CLEAR_ALL_DATA } from "../Constants/article";
import { ConfigureRequest } from "../API/requestBuilder";
import { ARTICLES_URL } from "../API/main";

export function getArticleSuccess(result){
    return {
        type: GET_ARTICLE_FROM_BACKEND_SUCCESS,
        payload: result
    }
}

export function getArticleHasLoading(status) {
    return {
        type: GET_ARTICLE_LOADING,
        payload: status
    }
}

export function getArticleHasError(err){
    return{
        type: GET_ARTICLE_FAILED,
        payload: err
    }
}

export function getArticleByID(identificator){
    return (dispatch) => {
        console.log('tururu');
        dispatch(getArticleHasLoading(true));

        let request = ConfigureRequest(ARTICLES_URL + "/"+ identificator, "GET", null)

        setTimeout(() => 
            fetch(request)
            .then((response ) => {
                if (!response.ok) {
                    throw Error(response.statusText)
                }

                dispatch(getArticleHasLoading(false));
                return response;
            })
            .then((reponse) => reponse.json())
            .then((item) => dispatch(getArticleSuccess(item)))
            .catch(err => getArticleHasError(err)),
            100
        );
    }
}

export function ArticleClear(){
    return {
        type: CLEAR_ALL_DATA,
        payload: null
    }
}