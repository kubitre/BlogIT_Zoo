import { GET_TAGS_FROM_BACKEND_SUCCESS, GET_TAGS_FROM_BACKEND_LOADING, GET_TAGS_FROM_BACKEND_FAILED } from "../Constants/tags";

const initState = {
    loading: false,
    loaded: false,
    failed: false,
    data: null
}

export function TSI_tagState(state = initState, action){
    const {type, payload} = action;

    switch(type){
        case GET_TAGS_FROM_BACKEND_SUCCESS:
            return {...state, data: payload, loaded: true}

        case GET_TAGS_FROM_BACKEND_LOADING:
            return {...state, loading: payload}

        case GET_TAGS_FROM_BACKEND_FAILED:
            return {...state, loading: false, loaded: false, failed: payload}
        default:
            return state;
    }
}