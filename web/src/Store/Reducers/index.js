import { combineReducers } from 'redux';
import {AI_stateArticles, AI_itemsHasError, AI_itemsIsLoading} from './articles';
import {ASI_stateArticle} from './article';
import {TSI_tagState} from './tags';
import {USI_userState} from './user';
import {CSI_commentsState} from './comments';

export default combineReducers({
    AI_stateArticles,
    AI_itemsIsLoading,
    AI_itemsHasError,
    ASI_stateArticle,
    TSI_tagState,
    USI_userState,
    CSI_commentsState,
});