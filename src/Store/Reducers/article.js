import { GET_ARTICLE_LOADING, GET_ARTICLE_FAILED, GET_ARTICLE_FROM_BACKEND_SUCCESS, CLEAR_ALL_DATA } from "../Constants/article";

const initState = {
    data: null,
    loading: false,
    loaded: false, 
    failed: false,
}

export function ASI_stateArticle(state = initState, action){
    const {type, payload} = action;

    switch(type){
        case GET_ARTICLE_LOADING:
            return {...state, loading: payload, loaded: false}
        case GET_ARTICLE_FAILED:
            return {...state, failed: payload, loaded: false}
        case GET_ARTICLE_FROM_BACKEND_SUCCESS:
            return {...state, data: payload, loaded: true}
        case CLEAR_ALL_DATA:
            return {...state, data: null, loading: false, failed: false, loaded: false}
        default:
            return state;
    }
}