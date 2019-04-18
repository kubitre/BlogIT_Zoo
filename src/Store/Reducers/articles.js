import { ARITCLES_IS_LOADING, ARTICLES_HAS_ERROR, GET_ARTICLES_FROM_BACKEND_SUCCESS } from "../Constants/articles";

export function AI_stateArticles(state = {"articles": []}, action){
    const {type, payload} = action;

    switch(type){
        case GET_ARTICLES_FROM_BACKEND_SUCCESS:
            return {...state, articles: payload}
        default:
            return state;
    }
}

export function AI_itemsIsLoading(state = false, action){
    const {type, payload} = action;

    switch(type){
        case ARITCLES_IS_LOADING:
            return payload;
            
        default:
            return state;
    }
}

export function AI_itemsHasError(state = false, action){
    const {type, payload} = action;

    switch(type){
        case ARTICLES_HAS_ERROR:
            return payload;
        default:
            return state;
    }
}