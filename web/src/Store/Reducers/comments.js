import { GET_COMMENTS_LOADING, GET_COMMENTS_FAILED, GET_COMMENTS_SUCCES, SET_COMMENTS } from "../Constants/comments";

const initState ={
    data: [],
    loaded: false,
    loading: false,
    failed: false,
}

export function CSI_commentsState(state = initState, action){
    const {type, payload} = action;

    switch(type){
        case GET_COMMENTS_LOADING:
            return {...state, loaded: false, loading: payload, failed: false}

        case GET_COMMENTS_FAILED:
            return{...state, loaded: false, loading: false, failed: payload}

        case GET_COMMENTS_SUCCES:
            return {...state, loaded: true, loading: false, failed: false, data: payload}

        case SET_COMMENTS:
            let dataT = state.data;
            dataT.push(payload);
            
            return {...state, data: dataT};
        
        default:
            return state;
    }
}