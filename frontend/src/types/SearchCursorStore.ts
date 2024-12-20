import type {SearchPageParams} from "@/types/SearchPageParams.ts";
import {SearchRequestResponse} from "@/types/SearchRequestResponse.ts";

export type SearchCursorStore = {
    originalSearchPageParams?: SearchPageParams,
    searchRequestResponse?: SearchRequestResponse,
}